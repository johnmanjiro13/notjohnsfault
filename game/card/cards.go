package card

type Cards []ICard

type Card struct {
	Number int
}

func (c Card) GetNumber() int {
	return c.Number
}

func GenerateCards() Cards {
	cards := Cards{}
	for i := 0; i < 32; i++ {
		switch {
		case (i >= 0 && i < 2):
			newCard := Card{Number: 0}
			cards = append(cards, newCard)
		case (i >= 2 && i < 6):
			newCard := Card{Number: 1}
			cards = append(cards, newCard)
		case (i >= 6 && i < 11):
			newCard := Card{Number: 2}
			cards = append(cards, newCard)
		case (i >= 11 && i < 16):
			newCard := Card{Number: 3}
			cards = append(cards, newCard)
		case (i >= 16 && i < 21):
			newCard := Card{Number: 4}
			cards = append(cards, newCard)
		case (i >= 21 && i < 26):
			newCard := Card{Number: 5}
			cards = append(cards, newCard)
		case (i >= 26 && i < 31):
			newCard := Card{Number: 6}
			cards = append(cards, newCard)
		}
	}
	return cards
}
