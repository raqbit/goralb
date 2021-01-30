package goralb

type BrushState int

//go:generate stringer -type=BrushState -output state_string.go -linecomment
const (
	StateUnknown BrushState = 0x00 // Unknown
	StateIdle    BrushState = 0x02 // Idle
	StateRun     BrushState = 0x03 // Run
	StateCharge  BrushState = 0x04 // Charge
)
