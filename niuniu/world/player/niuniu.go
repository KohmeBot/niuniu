package player

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/buff"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"github.com/kohmebot/niuniu/niuniu/world/skill"
	"math"
	"math/rand/v2"
	"strings"
)

type NiuNiu struct {
	// 名称
	Name string
	// 所有者
	Owner UID
	// 长度
	Length float64
	// 硬度
	Hardness float64
	// 弹性
	Flexibility float64
	// 幸运值
	Luck float64
	// 拥有的buff
	Buffs []buff.Buff
	// 拥有的技能
	Skills []skill.Skill
}

func (n *NiuNiu) IsValid() bool {
	return n.Owner.IsValid()
}

// RandNiuniu 随机生成一个牛牛
func RandNiuniu(owner int64, name string) *NiuNiu {
	c := float64(digitCount(owner))
	// c 是一个基础值
	length := c + randomFloat(5, 8)
	hardness := c + randomFloat(60, 80)
	flex := c + randomFloat(5, 20)
	luck := c + randomFloat(1, 20)

	return &NiuNiu{
		Name:        name,
		Length:      length,
		Hardness:    hardness,
		Flexibility: flex,
		Luck:        luck,
		Owner:       UID(owner),
	}
}

// Entity 生成实体
func (n *NiuNiu) Entity() *NiuNiuEntity {
	return &NiuNiuEntity{
		ctx:      event.NewContext(),
		snapshot: n,
		uid:      int64(n.Owner),
		name:     n.Name,
		hp: base.HP{
			HP:    100,
			MaxHP: 100,
		},
		status: base.Status{
			Critical:       0.5,
			CriticalDamage: 1.5,
			Dodge:          0.05,
			Luck:           n.Luck,
			Power:          n.Length + 0.3*n.Hardness,
			Defense:        0.2 * n.Hardness,
		},
		skills: n.Skills,
		buffs:  n.Buffs,
	}
}

func (n *NiuNiu) String() string {
	var builder strings.Builder
	builder.WriteString("-----牛牛属性-----\n")
	builder.WriteString(fmt.Sprintf("名称：%s\n长度：%.2f\n硬度：%.2f\n弹性：%.2f\n幸运值：%.2f\n",
		n.Name, n.Length, n.Hardness, n.Flexibility, n.Luck))
	builder.WriteString("-----拥有的技能-----\n")
	for _, s := range n.Skills {
		builder.WriteString(s.String())
		builder.WriteByte('\n')
	}
	builder.WriteString("-----存续的Buff-----\n")
	for _, b := range n.Buffs {
		builder.WriteString(b.String())
		builder.WriteByte('\n')
	}
	return builder.String()

}

type NiuNiuEntity struct {
	// 这里表明是从哪个快照生成的
	snapshot *NiuNiu

	ctx    *event.Context
	uid    int64
	name   string
	hp     base.HP
	status base.Status
	buffs  []buff.Buff
	skills []skill.Skill
}

func (n *NiuNiuEntity) Name() string {
	return n.name
}

func (n *NiuNiuEntity) ID() int64 {
	return n.uid
}

func (n *NiuNiuEntity) HP() *base.HP {
	return &n.hp
}

func (n *NiuNiuEntity) Status() *base.Status {
	return &n.status
}

func (n *NiuNiuEntity) Buffs() []buff.Buff {
	return n.buffs
}

func (n *NiuNiuEntity) Skills() []skill.Skill {
	return n.skills
}

func (n *NiuNiuEntity) EventContext() *event.Context {
	return n.ctx
}

// Snapshot 获得快照
func (n *NiuNiuEntity) Snapshot() *NiuNiu {
	return n.snapshot
}

// 生成指定区间内的随机整数 [min, max]
func randomInt(min, max int) int {
	return rand.IntN(max-min+1) + min
}

// 生成指定区间内的随机浮点数 [min, max]
func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// 获取数字位数
func digitCount(n int64) int {
	if n == 0 {
		return 1
	}
	return int(math.Log10(math.Abs(float64(n)))) + 1
}
