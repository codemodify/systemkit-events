# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Events
[![GoDoc](https://godoc.org/github.com/codemodify/systemkit-logging?status.svg)](https://godoc.org/github.com/codemodify/systemkit-events)
[![0-License](https://img.shields.io/badge/license-0--license-brightgreen)](https://github.com/codemodify/TheFreeLicense)
[![Go Report Card](https://goreportcard.com/badge/github.com/codemodify/systemkit-logging)](https://goreportcard.com/report/github.com/codemodify/systemkit-logging)
[![Test Status](https://github.com/danawoodman/systemservice/workflows/Test/badge.svg)](https://github.com/danawoodman/systemservice/actions)
![code size](https://img.shields.io/github/languages/code-size/codemodify/SystemKit?style=flat-square)

#### Robust events for Go. It takes `0.1` seconds to send/receive 1 Million events.


# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Install
```go
go get github.com/codemodify/systemkit-events
```

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) API

&nbsp;										| &nbsp;
---											| ---
On(`event`, `handler`, `callOnce`)			| Subscribe
OnWithData(`event`, `handler`, `callOnce`)	| Subscribe with payload
Off(`event`, `handlerRef`)					| Unsubscribe
OffWithData(`event`, `handlerRef`)			| Unsubscribe with payload
Emit(`event`)								| Emit
EmitWithData(`event`, `data`)				| Emit with payload



# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Usage: Subscribe + Notify
```go
package main

import (
	events "github.com/codemodify/systemkit-events"
)

func main() {
	events.Events().On("PING", func() {
		// FIMXE: will be called
	})

	events.Events().Emit("PING")
}
```

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Usage: Subscribe (data) + Notify (data)
```go
package main

import (
	events "github.com/codemodify/systemkit-events"
)

func main() {
	events.Events().OnWithData("PING", func(data []byte) {
		// FIMXE: will be called
	})

	events.Events().EmitWithData("PING", []byte("PING-DATA"))
}
```

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Usage: Subscribe (yes-no-data) + Notify (data)
```go
package main

import (
	events "github.com/codemodify/systemkit-events"
)

func main() {
	events.Events().OnWithData("PING", func(data []byte) {
		// FIMXE: will be called
	})

	events.Events().On("PING", func() {
		// FIMXE: will be called
	})

	events.Events().EmitWithData("PING", []byte("PING-DATA"))
}
```

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Usage: Subscribe (yes-no-data) + Notify
```go
package main

import (
	events "github.com/codemodify/systemkit-events"
)

func main() {
	events.Events().OnWithData("PING", func(data []byte) {
		// FIMXE: will be called, data will be nil
	})

	events.Events().On("PING", func() {
		// FIMXE: will be called
	})

	events.Events().Emit("PING")
}
```

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Usage: Subscribe Once + Notify
```go
package main

import (
	events "github.com/codemodify/systemkit-events"
)

func main() {

	handler1CallOnce := true
	events.Events().On("PING", func() {
		// FIMXE: will be called ONCE
	}, handler1CallOnce)

	handler3CallOnce := false
	events.Events().On("PING", func() {
		// FIMXE: will be called UNTIL "events.Events().Off()" is called
	}, handler3CallOnce)

	events.Events().Emit("PING")
}
```
