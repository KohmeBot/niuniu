package buff

import (
	"encoding/gob"
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/base"
	"github.com/kohmebot/niuniu/niuniu/world/event"
	"time"
)

func init() {
	gob.Register(&PowerBuff{})
}

type PowerBuff struct {
	Ddl time.Time
}

func NewPowerBuff(ddl time.Time) *PowerBuff {
	return &PowerBuff{Ddl: ddl}
}

func (p *PowerBuff) Name() string {
	return "牛子力up"
}

func (p *PowerBuff) Desc() string {
	return "在战斗开始时，获得10%的力量"
}

func (p *PowerBuff) Do(ctx *event.Context, c base.Character) {
	if !c.EventContext().IsFirst() {
		return
	}
	s := c.Status()
	s.Power += s.Power * 0.1
	ctx.AppendEvent(event.TextEvent(fmt.Sprintf("%s 获得了10%%的力量", c.Name())))
}

func (p *PowerBuff) String() string {
	dur := p.Ddl.Sub(time.Now())
	return fmt.Sprintf("[%s]%s %s后消失", p.Name(), p.Desc(), dur.String())
}

func (p *PowerBuff) DDL() time.Time {
	return p.Ddl
}
