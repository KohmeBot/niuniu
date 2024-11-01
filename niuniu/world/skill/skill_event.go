package skill

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/event"
)

// EventSkill 技能事件
type EventSkill interface {
	event.Event
	// Type 获取技能类型
	Type() Type
	// Skill 具体的skill
	Skill() Skill
	// Initiator 发起者
	Initiator() base.Character
}

type baseEvent struct {
	initiator base.Character
	skill     Skill
	typ       Type
	desc      string
}

func (e *baseEvent) String() string {
	switch e.typ {
	case Damage, Heal, Support, Buff:
		return fmt.Sprintf("%s使用了%s\n%s", e.initiator.Name(), e.skill.Name(), e.desc)
	}
	return ""
}

func (e *baseEvent) Type() Type {
	return e.typ
}

func (e *baseEvent) Skill() Skill {
	return e.skill
}

func (e *baseEvent) Initiator() base.Character {
	return e.initiator
}

// NewEventSkill 创建一个技能事件
// desc: 本次事件发生了什么的描述
func NewEventSkill[T Skill](initiator base.Character, typ Type, desc string) EventSkill {
	return &baseEvent{
		initiator: initiator,
		skill:     *new(T),
		typ:       typ,
		desc:      desc,
	}
}
