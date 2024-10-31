package trigger

import (
	"github.com/kohmebot/niuniu/niuniu/world/player"
	"github.com/kohmebot/plugin"
	zero "github.com/wdvxdr1123/ZeroBot"
	"strings"
)

var triggers = []Trigger{
	new(RegisterTrigger),
	new(LookTrigger),
	new(PKTrigger),
}

const (
	niuniuNotFound       = `你没有牛牛，发送"领养牛牛 名称"来领养一只可爱的牛牛吧！"`
	targetNiuniuNotFound = `对方没有牛牛！`
)

type Trigger interface {
	// TriggerText 触发文本
	TriggerText() []string
	// On 触发
	On(ctx *zero.Ctx, g *player.Generator) error
}

func Register(engine *zero.Engine, g *player.Generator, env plugin.Env) {
	for _, trigger := range triggers {
		engine.OnPrefixGroup(trigger.TriggerText(), env.Groups().Rule()).SetBlock(true).Handle(func(ctx *zero.Ctx) {
			err := trigger.On(ctx, g)
			if err != nil {
				env.Error(ctx, err)
			}
		})
	}
}

func ParseArgs(ctx *zero.Ctx) []string {
	res := make([]string, 0, 1)
	if ctx.Event == nil {
		return res
	}
	for _, segment := range ctx.Event.Message {
		if segment.Type == "text" {
			res = append(res, strings.Split(segment.Data["text"], " ")...)
		}
	}
	return res
}
