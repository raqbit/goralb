package goralb

import "strconv"

type BrushSector int

func (b BrushSector) String() string {
	if b == Sector8 {
		return "last"
	} else if 0x00 <= b && b <= 0x06 {
		return strconv.Itoa(int(b + 1))
	} else {
		return "none"
	}
}

const (
	Sector1    BrushSector = 0x00
	Sector2    BrushSector = 0x01
	Sector3    BrushSector = 0x02
	Sector4    BrushSector = 0x03
	Sector5    BrushSector = 0x04
	Sector6    BrushSector = 0x05
	Sector7    BrushSector = 0x06
	Sector8    BrushSector = 0x07
	SectorLast BrushSector = 0xFE
	SectorNone BrushSector = 0xFF
)

func ParseSector(value int) BrushSector {
	if value == int(Sector8) {
		return SectorLast
	} else if 0x00 <= value && value <= 0x06 {
		return BrushSector(value - 1)
	} else {
		return SectorNone
	}
}
