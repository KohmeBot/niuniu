package player

import (
	"errors"
	"gorm.io/gorm"
)

type Generator struct {
	db *gorm.DB
}

func NewGenerator(db *gorm.DB) (*Generator, error) {
	g := &Generator{db: db}
	return g, autoMigrate(g.db)
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Player{})
}

func (g *Generator) NewPlayer(qq int64) (*Player, error) {
	p := &Player{UID: UID(qq)}
	if err := g.db.First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			p.Niu = new(NiuNiu)
			// 第一次，插入记录
			p.update()
			return p, g.db.Create(p).Error
		}
		return nil, err
	}
	// 不是第一次, 尝试更新
	p.tryUpdate()
	return p, nil
}

func (g *Generator) SavePlayer(ps ...*Player) error {
	err := g.db.Transaction(func(tx *gorm.DB) error {
		for _, p := range ps {
			if err := tx.Save(p).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
