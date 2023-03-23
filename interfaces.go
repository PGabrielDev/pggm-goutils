package events

import (
	"time"
)

type EventInterface interface {
	GetName() string
	GetDate() time.Time
	GetPayload() interface{}
}

type EventHanlerInterface interface {
	Handle(event EventInterface)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHanlerInterface) error
	Dispatcher(event EventInterface) error
	Remove(eventName string, handler EventHanlerInterface) error
	Has(eventName string, handler EventHanlerInterface) bool
	Clear() error
}
