package deck

import (
	"github.com/johnmanjiro13/notjohnsfault/player/card"
)

type IDeck interface {
	Add(c card.ICard)
	Remove() card.ICard
	GetLength() int
}
