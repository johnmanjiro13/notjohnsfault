package match

import (
	"github.com/johnmanjiro13/notjohnsfault/player/deck"
	"github.com/johnmanjiro13/notjohnsfault/player/discard"
	"github.com/johnmanjiro13/notjohnsfault/player/downcard"
)

type Field struct {
	Match
}

func NewField(dc deck.IDeck, dwc downcard.IDowncard, dic discard.IDiscard) *Field {
	return &Field{
		Match: Match{
			Deck:     dc,
			Downcard: dwc,
			Discard:  dic,
		},
	}
}
