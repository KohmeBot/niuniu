package stage

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/entity"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"github.com/kohmebot/niuniu/niuniu/world/player"
	"github.com/kohmebot/niuniu/niuniu/world/scene"
	"math/rand/v2"
	"strings"
)

type PkStage struct {
	PlayerA *player.Player
	PlayerB *player.Player
}

func NewPkStage(playerA *player.Player, playerB *player.Player) *PkStage {
	return &PkStage{
		PlayerA: playerA,
		PlayerB: playerB,
	}
}

func (p *PkStage) Do() base.Narration {
	var first, second *player.NiuNiu
	// 看谁先手
	if rand.Float64() > 0.5 {
		first = p.PlayerA.Niu
		second = p.PlayerB.Niu
	} else {
		first = p.PlayerB.Niu
		second = p.PlayerA.Niu
	}
	var builder strings.Builder
	builder.WriteString("开始击剑！\n")
	builder.WriteString(fmt.Sprintf("%s获得了先手机会！\n", first.Name))
	ctx := event.NewContext()
	// 生成实体
	firstE := first.Entity()
	secondE := second.Entity()
	var winner, loser entity.Entity
	for !ctx.IsDone() {
		p.round(ctx, firstE, secondE)
		if firstE.HP().HP <= 0 {
			winner = secondE
			loser = firstE
			ctx.Done(fmt.Sprintf("%s 胜利！", winner.Name()))
		} else if secondE.HP().HP <= 0 {
			winner = firstE
			loser = secondE
			ctx.Done(fmt.Sprintf("%s 胜利！", winner.Name()))
		}

		// 回合制，交换角色
		firstE, secondE = secondE, firstE
	}

	scene.NewPKWinScene().Do(ctx, winner, loser)
	scene.NewPKLoseScene().Do(ctx, winner, loser)

	for e, step := range ctx.RangeEvent {
		builder.WriteString(fmt.Sprintf("-----Event%d-----\n", step))
		builder.WriteString(e.String())
		builder.WriteByte('\n')
	}

	return &builder
}

// 进行一个回合
func (p *PkStage) round(ctx *event.Context, attacker entity.Entity, defender entity.Entity) {
	// 应用buff
	for _, buff := range attacker.Buffs() {
		buff.Do(ctx, attacker)
	}
	// 使用技能
	for _, skill := range attacker.Skills() {
		skill.Do(ctx, attacker, defender)
	}
	// 攻击！
	scene.NewAttackScene().Do(ctx, attacker, defender)
}
