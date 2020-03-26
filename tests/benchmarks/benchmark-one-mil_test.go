package tests

import (
	"testing"

	events "github.com/codemodify/systemkit-events"
)

func Benchmark_benchmark_one_mil(b *testing.B) {

	const pingEvent = "BENCH"
	const pingData = "PING-DATA"

	events.Events().OnWithData(pingEvent, func(data []byte) {
		// DO SOMETHING WITH DATA
	})

	b.ResetTimer()
	b.StartTimer()
	count := 0
	countMax := 1000000
	for ; count < countMax; count++ {
		events.Events().EmitWithData(pingEvent, []byte(pingData))
	}
	b.StopTimer()

	b.Logf("SEND/RECV: %d", count)
}
