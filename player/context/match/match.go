package match

import (
	"github.com/johnmanjiro13/notjohnsfault/player/card"
	"github.com/johnmanjiro13/notjohnsfault/player/deck"
	"github.com/johnmanjiro13/notjohnsfault/player/discard"
	"github.com/johnmanjiro13/notjohnsfault/player/downcard"
)

type Match struct {
	ID       PlayerID
	Deck     deck.IDeck
	Downcard downcard.Downcard
	Discard  discard.Discard
}

func (m *Match) GetID() PlayerID {
	return m.ID
}

func (m *Match) Draw() card.ICard {
	if c := deck.Remove; c != nil {
		return deck.Remove()
	}
	return nil
}

func (m *Match) ToDowncard(useCard card.ICard) {
	downcard.Add(useCard)
}

func (m *Match) DowncardToDiscard(deadCards downcard) {
	cs := deadCards.GetCards()
	discard.Add(cs)
	deadCards.Remove()
}

func (m *Match) DiscardToDeck(reuseCards discard) {
	cs := reuseCards.GetCards()
	deck.Add(cs)
	reuseCards.Remove()
}
