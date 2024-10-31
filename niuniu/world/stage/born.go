package stage

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/buff"
	"github.com/kohmebot/niuniu/niuniu/world/player"
	"github.com/kohmebot/niuniu/niuniu/world/skill"
	"strings"
	"time"
)

type BornStage struct {
	niu *player.NiuNiu
}

func NewBornStage(niu *player.NiuNiu) *BornStage {
	return &BornStage{niu: niu}
}

func (b *BornStage) Do() base.Narration {
	var builder strings.Builder
	pbuff := buff.NewPowerBuff(time.Now().Add(3 * time.Hour))
	pSkill := &skill.PushSkill{}
	builder.WriteString(fmt.Sprintf(`似乎是某位好战者的牛子转生了？\n%s获得了Buff"%s\n获得了技能%s"`, b.niu.Name, pbuff.Name(), pSkill.Name()))
	b.niu.Buffs = append(b.niu.Buffs, pbuff)
	b.niu.Skills = append(b.niu.Skills, pSkill)
	return &builder
}
