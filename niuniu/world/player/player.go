package player

import (
	"github.com/kohmebot/niuniu/niuniu/world/buff"
	"time"
)

const (
	fatigue = 100
	rebirth = 1
)

type UID int64

func (u UID) IsValid() bool {
	return u != 0
}

type Player struct {
	// uid 一般是QQ号
	UID UID `gorm:"primaryKey"`
	// 从属的niuniu属性
	Niu *NiuNiu `gorm:"serializer:gob"`
	// 疲劳值
	Fatigue int
	// 转生次数
	Rebirth int
	// 上一次更新时间 (用于过了某个点更新)
	LastUpdated time.Time
}

// 更新状态
func (p *Player) update() {
	p.Fatigue = fatigue
	p.Rebirth = rebirth
	p.LastUpdated = time.Now()
}

func (p *Player) tryUpdate() {
	now := time.Now()
	var updateBufs []buff.Buff
	for _, b := range p.Niu.Buffs {
		if !now.After(b.DDL()) {
			updateBufs = append(updateBufs, b)
		}
	}
	p.Niu.Buffs = updateBufs

	// 获取今天的凌晨 4 点
	today4AM := time.Date(now.Year(), now.Month(), now.Day(), 4, 0, 0, 0, now.Location())
	// 如果 LastUpdated 在今天的凌晨 4 点之前，则更新
	if p.LastUpdated.Before(today4AM) {
		p.update()
	}
}
