package bornsc

import (
	"github.com/kohmebot/niuniu/niuniu/world/entity"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"strings"
)

type HaoZhanScene struct {
}

func NewHaoZhanScene() *HaoZhanScene {
	return new(HaoZhanScene)
}

func (h *HaoZhanScene) Do(ctx *event.Context, es ...entity.Entity) {
	var s strings.Builder
	p := ExtractNiuNiu(es[0])
	s.WriteString("似乎是某位击剑好战者的牛牛转生了？\n")
	p.Length += p.Length * 0.1
	s.WriteString("长度增加了30%，")
	p.Hardness += p.Hardness * 0.05
	s.WriteString("硬度增加了5%，")
	p.Flexibility -= p.Flexibility * 0.3
	s.WriteString("弹性减少了30%")
	e := event.TextEvent(s.String())
	ctx.AppendEvent(e)
	es[0].EventContext().AppendEvent(e)
}
