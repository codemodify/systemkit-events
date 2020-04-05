package events

import "encoding/json"

// Emit - informs all subscribers about the event
func (thisRef *Manager) Emit(eventName string) {
	thisRef.eventsMutex.RLock()
	defer thisRef.eventsMutex.RUnlock()

	thisRef.addEventIfNotExists(eventName)

	thisRef.raiseHelper(eventName)
	thisRef.raiseWithDataHelper(eventName, nil)

	thisRef.removeAllCallOnce(eventName)
}

// EmitWithData - informs all subscribers about the event with data
func (thisRef *Manager) EmitWithData(eventName string, data []byte, autoMarshal ...bool) {
	thisRef.eventsMutex.RLock()
	defer thisRef.eventsMutex.RUnlock()

	thisRef.addEventIfNotExists(eventName)

	thisRef.raiseWithDataHelper(eventName, data)
	thisRef.raiseHelper(eventName)

	thisRef.removeAllCallOnce(eventName)
}

// EmitWithDataM - marshal data and informs all subscribers about the event with data
func (thisRef *Manager) EmitWithDataM(eventName string, structAsInterface interface{}) error {
	structAsBytes, err := json.Marshal(structAsInterface)
	if err != nil {
		return err
	}

	thisRef.EmitWithData(eventName, structAsBytes)

	return nil
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
