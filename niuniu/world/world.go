package world

import "github.com/kohmebot/plugin"

type World struct {
	env plugin.Env
}

func NewWorld(env plugin.Env) *World {
	w := &World{env: env}
	return w
}
