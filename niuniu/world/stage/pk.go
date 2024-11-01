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
	// 生成实体
	aE := p.PlayerA.Niu.Entity()
	bE := p.PlayerB.Niu.Entity()
	var firstE, secondE entity.Entity
	// 看谁先手
	if rand.Float64() > 0.5 {
		firstE = aE
		secondE = bE
	} else {
		firstE = bE
		secondE = aE
	}
	var builder strings.Builder
	builder.WriteString("开始击剑！\n")
	builder.WriteString(fmt.Sprintf("%s获得了先手机会！\n", firstE.Name()))
	ctx := event.NewContext()

	var winner, loser entity.Entity
	for !ctx.IsDone() {
		// 每回合计算HP
		beforeAhp := aE.HP().HP
		beforeBhp := bE.HP().HP
		p.round(ctx, firstE, secondE)
		afterAhp := aE.HP().HP
		afterBhp := bE.HP().HP
		var b strings.Builder
		if beforeAhp != afterAhp {
			b.WriteString(fmt.Sprintf("%s:%.0f -> %.0f\n", aE.Name(), beforeAhp, afterAhp))
		}
		if beforeBhp != afterBhp {
			b.WriteString(fmt.Sprintf("%s:%.0f -> %.0f\n", bE.Name(), beforeBhp, afterBhp))
		}
		if b.Len() > 0 {
			ctx.AppendEvent(event.TextEvent(b.String()))
		}

		if firstE.HP().HP <= 0 {
			winner = secondE
			loser = firstE
			ctx.Done(fmt.Sprintf("%s 胜利！", winner.Name()))
		} else if secondE.HP().HP <= 0 {
			winner = firstE
			loser = secondE
			ctx.Done(fmt.Sprintf("%s 胜利！", winner.Name()))
		}
		// TODO 平局
		// 回合制，交换角色
		firstE, secondE = secondE, firstE
	}

	// 只结算发起者
	if winner.ID() == int64(p.PlayerA.UID) {
		scene.NewPKWinScene().Do(ctx, winner, loser)
	}
	if loser.ID() == int64(p.PlayerA.UID) {
		scene.NewPKLoseScene().Do(ctx, winner, loser)
	}

	return base.JoinNarration(builder.String(), base.EventNarration(ctx))
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
