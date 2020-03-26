package events

// Handler - event handler prototype
type Handler func()

// HandlerWithData - event handler with data prototype
type HandlerWithData func([]byte)

type handlerContainer struct {
	EventHandler Handler
	CallOnce     bool
}

type handlerContainerWithData struct {
	EventHandler HandlerWithData
	CallOnce     bool
}

type event struct {
	subscribers         []handlerContainer
	subscribersWithData []handlerContainerWithData
}
