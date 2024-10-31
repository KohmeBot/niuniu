package trigger

import (
	"github.com/kohmebot/niuniu/niuniu/world/player"
	"github.com/kohmebot/niuniu/niuniu/world/stage"
	"github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strconv"
)

type PKTrigger struct {
}

func (P *PKTrigger) TriggerText() []string {
	return []string{"击剑"}
}

func (P *PKTrigger) On(ctx *zero.Ctx, g *player.Generator) (err error) {
	var sender, target int64
	sender = ctx.Event.UserID
	for _, segment := range ctx.Event.Message {
		if segment.Type == "at" {
			v := segment.Data["qq"]
			target, err = strconv.ParseInt(v, 10, 64)
			if err != nil {
				return
			}
			break
		}
	}
	if target == 0 {
		return nil
	}
	senderPlayer, err := g.NewPlayer(sender)
	if err != nil {
		return err
	}
	targetPlayer, err := g.NewPlayer(target)
	if err != nil {
		return err
	}
	if !senderPlayer.Niu.IsValid() {
		// 让它先创建
		ctx.Send(message.Text(niuniuNotFound))
		return nil
	}
	if !targetPlayer.Niu.IsValid() {
		// 指对方没有角色
		ctx.Send(message.Text(targetNiuniuNotFound))
		return nil
	}
	s := stage.NewPkStage(senderPlayer, targetPlayer)
	msg := s.Do()

	err = g.SavePlayer(senderPlayer, targetPlayer)
	if err != nil {
		return err
	}
	if code := ctx.Send(message.Text(msg.String())); code.ID() == 0 {
		return nil
	}
	return nil
}
