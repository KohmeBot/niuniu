package player

import "gorm.io/gorm"

type Generator struct {
	db *gorm.DB
}

func NewGenerator(db *gorm.DB) (*Generator, error) {
	g := &Generator{db: db}

	return g, nil
}

func (g *Generator) NewPlayer(qq int64) *Player {
	return new(Player)
}
