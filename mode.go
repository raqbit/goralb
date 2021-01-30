package goralb

//go:generate stringer -type=BrushMode -output mode_string.go -linecomment
type BrushMode int

const (
	ModeOff            BrushMode = 0x00 // Off
	ModeDailyClean     BrushMode = 0x01 // Daily
	ModeSensitive      BrushMode = 0x02 // Sensitive
	ModeMassage        BrushMode = 0x03 // Massage
	ModeWhitening      BrushMode = 0x04 // Whitening
	ModeDeepClean      BrushMode = 0x05 // Deep Clean
	ModeTongueCleaning BrushMode = 0x06 // Tongue Cleaning
	ModeTurbo          BrushMode = 0x07 // Turbo
	ModeUnknown        BrushMode = 0xFF // Unknown
)
