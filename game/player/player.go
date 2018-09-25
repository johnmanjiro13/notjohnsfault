package player

import (
	"github.com/johnmanjiro13/notjohnsfault/game/card"
	"github.com/johnmanjiro13/notjohnsfault/game/deck"
	"github.com/johnmanjiro13/notjohnsfault/game/discard"
	"github.com/johnmanjiro13/notjohnsfault/game/downcard"
)

type Player struct {
	ID        PlayerID
	Deck      deck.IDeck
	Downcard  downcard.IDowncard
	Discard   discard.IDiscard
	Warned    bool
	Suspended bool
}

func NewPlayer(id PlayerID, dc deck.IDeck, dwc downcard.IDowncard) *Player {
	return &Player{
		ID:        id,
		Deck:      dc,
		Downcard:  dwc,
		Warned:    false,
		Suspended: false,
	}
}

func (p *Player) GetID() PlayerID {
	return p.ID
}

func (p *Player) Draw() card.ICard {
	if c := p.Deck.Remove(); c != nil {
		return c
	}
	return nil
}

func (p *Player) ToDowncard(useCard card.ICard) {
	p.Downcard.Add(useCard)
}

func (p *Player) UseYellowCard() {
	p.Warned = true
}

func (p *Player) ResetWarn() {
	p.Warned = false
}

func (p *Player) SetSuspend() {
	p.Suspended = true
}

func (p *Player) ResetSuspend() {
	p.Suspended = false
}
