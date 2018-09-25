package event

import (
	"fmt"
)

type DoneNotify struct{}

type Done chan *DoneNotify

func GetDoneEvnentID() EventID {
	return EventID("doneEvent")
}

func (e Done) GetID() EventID {
	return GetDoneEvnentID()
}

func (e Done) Emit(in IEventNotify) error {
	n, ok := in.(*DoneNotify)
	if !ok {
		return fmt.Errorf("unexpected event notify specified. expected=Donenotify, actual=%s", n)
	}
	e <- n
	return nil
}
