package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var number int64
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// Увеличиваем значение на 1-м, number++
				atomic.AddInt64(&number, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadInt64(&number)
	fmt.Println("final result:", opsFinal)
}
