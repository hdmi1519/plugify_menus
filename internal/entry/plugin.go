package entry

import (
	polyhook "github.com/fr0nch/go-plugify-polyhook/v2"
	"github.com/untrustedmodders/go-plugify"
)

type Plugin struct {
	gameDataHandle uint32
	hookHandle     polyhook.HookHandle
}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Run() {
	plugify.NewPlugin("menus", p.start, nil, p.stop)
}
