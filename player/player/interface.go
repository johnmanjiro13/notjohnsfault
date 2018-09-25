package player

import (
	"github.com/johnmanjiro13/notjohnsfault/player/card"
)

type IPlayer interface {
	GetID() PlayerID
	Draw() card.ICard
	ToDowncard(useCard card.ICard)
	useYellowCard()
	setSuspend()
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
