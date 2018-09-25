package field

import (
	"github.com/johnmanjiro13/notjohnsfault/game/deck"
	"github.com/johnmanjiro13/notjohnsfault/game/discard"
	"github.com/johnmanjiro13/notjohnsfault/game/downcard"
	"github.com/johnmanjiro13/notjohnsfault/game/player"
)

type Field struct {
	Deck          deck.IDeck
	Downcard      downcard.IDowncard
	Discard       discard.IDiscard
	CurrentPlayer player.Player
	NextPlayer    player.Player
	OppPlayer     player.Player
	LastPlayer    player.Player
}

func NewField(dc deck.IDeck, dwc downcard.IDowncard, dic discard.IDiscard,
	p1 player.Player, p2 player.Player, p3 player.Player, p4 player.Player) *Field {
	return &Field{
		Deck:          dc,
		Downcard:      dwc,
		Discard:       dic,
		CurrentPlayer: p1,
		NextPlayer:    p2,
		OppPlayer:     p3,
		LastPlayer:    p4,
	}
}

func (f *Field) DowncardToDiscard() {
	for _, c := range f.Downcard.GetCards() {
		f.Discard.Add(c)
	}
	f.Downcard.Remove()
}

func (f *Field) DiscardToDeck() {
	for _, c := range f.Discard.GetCards() {
		f.Deck.Add(c)
	}
	f.Discard.Remove()
}

func (f *Field) ComputeSumProgress() int {
	return f.Downcard.GetSum()
}

func (f *Field) GetDeckLength() int {
	return f.Deck.GetLength()
}

func (f *Field) GetDiscardLength() int {
	return f.Discard.GetLength()
}

func (f *Field) GetDowncardLength() int {
	return f.Downcard.GetLength()
}

func (f *Field) SetCurrentPlayer(p player.Player) {
	f.CurrentPlayer = p
}

func (f *Field) SetLastPlayer(p player.Player) {
	f.LastPlayer = p
}

func (f *Field) SetNextPlayer(p player.Player) {
	f.NextPlayer = p
}

func (f *Field) SetOppPlayer(p player.Player) {
	f.OppPlayer = p
}

func (f *Field) ResetYellowCards() {
	f.CurrentPlayer.ResetWarn()
	f.NextPlayer.ResetWarn()
	f.LastPlayer.ResetWarn()
	f.OppPlayer.ResetWarn()
}

func (f *Field) ResetRedCards() {
	f.CurrentPlayer.ResetSuspend()
	f.NextPlayer.ResetSuspend()
	f.LastPlayer.ResetSuspend()
	f.OppPlayer.ResetSuspend()
}
