package skill

import (
	"encoding/gob"
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/event"
)

func init() {
	gob.Register(new(PushSkill))
}

type PushSkillEvent struct {
	attacker base.Character
	damage   float64
}

func (p PushSkillEvent) Type() Type {
	return DamageSkillType
}

func (p PushSkillEvent) Skill() Skill {
	return new(PushSkill)
}

func (p PushSkillEvent) Initiator() base.Character {
	return p.attacker
}

func (p PushSkillEvent) String() string {
	return fmt.Sprintf("%s使用了巨龙撞击！！造成了%.2f点伤害", p.attacker.Name(), p.damage)
}

type PushSkill struct {
}

func (p *PushSkill) Name() string {
	return "巨龙撞击"
}

func (p *PushSkill) Desc() string {
	return "对敌方造成80%力量的真实伤害"
}

func (p *PushSkill) Do(ctx *event.Context, cs ...base.Character) {
	attacker := cs[0]
	defender := cs[1]
	var use bool
	Backward(attacker.EventContext(), DamageSkillType, func(e EventSkill, step int) bool {
		if e.Skill().Name() == p.Name() {
			use = true
			return false
		}
		return true
	})
	if use {
		return
	}

	damage := attacker.Status().Power * 0.8

	defender.HP().Hit(damage)
	e := PushSkillEvent{
		attacker: attacker,
		damage:   damage,
	}
	ctx.AppendEvent(e)
	attacker.EventContext().AppendEvent(e)

}

func (p *PushSkill) String() string {
	return fmt.Sprintf("[%s]%s", p.Name(), p.Desc())
}
