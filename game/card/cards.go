package card

type Cards []ICard

type Card struct {
	Number int
}

func (c Card) GetNumber() int {
	return c.Number
}
