// Analysis with the following command: go tool compile -m memory_escape_benchmark_test.go

package memory

import "testing"

type Hello interface {
	SayHi(format string, args ...interface{})
}

type HelloImpl struct {
}

func (h *HelloImpl) SayHi(format string, args ...interface{}) {

}

type HelloImpl2 struct {
}

func (h *HelloImpl2) SayHi(format string, args ...interface{}) {

}

var h1 *HelloImpl = &HelloImpl{}
var h2 Hello = &HelloImpl{}

// BenchmarkHello_Global_With_Interface-16     1000000000     0.2187 ns/op     0 B/op     0 allocs/op
func BenchmarkHello_Global_With_Interface(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h1.SayHi("hi", "go")
	}
}

// BenchmarkHello_Global_With_Struct-16     57404731     20.64 ns/op     16 B/op     1 allocs/op
func BenchmarkHello_Global_With_Struct(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h2.SayHi("hi", "go")
	}
}

// BenchmarkHello_Local_With_Struct-16     1000000000     0.2181 ns/op     0 B/op     0 allocs/op
func BenchmarkHello_Local_With_Struct(b *testing.B) {
	var h *HelloImpl = &HelloImpl{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.SayHi("hi", "go")
	}
}

// BenchmarkHello_Local_With_Interface_SingleImpl-16     1000000000     1.110 ns/op     0 B/op     0 allocs/op
func BenchmarkHello_Local_With_Interface_SingleImpl(b *testing.B) {
	var h Hello = &HelloImpl{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.SayHi("hi", "go")
	}
}

// BenchmarkHello_Local_With_Interface_MultiImpl-16     56007765     20.68 ns/op     16 B/op     1 allocs/op
func BenchmarkHello_Local_With_Interface_MultiImpl(b *testing.B) {
	var h Hello = &HelloImpl{}
	h = &HelloImpl2{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.SayHi("hi", "go")
	}
}
