package goralb

type statusConfigure byte

const (
	ExtendConnection  statusConfigure = 0x00
	ReadParam         statusConfigure = 0x01
	ReadData          statusConfigure = 0x02
	CalibrationRead   statusConfigure = 0x04
	ReadMetadata      statusConfigure = 0x05
	RTC               statusConfigure = 0x26
	BrushTimer        statusConfigure = 0x28
	BrushModes        statusConfigure = 0x29
	QuadrantTimers    statusConfigure = 0x2a
	TongueTime        statusConfigure = 0x2b
	MyColor           statusConfigure = 0x2f
	Dashboard         statusConfigure = 0x30
	RefillReminder    statusConfigure = 0x31
	FactoryReset      statusConfigure = 0x32
	SmartGuideDisable statusConfigure = 0x50
	SmartGuideEnable  statusConfigure = 0x51
)
