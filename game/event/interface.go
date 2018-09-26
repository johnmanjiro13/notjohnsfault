package event

type IEvent interface {
	GetCurrentEvent() string
	SetEvent(nextEvent string)
}
