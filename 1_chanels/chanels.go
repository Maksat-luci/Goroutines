package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаём канал, он может быть и структурой и интерфейсом любым типом данных
	c := make(chan string)
	// запускаем горутину ping_pong и передаём туда наш канал типа string
	go pingPong(c)
	// здесь мы запускаем цикл и считываем с нашего канала в который записала наша горутина PING_PONG
	for i := 0; i < 5; i++ {
		fmt.Println(<-c, ", ", <-c)
	}
	// Создаём буферзированный канал, тоесть канал у которого есть лимит записи, если привысить указанное значение то произайдёт DeadLock  
	rangeChan := make(chan int, 10)
	for i := 1; i < 10; i++ {
		// записываем в канал значение i 
		rangeChan <- i
	}
	close(rangeChan)// закрываем канал и сюда больше нельзя писать иначе будет паника
	// отправляем в функцию sum() наш канал заполненный числами, функция считывает их как в массиве и суммирует
	fmt.Println("Sum", sum(rangeChan))
	g := generatorData(true)
	for i := 0; i < 5; i++ {
		fmt.Printf("Print %d\n", <-g)
	}

}
//ping_pong функция которая принимает канал данная стрелка говорит о том что внутри функции мы можем тольуо записывать в канал.
func pingPong(c chan<- string) {
	for {
		c <- fmt.Sprintf("Ping")
		c <- fmt.Sprintf("Pong")
	}
}

func sum(input <-chan int) (res int) {
	for r := range input {
		res += r
	}
	return
}

func generatorData(even bool) <-chan int {
	j := make(chan int)
	go func() {
		i := 0
		for {
			if i%2 == 0 {
				if even {
					j <- i
				}
			} else {
				if !even {
					j <- i
				}
			}
			time.Sleep(time.Second)
			i++
		}
	}()
	return j
}
