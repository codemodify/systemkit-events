package events

import (
	"sync"
)

// Manager - event manager
type Manager struct {
	events      map[string]*event
	eventsMutex *sync.RWMutex
}

var managerInstance *Manager
var managerOnce sync.Once

// Events - singleton to get the events manager instance
func Events() *Manager {
	managerOnce.Do(func() {
		managerInstance = &Manager{
			events:      make(map[string]*event),
			eventsMutex: &sync.RWMutex{},
		}
	})

	return managerInstance
}
