package skill

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/event"
)

func CommonFormat(s Skill) string {
	return fmt.Sprintf("[%s]%s", s.Name(), s.Desc())
}

type Skill interface {
	Name() string
	Desc() string
	Do(ctx *event.Context, cs ...base.Character)
	String() string
}

// IsUse 判断之前最近是否用过了某个技能,返回对应的技能事件
func IsUse[T Skill](ctx *event.Context) (e EventSkill, step int, ok bool) {
	s := *new(T)
	event.Backward[EventSkill](ctx, func(e2 EventSkill, step2 int) bool {
		if e2.Skill().Name() == s.Name() {
			e = e2
			step = step2
			ok = true
		}
		return !ok
	})
	return
}

// UseN 返回某个技能用了多少次
func UseN[T Skill](ctx *event.Context) int {
	s := *new(T)
	c := 0
	event.Backward[EventSkill](ctx, func(e EventSkill, step int) bool {
		if e.Skill().Name() == s.Name() {
			c++
		}
		return true
	})
	return c
}
