package scene

import (
	"github.com/kohmebot/niuniu/niuniu/world/entity"
	"github.com/kohmebot/niuniu/niuniu/world/event"
)

// Scene 场景
type Scene interface {
	// Do 进入场景，返回事件和场景描述
	Do(ctx *event.Context, entities ...entity.Entity)
}
