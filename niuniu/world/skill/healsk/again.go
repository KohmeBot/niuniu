package healsk

import (
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"github.com/kohmebot/niuniu/niuniu/world/skill"
	"github.com/kohmebot/niuniu/niuniu/world/util"
)

func init() {
	util.RegGob[Again]()
}

type Again struct {
}

func (a *Again) Name() string {
	return "再度雄起"
}

func (a *Again) Desc() string {
	return "当HP归0时，有60%的几率以1点血量复活"
}

func (a *Again) Do(ctx *event.Context, cs ...base.Character) {
	if _, _, use := skill.IsUse[*Again](ctx); use {
		return
	}
	target := cs[0]
	ctx.Future().AnyTimes().DoAndDelete(func(event.Event, int) (event.Event, bool) {
		hp := target.HP()
		if hp.HP > 0 {
			return nil, false
		}
		hp.HP = 1
		e := skill.NewEventSkill[*Again](target, skill.Heal, "以一滴血雄起了！")
		target.EventContext().AppendEvent(e)
		return e, true
	})
}

func (a *Again) String() string {
	return skill.CommonFormat(a)
}
