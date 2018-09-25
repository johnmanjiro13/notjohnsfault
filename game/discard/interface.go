package discard

import "github.com/johnmanjiro13/notjohnsfault/game/card"

type IDiscard interface {
	GetCards() card.Cards
	Add(c card.ICard)
	Remove()
	GetLength() int
}
