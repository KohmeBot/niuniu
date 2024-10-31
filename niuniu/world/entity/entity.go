package entity

import (
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/buff"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"github.com/kohmebot/niuniu/niuniu/world/skill"
)

// Entity 实体
type Entity interface {
	ID() int64
	// Name 实体名称
	Name() string
	// HP 实体的生命值
	HP() *base.HP
	// Status 实体的各项状态
	Status() *base.Status
	// Buffs 实体拥有的buff
	Buffs() []buff.Buff
	// Skills 实体拥有的技能
	Skills() []skill.Skill
	// EventContext 获取该实体的事件上下文
	EventContext() *event.Context
}
