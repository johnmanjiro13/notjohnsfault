package downcard

import "github.com/johnmanjiro13/notjohnsfault/game/card"

type Downcard struct {
	Cards card.Cards
}

func NewDowncard() *Downcard {
	return &Downcard{}
}

func (d Downcard) GetCards() card.Cards {
	return d.Cards
}

func (d Downcard) GetSum() int {
	sum := 0
	for _, c := range d.Cards {
		sum += c.GetNumber()
	}
	return sum
}

func (d *Downcard) Add(c card.ICard) {
	d.Cards = append(d.Cards, c)
}

func (d *Downcard) Remove() {
	d.Cards = nil
}

func (d Downcard) GetLength() int {
	return len(d.Cards)
}
