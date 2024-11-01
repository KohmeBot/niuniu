package event

import (
	"reflect"
)

type Type reflect.Type

type Future struct {
	typ      Type
	on       func(e Event, step int) (Event, bool)
	times    int
	excludes []Type
}

// Times 指定执行次数
func (f *Future) Times(v int) *Future {
	f.times = v
	return f
}

// AnyTimes 任意次数
func (f *Future) AnyTimes() *Future {
	f.times = -1
	return f
}

// On 事件处理
func (f *Future) On(h func(e Event, step int)) *Future {
	f.on = func(e Event, step int) (Event, bool) {
		h(e, step)
		return nil, false
	}
	return f
}

// OnAndDelete 事件处理并返回是否删除
// Times 和 OnAndDelete 是||的关系
func (f *Future) OnAndDelete(h func(e Event, step int) bool) *Future {
	f.on = func(e Event, step int) (Event, bool) {
		return nil, h(e, step)
	}
	return f.AnyTimes()
}

// Do 事件处理,返回一个Event作为事件处理结果
func (f *Future) Do(h func(e Event, step int) Event) *Future {
	f.on = func(e Event, step int) (Event, bool) {
		return h(e, step), false
	}
	return f
}

// DoAndDelete 事件处理并返回一个Event作为事件处理结果,并返回是否删除Future
func (f *Future) DoAndDelete(h func(e Event, step int) (Event, bool)) *Future {
	f.on = h
	return f.AnyTimes()
}

// When 匹配事件类型
//
//	example: When(For[TextEvent]())
func (f *Future) When(typ Type) *Future {
	f.typ = typ
	return f
}

// Exclude 排除指定的事件类型
//
// example: Exclude(For[TextEvent]())
func (f *Future) Exclude(ts ...Type) *Future {
	f.excludes = append(f.excludes, ts...)
	return f
}

func (f *Future) is(e Event) bool {
	if f.typ == nil {
		return !f.isExclude(e)
	}
	switch f.typ.Kind() {
	case reflect.Interface:
		return reflect.TypeOf(e).Implements(f.typ)
	case reflect.Pointer:
		return reflect.TypeOf(e).Elem() == f.typ
	default:
		return reflect.TypeOf(e) == f.typ
	}
}

func (f *Future) isExclude(e Event) bool {
	if len(f.excludes) <= 0 {
		return false
	}
	typ := Type(reflect.TypeOf(e))
	for _, exclude := range f.excludes {
		switch f.typ.Kind() {
		case reflect.Interface:
			if typ.Implements(exclude) {
				return true
			}
		case reflect.Pointer:
			if typ.Elem() == exclude {
				return true
			}
		default:
			if typ == exclude {
				return true
			}
		}
	}
	return false
}

func (f *Future) handle(e Event, step int) (newEv Event, needDelete bool) {
	if f.on == nil {
		return nil, true
	}
	newEv, needDelete = f.on(e, step)
	if f.times == -1 {
		return
	}
	f.times--
	if f.times <= 0 {
		needDelete = true
	}
	return
}

// For 获取事件类型
func For[T Event]() Type {
	return reflect.TypeFor[T]()
}
