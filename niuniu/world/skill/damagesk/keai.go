package damagesk

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"github.com/kohmebot/niuniu/niuniu/world/scene"
	"github.com/kohmebot/niuniu/niuniu/world/skill"
	"github.com/kohmebot/niuniu/niuniu/world/util"
	"github.com/kohmebot/niuniu/niuniu/world/util/prob"
)

func init() {
	util.RegGob[KeAi]()
}

type KeAi struct {
}

func (k *KeAi) Name() string {
	return "小小的也很可爱"
}

func (k *KeAi) Desc() string {
	return "60%几率发动，全额返还对方的普攻伤害"
}

func (k *KeAi) Do(ctx *event.Context, cs ...base.Character) {
	self := cs[0]
	target := cs[1]
	_, _, use := skill.IsUse[*KeAi](self.EventContext())
	if use || !prob.HitProb(0.6) {
		return
	}

	// 0.6概率且之前没法动
	event.FutureWhenOn[scene.AttackEvent](target.EventContext(), func(e scene.AttackEvent, _ int) {
		target.HP().Hit(e.Damage)
		ev := skill.NewEventSkill[*KeAi](self, skill.Damage, fmt.Sprintf("反弹了%.0f点伤害！", e.Damage))
		ctx.AppendEvent(ev)
		self.EventContext().AppendEvent(ev)
	})

}

func (k *KeAi) String() string {
	return skill.CommonFormat(k)
}
