package int

import (
	"testing"
	"unsafe"
)

func TestInt(t *testing.T) {
	t.Logf("get int => %d size=%d bytes", Int, unsafe.Sizeof(Int))
}
