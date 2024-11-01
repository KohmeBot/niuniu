package bornsc

import (
	"github.com/kohmebot/niuniu/niuniu/world/entity"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"strings"
)

type PuTongScene struct {
}

func NewPuTongScene() *PuTongScene {
	return new(PuTongScene)
}

func (h *PuTongScene) Do(ctx *event.Context, es ...entity.Entity) {
	var s strings.Builder
	p := ExtractNiuNiu(es[0])
	s.WriteString("嗯,普普通通\n")
	p.Length += p.Length * 0.15
	s.WriteString("长度增加了15%，")
	p.Hardness += p.Hardness * 0.15
	s.WriteString("硬度增加了15%，")
	p.Flexibility += p.Flexibility * 0.15
	s.WriteString("弹性增加了15%，")
	p.Luck += p.Luck * 0.15
	s.WriteString("幸运值增加了15%")
	e := event.TextEvent(s.String())
	ctx.AppendEvent(e)
	es[0].EventContext().AppendEvent(e)
}
