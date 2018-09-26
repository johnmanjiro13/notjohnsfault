package util

import "github.com/johnmanjiro13/notjohnsfault/game/field"

func TransitState(playField *field.Field) {
	tmpPlayer := playField.CurrentPlayer
	playField.SetCurrentPlayer(playField.NextPlayer)
	playField.SetNextPlayer(playField.OppPlayer)
	playField.SetOppPlayer(playField.LastPlayer)
	playField.SetLastPlayer(tmpPlayer)
}
