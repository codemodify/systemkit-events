package events

import (
	"reflect"
)

func (thisRef *Manager) addEventIfNotExists(eventName string) {
	// thisRef.eventsMutex.Lock()
	// defer thisRef.eventsMutex.Unlock()

	if _, ok := thisRef.events[eventName]; !ok {
		thisRef.events[eventName] = &event{
			subscribers:         []handlerContainer{},
			subscribersWithData: []handlerContainerWithData{},
		}
	}
}

func (thisRef *Manager) addSubscriberIfNotExists(eventName string, eventHandler Handler, shouldCallOnce bool) bool {
	// thisRef.eventsMutex.Lock()
	// defer thisRef.eventsMutex.Unlock()

	// 1. Check if delegate for the event already there, assumes map-key exists
	var alreadyThere = false

	for _, existingEventHandler := range thisRef.events[eventName].subscribers {
		alreadyThere = (reflect.ValueOf(eventHandler) == reflect.ValueOf(existingEventHandler.EventHandler))
		if alreadyThere {
			break
		}
	}

	// 2. Add the delegate
	if !alreadyThere {
		thisRef.events[eventName].subscribers = append(
			thisRef.events[eventName].subscribers,
			handlerContainer{
				EventHandler: eventHandler,
				CallOnce:     shouldCallOnce,
			},
		)
	}

	return alreadyThere
}

func (thisRef *Manager) addSubscriberWithDataIfNotExists(eventName string, eventHandlerWithData HandlerWithData, shouldCallOnce bool) bool {
	// thisRef.eventsMutex.Lock()
	// defer thisRef.eventsMutex.Unlock()

	// 1. Check if delegate for the event already there, assumes map-key exists
	var alreadyThere = false

	for _, existingEventHandler := range thisRef.events[eventName].subscribersWithData {
		alreadyThere = (reflect.ValueOf(eventHandlerWithData) == reflect.ValueOf(existingEventHandler.EventHandler))
		if alreadyThere {
			break
		}
	}

	// 2. Add the delegate
	if !alreadyThere {
		thisRef.events[eventName].subscribersWithData = append(
			thisRef.events[eventName].subscribersWithData,
			handlerContainerWithData{
				EventHandler: eventHandlerWithData,
				CallOnce:     shouldCallOnce,
			},
		)
	}

	return alreadyThere
}
