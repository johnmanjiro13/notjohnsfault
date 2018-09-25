package deck

import (
	"math/rand"
	"time"

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
	rand.Seed(time.Now().UnixNano())
	c := d.Cards[rand.Intn(len(d.Cards))]
	d.Cards = d.Cards[1:]
	return c
}

func (d Deck) GetLength() int {
	return len(d.Cards)
}
