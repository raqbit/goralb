package goralb

type BrushMode int

func (b BrushMode) String() string {
	str := []string{
		"Off",
		"Daily",
		"Sensitive",
		"Massage",
		"Whitening",
		"DeepClean",
		"TongueCleaning",
		"Turbo",
		"Unknown",
	}
	return str[int(b)]
}

const (
	ModeOff            BrushMode = 0x00
	ModeDailyClean     BrushMode = 0x01
	ModeSensitive      BrushMode = 0x02
	ModeMassage        BrushMode = 0x03
	ModeWhitening      BrushMode = 0x04
	ModeDeepClean      BrushMode = 0x05
	ModeTongueCleaning BrushMode = 0x06
	ModeTurbo          BrushMode = 0x07
	ModeUnknown        BrushMode = 0xFF
)
