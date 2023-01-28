package slice

import (
	"reflect"
	"testing"
)

func TestSlice_Builtin(t *testing.T) {
	var list []int
	for i := 1; i <= 2000; i++ {
		list = append(list, i)
		t.Logf("len=%d cap=%d", len(list), cap(list))
	}
}

func TestSlice_Reflect(t *testing.T) {
	// the growth rule of the reflect slice is different from the builtin slice
	list := reflect.ValueOf([]int{})
	for i := 1; i <= 2000; i++ {
		list = reflect.AppendSlice(list, reflect.ValueOf([]int{i}))
		t.Logf("len=%d cap=%d", list.Len(), list.Cap())
	}
}
