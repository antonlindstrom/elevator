package main

import (
	"sync"
)

type observer interface {
	Update()
}

type ElevatorObserver interface {
	Add(observer)
	Remove(observer)
	NotifyAll()
}

type elevatorObserver struct {
	observers map[observer]bool

	sync.RWMutex
}

func NewObserver() ElevatorObserver {
	return &elevatorObserver{
		observers: make(map[observer]bool),
	}
}

func (e *elevatorObserver) Add(ob observer) {
	e.Lock()
	defer e.Unlock()
	e.observers[ob] = true
}

func (e *elevatorObserver) Remove(ob observer) {
	e.Lock()
	defer e.Unlock()
	delete(e.observers, ob)
}

func (e *elevatorObserver) NotifyAll() {
	e.RLock()
	observers := e.observers
	e.RUnlock()

	for observer := range observers {
		observer.Update()
	}
}
