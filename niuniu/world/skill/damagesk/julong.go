package damagesk

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"github.com/kohmebot/niuniu/niuniu/world/skill"
	"github.com/kohmebot/niuniu/niuniu/world/util"
)

func init() {
	util.RegGob[JuLong]()
}

type JuLong struct {
}

func (p *JuLong) Name() string {
	return "巨龙撞击"
}

func (p *JuLong) Desc() string {
	return "对敌方造成80%攻击力的真实伤害"
}

func (p *JuLong) Do(ctx *event.Context, cs ...base.Character) {
	attacker := cs[0]
	defender := cs[1]
	_, _, use := skill.IsUse[*JuLong](ctx)
	if use {
		return
	}
	damage := attacker.Status().Power * 0.8
	defender.HP().Hit(damage)
	e := skill.NewEventSkill[*JuLong](attacker, skill.Damage, fmt.Sprintf("造成了%.2f点伤害！", damage))
	ctx.AppendEvent(e)
	attacker.EventContext().AppendEvent(e)

}

func (p *JuLong) String() string {
	return skill.CommonFormat(p)
}
