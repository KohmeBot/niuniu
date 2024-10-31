package scene

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/entity"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"github.com/kohmebot/niuniu/niuniu/world/player"
)

type PKWinScene struct {
}

// NewPKWinScene PK胜利场景
func NewPKWinScene() *PKWinScene {
	return new(PKWinScene)
}

func (P *PKWinScene) Do(ctx *event.Context, es ...entity.Entity) {
	winner := es[0]
	loser := es[1]

	hpDelta := winner.HP().HP - loser.HP().HP
	// 剩余的HP
	// 越是惊险，给的越多
	point := winner.HP().MaxHP - hpDelta
	// 百分比
	percent := point * 0.01
	ne, ok := winner.(*player.NiuNiuEntity)
	if !ok {
		return
	}
	niu := ne.Snapshot()
	up := percent * niu.Length
	niu.Length += up

	ctx.AppendEvent(event.TextEvent(fmt.Sprintf("%s 变长了！(%.2f -> %.2f)", niu.Name, niu.Length-up, niu.Length)))

}

type PKLoseScene struct {
}

func NewPKLoseScene() *PKLoseScene {
	return new(PKLoseScene)
}

func (P *PKLoseScene) Do(ctx *event.Context, es ...entity.Entity) {
	winner := es[0]
	loser := es[1]

	_ = winner
	_ = loser
}
