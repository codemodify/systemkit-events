package events

// On - add a subscriber for an event
func (thisRef *Manager) On(eventName string, eventHandler Handler, callOnce ...bool) {
	thisRef.eventsMutex.Lock()
	defer thisRef.eventsMutex.Unlock()

	shouldCallOnce := false
	if len(callOnce) > 0 {
		shouldCallOnce = callOnce[0]
	}

	thisRef.addEventIfNotExists(eventName)
	thisRef.addSubscriberIfNotExists(eventName, eventHandler, shouldCallOnce)
}

// OnWithData - add a subscriber with data for an event
func (thisRef *Manager) OnWithData(eventName string, eventHandlerWithData HandlerWithData, callOnce ...bool) {
	thisRef.eventsMutex.Lock()
	defer thisRef.eventsMutex.Unlock()

	shouldCallOnce := false
	if len(callOnce) > 0 {
		shouldCallOnce = callOnce[0]
	}

	thisRef.addEventIfNotExists(eventName)
	thisRef.addSubscriberWithDataIfNotExists(eventName, eventHandlerWithData, shouldCallOnce)
}
