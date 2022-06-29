package goroutine

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestIncrement(t *testing.T) {
	var num int64 = 0
	go func() {
		for {
			num++
		}
	}()
	time.Sleep(time.Second)
	fmt.Printf("num=%d\n", num)
}

func TestIncrement_With_UnreachableCode(t *testing.T) {
	var num int64 = 0
	go func() {
		for {
			num++
			if num < 0 {
				break
			}
		}
	}()
	time.Sleep(time.Second)
	fmt.Printf("num=%d\n", num)
}

func TestIncrement_With_Atomic(t *testing.T) {
	var num int64 = 0
	go func() {
		for {
			atomic.AddInt64(&num, 1)
		}
	}()
	time.Sleep(time.Second)
	fmt.Printf("num=%d\n", atomic.LoadInt64(&num))
}
