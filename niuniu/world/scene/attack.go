package scene

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/entity"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"github.com/kohmebot/niuniu/niuniu/world/skill"
	"strings"
)

type AttackScene struct {
}

// AttackEvent 攻击事件
type AttackEvent struct {
	// 攻击者
	Attacker entity.Entity
	// 被攻击者
	Defender entity.Entity
	// 造成的伤害
	Damage float64
	// 攻击者是否暴击
	IsCritical bool
	// 攻击者是否闪避
	IsDodge bool
}

func (e AttackEvent) String() string {
	var builder strings.Builder
	if e.IsCritical {
		builder.WriteString(fmt.Sprintf("%s暴击了！\n", e.Attacker.Name()))
	}
	if e.IsDodge {
		builder.WriteString(fmt.Sprintf(" %s成功闪避！\n", e.Defender.Name()))
	}
	builder.WriteString(fmt.Sprintf("%s对%s造成了%.0f点伤害", e.Attacker.Name(), e.Defender.Name(), e.Damage))
	return builder.String()
}

// NewAttackScene 生成一个攻击场景
func NewAttackScene() *AttackScene {
	return new(AttackScene)
}

// Do 一般来说，是A攻击B
func (s *AttackScene) Do(ctx *event.Context, entities ...entity.Entity) {

	eA, eB := entities[0], entities[1]
	sA, sB := eA.Status(), eB.Status()
	if s.isAttackBefore(eA) {
		return
	}
	e := AttackEvent{
		Attacker: eA,
		Defender: eB,
	}
	// 计算本次A的伤害
	damage := s.damage(eA)

	if e.IsCritical = sA.HitCritical(); e.IsCritical {
		// 暴击了
		damage = sA.HitCriticalDamage(damage)
	}

	if e.IsDodge = sB.HitDodge(); e.IsDodge {
		// 闪避了，实际伤害减为0
		damage = 0
	}

	damage = sB.HitDefense(damage)

	eB.HP().Hit(damage)
	e.Damage = damage
	ctx.AppendEvent(e)
	eA.EventContext().AppendEvent(e)
}

// 计算伤害
func (s *AttackScene) damage(e entity.Entity) float64 {
	status := e.Status()
	// 伤害计算
	damage := status.Power

	return damage
}

// 返回某人本轮前有没有造成伤害
func (s *AttackScene) isAttackBefore(e entity.Entity) bool {
	sk, _, ok := event.IsAfter[skill.EventSkill, AttackEvent](e.EventContext())
	// 代表上一次攻击。之后，有伤害型技能触发
	if ok && sk.Type() == skill.Damage {
		return true
	}
	return false
}
