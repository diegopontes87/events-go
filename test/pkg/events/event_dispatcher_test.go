package events

import (
	"time"

	"github.com/diegopontes87/events-go/pkg/events/interfaces"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

type TestEventHandler struct {
	ID int
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

func (h *TestEventHandler) Handle(event interfaces.EventInterface) {

}
