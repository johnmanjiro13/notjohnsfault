package event

import (
	"fmt"

	"github.com/johnmanjiro13/notjohnsfault/util"
)

type Event struct {
	EventName string
}

var events = []string{"reset", "standby", "draw", "report", "milestone", "judge"}

func NewEvent() *Event {
	return &Event{
		EventName: "reset",
	}
}

func (e Event) GetCurrentEvent() string {
	return e.EventName
}

func (e *Event) SetEvent(nextEvent string) error {
	if util.ArrayContainsString(events, nextEvent) {
		e.EventName = nextEvent
		return nil
	}
	return fmt.Errorf("%s is not valid event", nextEvent)
}
