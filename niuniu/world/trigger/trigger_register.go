package trigger

import (
	"fmt"
	"github.com/kohmebot/niuniu/niuniu/world/player"
	"github.com/kohmebot/niuniu/niuniu/world/stage"
	"github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strings"
)

type RegisterTrigger struct {
}

func (r *RegisterTrigger) TriggerText() []string {
	return []string{"领养牛牛"}
}

func (r *RegisterTrigger) On(ctx *zero.Ctx, g *player.Generator) error {
	cs := ParseArgs(ctx)
	if len(cs) < 2 {
		return nil
	}
	niuniuName := cs[1]

	sender := ctx.Event.UserID
	p, err := g.NewPlayer(sender)
	if err != nil {
		return err
	}
	if p.Niu.IsValid() {
		ctx.Send(message.Text("你已经有牛牛了噢！"))
		return nil
	}

	niu := player.RandNiuniu(sender, niuniuName)
	p.Niu = niu

	s := stage.NewBornStage(niu)

	msg := s.Do()

	err = g.SavePlayer(p)
	if err != nil {
		return err
	}
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s 出生了\n", niu.Name))
	builder.WriteString(msg.String())
	ctx.Send(message.Text(builder.String()))
	ctx.Send(message.Text(niu.String()))

	return nil
}
