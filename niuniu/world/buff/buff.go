package buff

import (
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"time"
)

type Buff interface {
	// Name 获取buff名称
	Name() string
	// Desc 获取buff描述
	Desc() string
	// Do 启动buff
	Do(ctx *event.Context, c base.Character)
	String() string
	// DDL 到期时间
	DDL() time.Time
}
