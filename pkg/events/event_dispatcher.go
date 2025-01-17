package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("hanlder already registered")

type EventDispacher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispacher {
	return &EventDispacher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}
func (ed *EventDispacher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}
