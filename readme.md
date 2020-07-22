# ![](https://fonts.gstatic.com/s/i/materialiconsoutlined/flare/v4/24px.svg) Go Events
[![](https://img.shields.io/github/v/release/codemodify/systemkit-events?style=flat-square)](https://github.com/codemodify/systemkit-events/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-events?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-events?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-events/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-events?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-events?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-events)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-events)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-events?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-events?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-events?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-events?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-events?style=flat-square)

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
