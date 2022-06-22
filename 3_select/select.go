package main

import (
	"fmt"
	"time"
)

func main() {
	//создаём каналы
	c1 := make(chan string)
	c2 := make(chan string)

	//запускаем горутину которая будет писать сообщение в канал с1
	go func() {
		for {
			c1 <- "message №1"
			time.Sleep(time.Second * 2)
		}
	}()
	//запускаем горутину которая будет писать сообщение в канал c2
	go func() {
		for {
			c2 <- "message #2"
			time.Sleep(time.Second * 2)
		}
	}()

	// В горутине читаем каналы через оператор select
	// Приоритетность того кто выше по case
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("msg1", msg1)

			case msg2 := <-c2:
				fmt.Println("msg2", msg2)
			// case <-time.After(time.Second): // каждую секунду функция After присылает true
			// 	fmt.Println("timeout")
			}

		}
	}()
	fmt.Scanln()
}
