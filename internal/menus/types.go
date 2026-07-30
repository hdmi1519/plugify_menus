package menus

type Option struct {
	Type int32
	Back string
	Text string
}

type Instance struct {
	Title           string
	MaxVisibleItems int32
	HasBack         bool
	Options         []Option
	CurrentOption   int
	CurrentPage     int
	Callback        func(action string, slot int32)
	AutoCloseAfter  float64
	OpenedAt        float64
}

type Session struct {
	Menu           *Instance
	History        []*Instance
	LastButtons    uint64
	LastClickTime  float64
	LastRenderTime float64
	IsActive       bool
	IsLocked       bool
}
