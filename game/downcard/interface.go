package downcard

import "github.com/johnmanjiro13/notjohnsfault/game/card"

type IDowncard interface {
	GetCards() card.Cards
	Add(c card.ICard)
	Remove()
	GetLength() int
	GetSum() int
}
