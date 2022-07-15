package slice

import (
	"testing"
)

func TestIntSlice(t *testing.T) {
	t.Logf("int slice => %v len=%d cap=%d", IntSlice, len(IntSlice), cap(IntSlice))
}
