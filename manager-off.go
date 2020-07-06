package events

import (
	"reflect"
)

// Off - remove a subscriber for an event
func (thisRef *Manager) Off(eventName string, eventHandler Handler) {
	thisRef.eventsMutex.Lock()
	defer thisRef.eventsMutex.Unlock()

	thisRef.addEventIfNotExists(eventName)

	thisRef.removeHandler(eventName, eventHandler)
}

// OffWithData - remove a subscriber with data for an event
func (thisRef *Manager) OffWithData(eventName string, eventHandlerWithData HandlerWithData) {
	thisRef.eventsMutex.Lock()
	defer thisRef.eventsMutex.Unlock()

	thisRef.addEventIfNotExists(eventName)

	thisRef.removeHandlerWithData(eventName, eventHandlerWithData)
}

func (thisRef *Manager) removeHandler(eventName string, eventHandler Handler) {
	var foundIndex = -1
	for index, existingEventHandler := range thisRef.events[eventName].subscribers {
		if reflect.ValueOf(eventHandler) == reflect.ValueOf(existingEventHandler.EventHandler) {
			foundIndex = index
			break
		}
	}

	if foundIndex != -1 {
		thisRef.events[eventName].subscribers = append(
			thisRef.events[eventName].subscribers[:foundIndex],
			thisRef.events[eventName].subscribers[foundIndex+1:]...,
		)
	}
}

func (thisRef *Manager) removeHandlerWithData(eventName string, eventHandlerWithData HandlerWithData) {
	var foundIndex = -1
	for index, existingEventHandler := range thisRef.events[eventName].subscribersWithData {
		if reflect.ValueOf(eventHandlerWithData) == reflect.ValueOf(existingEventHandler.EventHandler) {
			foundIndex = index
			break
		}
	}

	if foundIndex != -1 {
		thisRef.events[eventName].subscribersWithData = append(
			thisRef.events[eventName].subscribersWithData[:foundIndex],
			thisRef.events[eventName].subscribersWithData[foundIndex+1:]...,
		)
	}
}
