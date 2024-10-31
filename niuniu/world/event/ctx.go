package event

import (
	"slices"
	"strings"
)

// Context 事件上下文
type Context struct {
	history []Event
}

func NewContext() *Context {
	return &Context{}
}

// AppendEvent 添加事件
func (c *Context) AppendEvent(e Event) {
	c.history = append(c.history, e)
}

// LastEvent 返回最后一个事件
func (c *Context) LastEvent() Event {
	return c.GetEvent(c.Step())
}

// IsFirst 是否处于第一个事件
func (c *Context) IsFirst() bool {
	return c.Step() == 1
}

// Step 目前处于第几个事件
func (c *Context) Step() int {
	return len(c.history) + 1
}

// GetEvent 获取某一个事件
func (c *Context) GetEvent(step int) Event {
	if c.IsFirst() {
		return nil
	}
	return c.history[step-2]
}

// IsDone 判断是否已经结束
func (c *Context) IsDone() bool {
	_, ok := c.LastEvent().(EndEvent)
	return ok
}

// Done 结束上下文
func (c *Context) Done(s ...string) {
	c.AppendEvent(EndEvent(strings.Join(s, " ")))
}

func (c *Context) RangeEvent(yield func(e Event, step int) bool) {
	for i, event := range c.history {
		if !yield(event, i+1) {
			return
		}
	}
}

// ExtractEvent 提取事件
func ExtractEvent[T Event](ctx *Context) []T {
	var res []T
	for _, event := range ctx.history {
		if e, ok := event.(T); ok {
			res = append(res, e)
		}
	}
	return res
}

// Backward 向前回溯某个事件
func Backward[T Event](ctx *Context, yield func(e T, step int) bool) {
	for i, event := range slices.Backward(ctx.history) {
		if e, ok := event.(T); ok {
			if !yield(e, i+1) {
				return
			}
		}
	}

}
