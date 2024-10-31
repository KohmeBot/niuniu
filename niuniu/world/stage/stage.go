package stage

import "github.com/kohmebot/niuniu/niuniu/world/base"

type Stage interface {
	Do() base.Narration
}
