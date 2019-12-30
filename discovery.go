package goralb

import (
	"context"
	"errors"
	"fmt"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/device"
)

const PGCompanyID = 0xDC

type BrushManager struct{}

func NewBrushManager() *BrushManager {
	return &BrushManager{}
}

// Searches for a single brush
func (bm BrushManager) FindBrush(ctx context.Context) (*Brush, error) {
	brushes, err := bm.FindBrushes(ctx, 1)

	if err != nil {
		return nil, err
	}

	if len(brushes) == 0 {
		return nil, errors.New("could not find brush")
	}

	return brushes[0], nil
}

// Searches for specified amount of brushes
func (bm BrushManager) FindBrushes(ctx context.Context, count int) ([]*Brush, error) {
	adapt, err := adapter.GetDefaultAdapter()

	if err != nil {
		return nil, err
	}

	err = bm.flushBrushDiscoveries(adapt)

	if err != nil {
		return nil, fmt.Errorf("could not flush brush discoveries: %w", err)
	}

	adverts, err := bm.discoverBrushes(ctx, adapt, count)

	if err != nil {
		return nil, fmt.Errorf("could not discover brushes: %w", err)
	}

	return adverts, nil
}

func (bm BrushManager) flushBrushDiscoveries(adapt *adapter.Adapter1) error {
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

func (bm BrushManager) discoverBrushes(ctx context.Context, adapt *adapter.Adapter1, count int) ([]*Brush, error) {
	// Only discover LE devices and do not give duplicates
	filter := &adapter.DiscoveryFilter{
		Transport:     adapter.DiscoveryFilterTransportLE,
		DuplicateData: false,
	}

	// Discover new devices
	discoveries, cancel, err := api.Discover(adapt, filter)

	if err != nil {
		return nil, err
	}

	// FIXME: this sometimes hangs
	defer cancel()

	brushes := make([]*Brush, 0, count)

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
			if _, exists := dev.Properties.ManufacturerData[PGCompanyID]; !exists {
				continue
			}

			// Create brush from found device, append to list
			brushes = append(brushes, NewBrush(dev))

			if len(brushes) == count {
				return brushes, nil
			}
		case <-ctx.Done():
			return brushes, nil
		}
	}
}

func (bm BrushManager) Close() error {
	return api.Exit()
}
