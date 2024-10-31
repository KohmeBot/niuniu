package main

import (
	"github.com/kohmebot/niuniu/niuniu"
	"github.com/kohmebot/plugin"
)

func NewPlugin() plugin.Plugin {
	return niuniu.NewPlugin()
}
