package tokenize

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var benchText = `Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.`

//go test -bench . -run NONE // NONE for not running any test
//go test -bench . -run NONE -benchmem // to get memory allocation
//go test -bench . -run NONE -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof // to get cpu and memory profile
//go tool pprof -http=:8080 cpu.prof // to get cpu profile
//go tool pprof -http=:8080 mem.prof // to get memory profile
//go test -bench . -run NONE -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof -blockprofile=block.prof // to get block profile
//go tool pprof -http=:8080 block.prof // to get block profile
//go test -bench . -run NONE -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof -blockprofile=block.prof -mutexprofile=mutex.prof // to get mutex profile
//go tool pprof -http=:8080 mutex.prof // to get mutex profile
//go test -bench . -run NONE -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof -blockprofile=block.prof -mutexprofile=mutex.prof -trace=trace.out // to get trace profile
//go tool trace trace.out // to get trace profile
//go test -bench . -run NONE -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof -blockprofile=block.prof -mutexprofile=mutex.prof -trace=trace.out -coverprofile=cover.out // to get coverage profile

func BenchmarkTokenize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tokens := Tokenize(benchText)
		require.Equal(b, 18, len(tokens))
	}
}
