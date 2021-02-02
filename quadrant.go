package goralb

import (
	"fmt"
)

type BrushQuadrant int

func (b BrushQuadrant) String() string {
	return fmt.Sprintf("%d", b)
}

const (
	Sector1 BrushQuadrant = 0x01
	Sector2 BrushQuadrant = 0x02
	Sector3 BrushQuadrant = 0x03
	Sector4 BrushQuadrant = 0x04
	Sector5 BrushQuadrant = 0x05
	Sector6 BrushQuadrant = 0x06
	Sector7 BrushQuadrant = 0x07
	Sector8 BrushQuadrant = 0x08
)
