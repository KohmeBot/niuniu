package event

import (
	"slices"
	"strings"
)

// Context 事件上下文
type Context struct {
	history []Event
	futures []*Future
}

func NewContext() *Context {
	return &Context{}
}

// AppendEvent 添加事件
func (c *Context) AppendEvent(e Event) {
	c.history = append(c.history, e)
	var newEvents []Event
	c.futures = slices.DeleteFunc(c.futures, func(future *Future) bool {
		if !future.is(e) {
			return false
		}
		newEv, d := future.handle(e, c.Step()-1)
		if newEv != nil {
			newEvents = append(newEvents, newEv)
		}
		return d
	})
	for _, ev := range newEvents {
		c.AppendEvent(ev)
	}
}

// Future 创建一个未来事件
func (c *Context) Future() *Future {
	f := new(Future)
	f.times = 1
	c.SetFuture(f)
	return f
}

func (c *Context) SetFuture(f *Future) {
	c.futures = append(c.futures, f)
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

// BackwardN 向前回溯某个事件直到N次为止，返回找到的次数
func BackwardN[T Event](ctx *Context, n int, yield func(e T, step int)) int {
	c := 0
	Backward[T](ctx, func(e T, step int) bool {
		c++
		yield(e, step)
		return c == n
	})
	return c
}

// BackwardFirst 向前回溯某个事件，返回第一个找到的
func BackwardFirst[T Event](ctx *Context) (e T, step int, ok bool) {
	BackwardN[T](ctx, 1, func(ev T, st int) {
		e = ev
		step = st
		ok = true
	})
	return
}

// IsBefore 判断A事件是否在B事件之前，返回找到的A事件和B事件实例，以及是否在之前
func IsBefore[A Event, B Event](ctx *Context) (eA A, eB B, ok bool) {
	eA, aStep, _ := BackwardFirst[A](ctx)
	eB, bStep, _ := BackwardFirst[B](ctx)
	return eA, eB, aStep < bStep
}

// IsAfter 判断A事件是否在B事件之后, 返回找到的A事件和B事件实例，以及是否在之后
func IsAfter[A Event, B Event](ctx *Context) (eA A, eB B, ok bool) {
	eA, aStep, _ := BackwardFirst[A](ctx)
	eB, bStep, _ := BackwardFirst[B](ctx)
	return eA, eB, aStep > bStep
}

// FutureWhenOn 创建一个未来事件，当匹配到指定类型时，执行回调
func FutureWhenOn[T Event](ctx *Context, h func(e T, step int)) *Future {
	future := ctx.Future()
	future.When(For[T]()).On(func(e Event, step int) {
		if v, ok := e.(T); ok {
			h(v, step)
		}
	})
	return future
}

func FutureWhenDoAndDelete[T Event](ctx *Context, h func(e T, step int) (Event, bool)) *Future {
	future := ctx.Future()
	future.When(For[T]()).DoAndDelete(func(e Event, step int) (Event, bool) {
		return h(e.(T), step)
	})
	return future
}
