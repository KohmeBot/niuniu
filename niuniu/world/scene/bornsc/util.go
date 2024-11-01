package bornsc

import (
	"github.com/kohmebot/niuniu/niuniu/world/entity"
	"github.com/kohmebot/niuniu/niuniu/world/player"
)

func ExtractNiuNiu(e entity.Entity) *player.NiuNiu {
	return e.(*player.NiuNiuEntity).Snapshot()
}
