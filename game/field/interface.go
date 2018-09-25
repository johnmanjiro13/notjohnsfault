package field

import (
	"github.com/johnmanjiro13/notjohnsfault/game/discard"
	"github.com/johnmanjiro13/notjohnsfault/game/downcard"
)

type IField interface {
	DowncardToDiscard(deadCards downcard.Downcard)
	DiscardToDeck(reuseCards discard.Discard)
}
