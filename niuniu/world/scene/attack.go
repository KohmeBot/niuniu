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
	attacker entity.Entity
	// 被攻击者
	defender entity.Entity
	// 造成的伤害
	damage float64
	// 攻击者是否暴击
	isCritical bool
	// 攻击者是否闪避
	isDodge bool
}

func (e AttackEvent) String() string {
	var builder strings.Builder
	if e.isCritical {
		builder.WriteString(fmt.Sprintf("%s暴击了！\n", e.attacker.Name()))
	}
	if e.isDodge {
		builder.WriteString(fmt.Sprintf(" %s成功闪避！\n", e.defender.Name()))
	}
	builder.WriteString(fmt.Sprintf("%s对%s造成了%.0f点伤害", e.attacker.Name(), e.defender.Name(), e.damage))
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
	if s.isAttackBefore(eA.EventContext()) {
		return
	}
	e := AttackEvent{
		attacker: eA,
		defender: eB,
	}
	// 计算本次A的伤害
	damage := s.damage(eA)

	if e.isCritical = sA.HitCritical(); e.isCritical {
		// 暴击了
		damage = sA.HitCriticalDamage(damage)
	}

	if e.isDodge = sB.HitDodge(); e.isDodge {
		// 闪避了，实际伤害减为0
		damage = 0
	}
	damage = sB.HitDefense(damage)
	eB.HP().Hit(damage)
	e.damage = damage
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

// 返回本轮前有没有造成伤害
func (s *AttackScene) isAttackBefore(ctx *event.Context) bool {
	var damageSkillStep, attackEventStep int
	skill.Backward(ctx, skill.DamageSkillType, func(e skill.EventSkill, step int) bool {
		damageSkillStep = step
		return false
	})
	event.Backward[AttackEvent](ctx, func(e AttackEvent, step int) bool {
		attackEventStep = step
		return false
	})
	// 表示 上一次攻击之后 触发了伤害技能
	return attackEventStep < damageSkillStep
}
