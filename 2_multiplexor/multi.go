package main

import (
	"fmt"
)

func main() {
	//
	// c := doubleChannel(say("Ping"), say("Pong"))
	c := make(chan string)
	// go func() {
	go func() {
		for i := 1000; i != 0; i-- {
			if i == 999 {
				c <- "hahas"
			}
		}
	}()

	go func() {
		for say := range c{
			fmt.Println(say)
		}
	}()
	fmt.Scanln()
	fmt.Println("konec igry!")
}

func doubleChannel(input1, input2 <-chan string) <-chan string {
	// Создаём канал стринговый
	c := make(chan string)
	// запускаем горутину
	go func() {
		for {
			// записываем то что пришло из канала
			// такой синтаксис говорит о том что мы одновременно считываем с канала и одновремено записываем
			c <- <-input1
		}
	}()
	go func() {
		for {
			// записываем то что пришло из канала
			// аналогично
			c <- <-input2
		}
	}()
	return c

}

func say(msg string) <-chan string {
	// создаём канал
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			// записываем в канал стринг из поступающего стринга и I
			c <- fmt.Sprintf("%s %d", msg, i)
			// приостанавливаем работу функции на 1 секунду
		}
	}()

	return c
}
