package base

import (
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"math/rand/v2"
)

type Character interface {
	ID() int64
	Name() string
	HP() *HP
	Status() *Status
	EventContext() *event.Context
}

type HP struct {
	// 最大生命值
	MaxHP float64
	// 现生命值
	HP float64
}

// Hit 攻击
func (h *HP) Hit(damage float64) {
	h.HP -= damage
}

type Status struct {
	// 攻击力
	Power float64
	// 闪避概率
	Dodge float64
	// 防御力
	Defense float64
	// 幸运值
	Luck float64
	// 暴击率
	Critical float64
	// 暴击伤害倍率
	CriticalDamage float64
}

// HitDodge 判定是否闪避
func (s *Status) HitDodge() bool {
	// TODO 可以根据幸运值加成
	return rand.Float64() < s.Dodge
}

// HitCritical 判定是否暴击
func (s *Status) HitCritical() bool {
	// TODO 可以根据幸运值加成
	return rand.Float64() < s.Critical
}

// HitCriticalDamage 判定暴击伤害
func (s *Status) HitCriticalDamage(damage float64) float64 {
	return damage * s.CriticalDamage
}

// HitDefense 伤害减免/判定防御
func (s *Status) HitDefense(damage float64) float64 {
	// TODO 可以根据幸运值加成
	return damage - s.Defense
}
