package goralb

import (
	"context"
	"errors"
	"fmt"
	"github.com/godbus/dbus"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/device"
)

const PGCompanyID = 0xDC

// Searches for a single brush
func FindBrush(ctx context.Context) (*Advertisement, error) {
	brushes, err := FindBrushes(ctx, 1)

	if err != nil {
		return nil, err
	}

	if len(brushes) == 0 {
		return nil, errors.New("could not find brush")
	}

	return brushes[0], nil
}

// Searches for specified amount of brushes
func FindBrushes(ctx context.Context, count int) ([]*Advertisement, error) {
	defer api.Exit()

	adapt, err := adapter.GetDefaultAdapter()

	if err != nil {
		return nil, err
	}

	err = flushBrushDiscoveries(adapt)

	if err != nil {
		return nil, fmt.Errorf("could not flush brush discoveries: %w", err)
	}

	adverts, err := discoverBrushes(ctx, adapt, count)

	if err != nil {
		return nil, fmt.Errorf("could not discover brushes: %w", err)
	}

	return adverts, nil
}

func flushBrushDiscoveries(adapt *adapter.Adapter1) error {
	devices, err := adapt.GetDevices()

	if err != nil {
		return err
	}

	for _, dev := range devices {
		// Do not try flush connected devices
		if dev.Properties.Connected {
			continue
		}

		// There should only be one companyId per device, but the data is exposed as a map
		for companyId := range dev.Properties.ManufacturerData {
			if companyId == PGCompanyID {
				// Remove device, ignore when unsuccessful
				err = adapt.RemoveDevice(dev.Path())

				if err != nil {
					return fmt.Errorf("could not remove %s from brush cache: %w", dev.Properties.Address, err)
				}

				// We only care about this companyID
				break
			}
		}
	}

	return nil
}

func discoverBrushes(ctx context.Context, adapt *adapter.Adapter1, count int) ([]*Advertisement, error) {
	// Only discover LE devices and do not give duplicates
	filter := &adapter.DiscoveryFilter{
		Transport:     adapter.DiscoveryFilterTransportLE,
		DuplicateData: false,
	}

	// Discover new devices
	discoveries, cancel, err := api.Discover(adapt, filter)
	defer cancel()

	if err != nil {
		return nil, err
	}

	adverts := make([]*Advertisement, 0, count)

	for {
		select {
		// When a new device has been discovered
		case discovery := <-discoveries:

			// Ignore devices which have been removed
			if discovery.Type == adapter.DeviceRemoved {
				continue
			}

			// Create device handle
			dev, err := device.NewDevice1(discovery.Path)

			if err != nil || dev == nil {
				continue
			}

			// Check if device has the right companyID
			mfd, exists := dev.Properties.ManufacturerData[PGCompanyID]
			if !exists {
				continue
			}

			// Get the raw manufacturer data bytes
			mfBytes := mfd.(dbus.Variant).Value().([]byte)

			// Parse advertisement
			adv := ParseAdvertisement(mfBytes)

			adverts = append(adverts, adv)

			if len(adverts) == count {
				return adverts, nil
			}
		case <-ctx.Done():
			return adverts, nil
		}
	}
}
