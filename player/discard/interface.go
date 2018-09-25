package discard

import "github.com/johnmanjiro13/notjohnsfault/player/card"

type IDiscard interface {
	Add(c card.ICard)
	Remove()
	GetLength() int
}
