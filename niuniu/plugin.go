package niuniu

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world"
	"github.com/kohmebot/pkg/command"
	"github.com/kohmebot/pkg/version"
	"github.com/kohmebot/plugin"
	"github.com/wdvxdr1123/ZeroBot"
)

type PluginNiuNiu struct {
	w *world.World
}

func NewPlugin() plugin.Plugin {
	return &PluginNiuNiu{}
}

func (p *PluginNiuNiu) Init(engine *zero.Engine, env plugin.Env) (err error) {
	p.w, err = world.NewWorld(env)
	if err != nil {
		return
	}
	p.w.InitWorld(engine)

	return
}

func (p *PluginNiuNiu) Name() string {
	return "niuniu"
}

func (p *PluginNiuNiu) Description() string {
	return ""
}

func (p *PluginNiuNiu) Commands() fmt.Stringer {
	return command.NewCommands()
}

func (p *PluginNiuNiu) Version() uint64 {
	return uint64(version.NewVersion(0, 0, 10))
}

func (p *PluginNiuNiu) OnBoot() {

}
