package goralb

type statusControl byte

const (
	SetMode          statusControl = 0x01
	StopTimerSignal  statusControl = 0x20
	ResetMemoryTimer statusControl = 0x29
	MyColorDisable   statusControl = 0x30
	MyColorEnable    statusControl = 0x31
	MotionDisable    statusControl = 0x40
	MotionEnable     statusControl = 0x41
)
