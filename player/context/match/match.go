package match

import (
	"github.com/johnmanjiro13/notjohnsfault/player/card"
	"github.com/johnmanjiro13/notjohnsfault/player/deck"
	"github.com/johnmanjiro13/notjohnsfault/player/discard"
	"github.com/johnmanjiro13/notjohnsfault/player/downcard"
)

type Match struct {
	Deck     deck.IDeck
	Downcard downcard.IDowncard
	Discard  discard.IDiscard
}

func (m *Match) Draw() card.ICard {
	if c := m.Deck.Remove(); c != nil {
		return c
	}
	return nil
}

func (m *Match) ToDowncard(useCard card.ICard) {
	m.Downcard.Add(useCard)
}

func (m *Match) DowncardToDiscard(deadCards downcard.Downcard) {
	for _, c := range deadCards.GetCards() {
		m.Discard.Add(c)
	}
	deadCards.Remove()
}

func (m *Match) DiscardToDeck(reuseCards discard.Discard) {
	for _, c := range reuseCards.GetCards() {
		m.Deck.Add(c)
	}
	reuseCards.Remove()
}
