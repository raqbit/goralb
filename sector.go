package goralb

import (
	"fmt"
)

type BrushSector int

func (b BrushSector) String() string {
	return fmt.Sprintf("%d", b)
}

const (
	Sector1 BrushSector = 0x01
	Sector2 BrushSector = 0x02
	Sector3 BrushSector = 0x03
	Sector4 BrushSector = 0x04
	Sector5 BrushSector = 0x05
	Sector6 BrushSector = 0x06
	Sector7 BrushSector = 0x07
	Sector8 BrushSector = 0x08
)
