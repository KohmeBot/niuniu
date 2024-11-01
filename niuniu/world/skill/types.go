package skill

// Type 表示技能类型
type Type int

const (
	// Damage 伤害型技能
	Damage Type = iota + 1
	// Heal 回复型技能
	Heal
	// Buff 增益型技能
	Buff
	// Support 辅助型技能
	Support
)
