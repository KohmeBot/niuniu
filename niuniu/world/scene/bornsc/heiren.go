package bornsc

import (
	"github.com/kohmebot/niuniu/niuniu/world/entity"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"strings"
)

type HeiRenScene struct {
}

func NewHeiRenScene() *HeiRenScene {
	return new(HeiRenScene)
}

func (h *HeiRenScene) Do(ctx *event.Context, es ...entity.Entity) {
	var s strings.Builder
	p := ExtractNiuNiu(es[0])
	s.WriteString("SSR！竟然是黑人牛牛\n")
	p.Length += p.Length * 0.5
	s.WriteString("长度增加了50%，")
	p.Hardness += p.Hardness * 0.5
	s.WriteString("硬度增加了50%，")
	p.Flexibility -= p.Flexibility * 0.5
	s.WriteString("弹性减少了50%，")
	p.Luck = 0
	s.WriteString("幸运值归零")
	e := event.TextEvent(s.String())
	ctx.AppendEvent(e)
	es[0].EventContext().AppendEvent(e)
}
