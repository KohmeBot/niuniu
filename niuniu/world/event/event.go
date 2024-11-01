package event

import "strings"

type Event interface {
	String() string
}

type SpecialEvent interface {
	Event
	EventName() string
}

// EndEvent 结束事件
type EndEvent string

func (e EndEvent) String() string {
	return string(e)
}

// TextEvent 文本事件，仅作为string，没有任何实际的作用
type TextEvent string

func (e TextEvent) String() string {
	return strings.TrimSpace(string(e))
}
