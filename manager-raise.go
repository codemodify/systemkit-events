package events

// Raise - informs all subscribers about the event
func (thisRef *Manager) Raise(eventName string) {
	thisRef.eventsMutex.RLock()
	defer thisRef.eventsMutex.RUnlock()

	thisRef.addEventIfNotExists(eventName)

	thisRef.raiseHelper(eventName)
	thisRef.raiseWithDataHelper(eventName, nil)

	thisRef.removeAllCallOnce(eventName)
}

// RaiseWithData - informs all subscribers about the event with data
func (thisRef *Manager) RaiseWithData(eventName string, data []byte) {
	thisRef.eventsMutex.RLock()
	defer thisRef.eventsMutex.RUnlock()

	thisRef.addEventIfNotExists(eventName)

	thisRef.raiseWithDataHelper(eventName, data)
	thisRef.raiseHelper(eventName)

	thisRef.removeAllCallOnce(eventName)
}

func (thisRef *Manager) raiseHelper(eventName string) {
	for _, eventHandlerC := range thisRef.events[eventName].subscribers {
		go eventHandlerC.EventHandler()
	}
}

func (thisRef *Manager) raiseWithDataHelper(eventName string, data []byte) {
	for _, eventHandlerC := range thisRef.events[eventName].subscribersWithData {
		go eventHandlerC.EventHandler(data)
	}
}

func (thisRef *Manager) removeAllCallOnce(eventName string) {
	for _, eventHandlerC := range thisRef.events[eventName].subscribers {
		if eventHandlerC.CallOnce {
			thisRef.removeHandler(eventName, eventHandlerC.EventHandler)
		}
	}

	for _, eventHandlerC := range thisRef.events[eventName].subscribersWithData {
		if eventHandlerC.CallOnce {
			thisRef.removeHandlerWithData(eventName, eventHandlerC.EventHandler)
		}
	}
}
