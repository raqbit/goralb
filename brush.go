package goralb

import (
	"context"
	"errors"
	"fmt"
	"github.com/godbus/dbus"
	"github.com/muka/go-bluetooth/bluez/profile/device"
	"image/color"
	"time"
)

type Brush struct {
	btDev                *device.Device1
	ProtocolVersion      int
	TypeID               int
	FirmwareVersion      int
	State                BrushState
	PressureDetected     bool
	HasReducedMotorSpeed bool
	HasProfessionalTimer bool
	BrushTime            time.Duration
	Mode                 BrushMode
	Sector               BrushSector
	Smiley               int
}

func (brush *Brush) WaitForServicesResolved(timeout context.Context) error {
	changed, err := brush.btDev.WatchProperties()

	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case change := <-changed:
			if change.Name == "ServicesResolved" && change.Value == true {
				// TODO: Should probably "unwatch", but that only closes the "changed" channel and then panic's the bt lib
				return nil
			}
		case <-timeout.Done():
			return errors.New("services did not resolve within given timeout")
		}
	}
}

func (brush *Brush) Connect() error {
	if err := brush.btDev.Connect(); err != nil {
		return fmt.Errorf("could not connect to brush: %w", err)
	}
	return nil
}

func (brush *Brush) Disconnect() error {
	if err := brush.btDev.Disconnect(); err != nil {
		return fmt.Errorf("could not disconnect from brush: %w", err)
	}
	return nil
}

// TODO: rewrite, this is quick'n dirty
func (brush *Brush) SetColor(rgba color.RGBA) error {
	colorChar, err := brush.btDev.GetCharByUUID("a0f0ff2b-5047-4d53-8208-4f72616c2d42")

	if err != nil {
		return fmt.Errorf("could get color characteristic: %w", err)
	}

	statusChar, err := brush.btDev.GetCharByUUID("a0f0ff21-5047-4d53-8208-4f72616c2d42")

	if err != nil {
		return fmt.Errorf("could not get status characteristic: %w", err)
	}

	err = colorChar.WriteValue([]byte{
		rgba.R,
		rgba.G,
		rgba.B,
		0x00,
	}, nil)

	if err != nil {
		return fmt.Errorf("could not set brush color: %w", err)
	}

	err = statusChar.WriteValue([]byte{
		0x10,
		0x31,
	}, nil)

	if err != nil {
		return fmt.Errorf("could not set brush color: %w", err)
	}

	err = statusChar.WriteValue([]byte{
		0x37,
		0x2f,
	}, nil)

	if err != nil {
		return fmt.Errorf("could not set brush color: %w", err)
	}

	return nil
}

func NewBrush(dev *device.Device1) *Brush {
	mfd := dev.Properties.ManufacturerData[PGCompanyID]
	mfBytes := mfd.(dbus.Variant).Value().([]byte)

	b := &Brush{
		btDev:                dev,
		ProtocolVersion:      int(mfBytes[0]),
		TypeID:               int(mfBytes[1]),
		FirmwareVersion:      int(mfBytes[2]),
		State:                BrushState(mfBytes[3]),
		PressureDetected:     mfBytes[4]&0x80 != 0,
		HasReducedMotorSpeed: mfBytes[4]&0x40 != 0,
		HasProfessionalTimer: mfBytes[4]&0x1 == 0,
		BrushTime:            time.Duration(mfBytes[5]*60+mfBytes[6]) * time.Second, // (mins * secs_in_min + secs) * nanos_in_sec
		Mode:                 BrushMode(mfBytes[7]),
		Sector:               ParseSector(int(mfBytes[8] & 0x07)),
		Smiley:               int(mfBytes[8] & 0x38 >> 3),
	}

	return b
}
