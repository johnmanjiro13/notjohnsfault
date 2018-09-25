package deck

import (
	"github.com/johnmanjiro13/notjohnsfault/player/card"
)

type Deck struct {
	Cards card.Cards
}

func NewDeck(c []card.ICard) *Deck {
	return &Deck{
		Cards: c,
	}
}

func (d Deck) Add(c card.ICard) {
	d.Cards = append(d.Cards, c)
}

func (d *Deck) Remove() card.ICard {
	if len(d.Cards) <= 0 {
		return nil
	}
	c := d.Cards[0]
	d.Cards = d.Cards[1:]
	return c
}

func (d Deck) GetLength() int {
	return len(d.Cards)
}
