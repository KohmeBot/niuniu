package base

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"io"
	"strings"
)

// Narration 念白，旁白
type Narration interface {
	String() string
}

type str string

func (s str) String() string {
	return string(s)
}

func EventNarration(ctx *event.Context) Narration {
	var builder strings.Builder
	for e, step := range ctx.RangeEvent {
		builder.WriteString(fmt.Sprintf("-----Event%d-----\n", step))
		builder.WriteString(e.String())
		builder.WriteByte('\n')
	}
	return str(builder.String())
}

func WriteNarration(wr io.StringWriter, na Narration) {
	_, _ = wr.WriteString(na.String())
}

func JoinNarration(s string, nas ...Narration) Narration {
	var builder strings.Builder
	builder.WriteString(s)
	for _, na := range nas {
		builder.WriteString(na.String())
	}
	return str(builder.String())
}
