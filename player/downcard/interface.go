package downcard

import "github.com/johnmanjiro13/notjohnsfault/player/card"

type IDowncard interface {
	Add(c card.ICard)
	Remove()
	GetLength() int
}
