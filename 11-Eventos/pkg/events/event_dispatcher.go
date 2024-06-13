package events

import (
	"errors"
	"sync"
)

var ErrHandleAlreadyRegister = errors.New("event already registred")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if handler == h {
				return ErrHandleAlreadyRegister
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}
func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if handler == h {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			handler.Handle(event, wg)
		}
		wg.Wait()
	}

	return nil
}

func (ed *EventDispatcher) Remove(eventName string, handle EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == h {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}

		}

	}
	return nil
}
