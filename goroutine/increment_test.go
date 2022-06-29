// Output the assembly code with the following command: go tool compile -S increment_test.go

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
	// Output: num=0
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
	// Output: num=464742893
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
	// Output: num=172373819
	fmt.Printf("num=%d\n", atomic.LoadInt64(&num))
}
