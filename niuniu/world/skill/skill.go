package skill

import (
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/event"
)

type Skill interface {
	Name() string
	Desc() string
	Do(ctx *event.Context, cs ...base.Character)
	String() string
}

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
