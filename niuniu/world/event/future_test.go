package event

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestFuture_Times(t *testing.T) {
	ctx := NewContext()

	times := 10
	trigger := 0
	ctx.Future().Times(times).On(func(e Event, step int) {
		trigger++
	})

	for i := 0; i < times*3; i++ {
		ctx.AppendEvent(&strings.Builder{})
	}

	assert.Equal(t, times, trigger)

}

func TestFuture_AnyTimes(t *testing.T) {
	ctx := NewContext()

	trigger := 0
	loop := 0

	ctx.Future().AnyTimes().On(func(e Event, step int) {
		trigger++
	})

	after := time.After(time.Second)
	b := new(strings.Builder)
LP:
	for {
		select {
		case <-after:
			break LP
		default:
		}
		loop++
		ctx.AppendEvent(b)
	}
	t.Log(loop)
	assert.Equal(t, loop, trigger)
}

func TestFuture_When(t *testing.T) {
	ctx := NewContext()

	trueStep := 3
	trigger := 0
	ev := EndEvent("end")
	ctx.Future().When(For[EndEvent]()).On(func(e Event, step int) {
		trigger++
		assert.Equal(t, trueStep, step)
		assert.Equal(t, ev, e)
	})

	for ctx.Step() < trueStep {
		ctx.AppendEvent(&strings.Builder{})
	}
	ctx.AppendEvent(ev)
	assert.Equal(t, 1, trigger)

}

func TestFutureWhen(t *testing.T) {
	ctx := NewContext()

	trueStep := 3
	trigger := 0
	ev := EndEvent("end")
	FutureWhenOn[EndEvent](ctx, func(e EndEvent, step int) {
		trigger++
		assert.Equal(t, trueStep, step)
		assert.Equal(t, ev, e)
	})

	for ctx.Step() < trueStep {
		ctx.AppendEvent(&strings.Builder{})
	}
	ctx.AppendEvent(ev)
	assert.Equal(t, 1, trigger)

}

func TestFuture_On(t *testing.T) {
	ctx := NewContext()

	count := 10
	trigger := 0
	ctx.Future().On(func(e Event, step int) {
		trigger++
	}).AnyTimes()

	for i := 0; i < count; i++ {
		ctx.AppendEvent(&strings.Builder{})
	}
	assert.Equal(t, count, trigger)
}
