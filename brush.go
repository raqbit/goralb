package goralb

import (
	"context"
	"errors"
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/muka/go-bluetooth/bluez/profile/device"
	"github.com/muka/go-bluetooth/bluez/profile/gatt"
	"image/color"
	"time"
)

type (
	Brush interface {
		Connect(timeout context.Context) error
		Disconnect() error
		SetRingEnabled(enabled bool) error
		SetMotionEnabled(enabled bool) error
		GetColor() (color.RGBA, error)
		SetColor(rgba color.RGBA) error
		Info() *BrushInfo
		String() string
	}

	BrushInfo struct {
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

	brush struct {
		btDev *device.Device1
		info  *BrushInfo

		status *gatt.GattCharacteristic1
		color  *gatt.GattCharacteristic1
	}
)

const (
	statusChar        string = "a0f0ff21-5047-4d53-8208-4f72616c2d42"
	rtcChar           string = "a0f0ff22-5047-4d53-8208-4f72616c2d42"
	timezoneChar      string = "a0f0ff23-5047-4d53-8208-4f72616c2d42"
	brushingTimerChar string = "a0f0ff24-5047-4d53-8208-4f72616c2d42"
	brushingModesChar string = "a0f0ff25-5047-4d53-8208-4f72616c2d42"
	quadrantTimesChar string = "a0f0ff26-5047-4d53-8208-4f72616c2d42"
	tongueTimeChar    string = "a0f0ff27-5047-4d53-8208-4f72616c2d42"
	pressureChar      string = "a0f0ff28-5047-4d53-8208-4f72616c2d42"
	dataChar          string = "a0f0ff29-5047-4d53-8208-4f72616c2d42"
	flightModeChar    string = "a0f0ff2a-5047-4d53-8208-4f72616c2d42"
	colorChar         string = "a0f0ff2b-5047-4d53-8208-4f72616c2d42"

	statusControlId   byte = 0x10
	statusConfigureId byte = 0x37
)

func (b *brush) Connect(timeout context.Context) error {
	var err error

	if err = b.btDev.Connect(); err != nil {
		return fmt.Errorf("could not connect to brush: %w", err)
	}

	if err = b.waitForServicesResolved(timeout); err != nil {
		return fmt.Errorf("could not resolve services: %w", err)
	}

	b.status, err = b.btDev.GetCharByUUID(statusChar)

	if err != nil {
		return fmt.Errorf("could not get status characteristic: %w", err)
	}

	b.color, err = b.btDev.GetCharByUUID(colorChar)

	if err != nil {
		return fmt.Errorf("could get color characteristic: %w", err)
	}

	return nil
}

func (b *brush) Disconnect() error {
	if err := b.btDev.Disconnect(); err != nil {
		return fmt.Errorf("could not disconnect from brush: %w", err)
	}
	return nil
}

func (b *brush) SetRingEnabled(enabled bool) error {
	option := MyColorDisable

	if enabled {
		option = MyColorEnable
	}

	return b.control(option)
}

func (b *brush) SetMotionEnabled(enabled bool) error {
	option := MotionDisable

	if enabled {
		option = MotionEnable
	}

	return b.control(option)
}

func (b *brush) GetColor() (color.RGBA, error) {
	data, err := b.color.ReadValue(nil)

	if err != nil {
		return color.RGBA{}, fmt.Errorf("could not read color: %w", err)
	}

	return color.RGBA{
		R: data[0],
		G: data[1],
		B: data[2],
	}, nil
}

func (b *brush) SetColor(rgba color.RGBA) error {
	var err error

	err = b.color.WriteValue([]byte{
		rgba.R,
		rgba.G,
		rgba.B,
		0x00,
	}, nil)

	if err != nil {
		return fmt.Errorf("could not set brush color: %w", err)
	}

	err = b.configure(MyColor)

	if err != nil {
		return fmt.Errorf("could not configure brush color: %w", err)
	}

	return nil
}

func (b brush) Info() *BrushInfo {
	return b.info
}

func (b brush) String() string {
	info := b.Info()

	return fmt.Sprintf(
		`Oral-B ToothbrushType: %d, Firmware version: v%d, Mode: %s, State: %s, Sector: %s, Pressure: %v`,
		info.TypeID, info.FirmwareVersion, info.Mode, info.State, info.Sector, info.PressureDetected)
}

func (b brush) control(option statusControl) error {
	return b.status.WriteValue([]byte{
		statusControlId,
		byte(option),
	}, nil)
}

func (b *brush) configure(option statusConfigure) error {
	return b.status.WriteValue([]byte{
		statusConfigureId,
		byte(option),
	}, nil)
}

func (b *brush) waitForServicesResolved(timeout context.Context) error {
	changed, err := b.btDev.WatchProperties()
	//defer b.btDev.UnwatchProperties(changed)

	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case change := <-changed:
			if change.Name == "ServicesResolved" && change.Value == true {
				return nil
			}
		case <-timeout.Done():
			return errors.New("services did not resolve within given timeout")
		}
	}
}

func NewBrush(dev *device.Device1) Brush {
	mfd := dev.Properties.ManufacturerData[PGCompanyID]
	mfBytes := mfd.(dbus.Variant).Value().([]byte)

	b := &brush{
		btDev: dev,
		info: &BrushInfo{
			ProtocolVersion:      int(mfBytes[0]),
			TypeID:               int(mfBytes[1]),
			FirmwareVersion:      int(mfBytes[2]),
			State:                BrushState(mfBytes[3]),
			PressureDetected:     mfBytes[4]&0x80 != 0,
			HasReducedMotorSpeed: mfBytes[4]&0x40 != 0,
			HasProfessionalTimer: mfBytes[4]&0x1 == 0,
			BrushTime:            time.Duration(mfBytes[5]*60+mfBytes[6]) * time.Second, // (mins * secs_in_min + secs) * nanos_in_sec
			Mode:                 BrushMode(mfBytes[7]),
			Sector:               BrushSector(int(mfBytes[8])),
			Smiley:               int(mfBytes[8] & 0x38 >> 3),
		},
	}

	return b
}
