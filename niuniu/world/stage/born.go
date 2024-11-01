package stage

import (
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"github.com/kohmebot/niuniu/niuniu/world/player"
	"github.com/kohmebot/niuniu/niuniu/world/scene"
	"github.com/kohmebot/niuniu/niuniu/world/scene/bornsc"
	"github.com/kohmebot/niuniu/niuniu/world/util/prob"
)

var (
	bornGroup = prob.ProbabilityGroup[scene.Scene]{
		prob.Value[scene.Scene]{V: bornsc.NewPuTongScene(), Prob: 0.4},
		prob.Value[scene.Scene]{V: bornsc.NewNanNiangScene(), Prob: 0.2},
		prob.Value[scene.Scene]{V: bornsc.NewHaoZhanScene(), Prob: 0.2},
		prob.Value[scene.Scene]{V: bornsc.NewHeiRenScene(), Prob: 1000.2},
	}
)

type BornStage struct {
	niu *player.NiuNiu
}

func NewBornStage(niu *player.NiuNiu) *BornStage {
	return &BornStage{niu: niu}
}

func (b *BornStage) Do() base.Narration {
	s := bornGroup.Hit().V
	ctx := event.NewContext()
	s.Do(ctx, b.niu.Entity())
	return base.EventNarration(ctx)
}
