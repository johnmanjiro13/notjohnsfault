package discard

import "github.com/johnmanjiro13/notjohnsfault/game/card"

type Discard struct {
	Cards card.Cards
}

func NewDiscard() *Discard {
	return &Discard{}
}

func (d Discard) GetCards() card.Cards {
	return d.Cards
}

func (d *Discard) Add(c card.ICard) {
	d.Cards = append(d.Cards, c)
}

func (d *Discard) Remove() {
	d.Cards = nil
}

func (d Discard) GetLength() int {
	return len(d.Cards)
}
