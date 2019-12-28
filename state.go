package goralb

type BrushState int

func (b BrushState) String() string {
	str := []string{
		"Idle",
		"Run",
		"Charge",
		"Unknown",
	}
	return str[int(b)]
}

const (
	StateUnknown BrushState = 0x00
	StateIdle    BrushState = 0x02
	StateRun     BrushState = 0x03
	StateCharge  BrushState = 0x04
)
