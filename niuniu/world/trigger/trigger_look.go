package trigger

import (
	"github.com/kohmebot/niuniu/niuniu/world/player"
	"github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type LookTrigger struct {
}

func (l *LookTrigger) TriggerText() []string {
	return []string{"看看牛牛"}
}

func (l *LookTrigger) On(ctx *zero.Ctx, g *player.Generator) error {
	sender := ctx.Event.UserID

	p, err := g.NewPlayer(sender)
	if err != nil {
		return err
	}

	if !p.Niu.IsValid() {
		ctx.Send(message.Text(`你没有牛牛，发送"领养牛牛 名称"来领养一只可爱的牛牛吧！"`))
		return nil
	}

	ctx.Send(message.Text(p.Niu.String()))
	return nil
}
