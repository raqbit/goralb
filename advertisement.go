package goralb

import (
	"time"
)

type Advertisement struct {
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

func ParseAdvertisement(data []byte) *Advertisement {
	return &Advertisement{
		ProtocolVersion:      int(data[0]),
		TypeID:               int(data[1]),
		FirmwareVersion:      int(data[2]),
		State:                BrushState(data[3]),
		PressureDetected:     data[4]&0x80 != 0,
		HasReducedMotorSpeed: data[4]&0x40 != 0,
		HasProfessionalTimer: data[4]&0x1 == 0,
		BrushTime:            time.Duration(data[5]*60+data[6]) * time.Second, // (mins * secs_in_min + secs) * nanos_in_sec
		Mode:                 BrushMode(data[7]),
		Sector:               ParseSector(int(data[8] & 0x07)),
		Smiley:               int(data[8] & 0x38 >> 3),
	}
}
