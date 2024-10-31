package skill

import (
	"github.com/kohmebot/niuniu/niuniu/world/event"
)

// 技能类型

type Type int

const (
	// DamageSkillType 伤害形技能
	DamageSkillType = iota + 1
)

// HasType 判断事件中是否包含某技能类型
func HasType(ctx *event.Context, t Type) bool {
	var has bool
	for _, e := range event.ExtractEvent[EventSkill](ctx) {
		if e.Type() == t {
			has = true
			break
		}
	}
	return has
}

func Backward(ctx *event.Context, t Type, yield func(e EventSkill, step int) bool) {
	event.Backward[EventSkill](ctx, func(e EventSkill, step int) bool {
		if e.Type() == t {
			return yield(e, step)
		}
		return true
	})
}
