package bornsc

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/entity"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"github.com/kohmebot/niuniu/niuniu/world/skill/damagesk"
	"strings"
)

type NanNiangScene struct {
}

func NewNanNiangScene() *NanNiangScene {
	return new(NanNiangScene)
}

func (h *NanNiangScene) Do(ctx *event.Context, es ...entity.Entity) {
	var s strings.Builder
	p := ExtractNiuNiu(es[0])
	s.WriteString("嗯，哪来的男娘？！\n")
	p.Length -= p.Length * 0.2
	s.WriteString("长度减少了20%，")
	p.Hardness -= p.Hardness * 0.5
	s.WriteString("硬度减少了50%，")
	p.Flexibility += p.Flexibility * 0.6
	s.WriteString("弹性增加了60%，")
	p.Luck += p.Luck * 1
	s.WriteString("幸运值提升100%\n")
	keai := new(damagesk.KeAi)
	p.Skills = append(p.Skills, keai)
	s.WriteString(fmt.Sprintf("获得了技能: %s", keai.Name()))
	e := event.TextEvent(s.String())
	ctx.AppendEvent(e)
	es[0].EventContext().AppendEvent(e)
}
