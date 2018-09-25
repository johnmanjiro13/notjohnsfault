package match

import (
	"github.com/johnmanjiro13/notjohnsfault/player/card"
)

type IPlayer interface {
	GetID() PlayerID
	Draw() card.ICard
	Play(id string) (card.ICard, error)
}

type Players []IPlayer

func (ps Players) FindByID(id PlayerID) IPlayer {
	for _, p := range ps {
		if p.GetID() == id {
			return p
		}
	}
	return nil
}
