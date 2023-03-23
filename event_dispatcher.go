package events

import (
	"errors"
)

var ErrorHandlerAlredyExists = errors.New("Handler Alredy Existes")

type EventDispatcher struct {
	Handlers map[string][]EventHanlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		Handlers: make(map[string][]EventHanlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHanlerInterface) error {
	if _, ok := ed.Handlers[eventName]; ok {
		for _, h := range ed.Handlers[eventName] {
			if h == handler {
				return ErrorHandlerAlredyExists
			}

		}
	}
	ed.Handlers[eventName] = append(ed.Handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Clear() error {
	ed.Handlers = make(map[string][]EventHanlerInterface)
	return nil
}

func (ed *EventDispatcher) Has(eventName string, handler EventHanlerInterface) bool {
	if _, ok := ed.Handlers[eventName]; ok {
		for _, h := range ed.Handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Dispatcher(event EventInterface) error {
	if handlers, ok := ed.Handlers[event.GetName()]; ok {
		for _, h := range handlers {
			h.Handle(event)
		}

	}
	return nil
}
