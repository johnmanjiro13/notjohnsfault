package field

import (
	"github.com/johnmanjiro13/notjohnsfault/player/discard"
	"github.com/johnmanjiro13/notjohnsfault/player/downcard"
)

type IField interface {
	DowncardToDiscard(deadCards downcard.Downcard)
	DiscardToDeck(reuseCards discard.Discard)
}
