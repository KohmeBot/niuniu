package world

import (
	"github.com/kohmebot/niuniu/niuniu/world/player"
	"github.com/kohmebot/niuniu/niuniu/world/trigger"
	"github.com/kohmebot/plugin"
	zero "github.com/wdvxdr1123/ZeroBot"
)

type World struct {
	env       plugin.Env
	generator *player.Generator
}

func NewWorld(env plugin.Env) (w *World, err error) {
	w = &World{env: env}
	db, err := env.GetDB()
	if err != nil {
		return
	}
	w.generator, err = player.NewGenerator(db)
	if err != nil {
		return
	}

	return
}

func (w *World) InitWorld(engine *zero.Engine) {
	trigger.Register(engine, w.generator, w.env)
}
