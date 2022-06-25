package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Worker структура
type Worker struct {
	id int
	die chan bool 

	*sync.WaitGroup
}

//DoWork Метод который возвращает работу нашего воркера
func (w *Worker) DoWork() <-chan string {
	c := make(chan string)
	// Выполняет определенную работу и заполняет канал С
	go func() {
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond):
				c <- fmt.Sprintf("Worker #%d, do some work - %d", w.id, rand.Intn(100))
				// слушаем канал die ЕСЛИ  рабочий умер то уведомляем 
			case <-w.die:
				fmt.Println("Finish work!")
				return
				
			}
		}
	}()
	// Возвращает канал С с данными 
	return c
}

// Die Метод для закрытия воркера
func (w *Worker) Die() {
	fmt.Printf("Worker die %d\n", w.id)
	w.die <- true
	// минусуем счётчик WaitGroup 
	w.Done()
	// закрываем канал
	close(w.die)
}

// NewWorker конструктор длс создания воркера
func NewWorker(id int, wg *sync.WaitGroup) *Worker {
	if id <= 0 {
		panic("id must be > 0")
	}else if wg == nil {
		panic("WaitGroup can`t be empty")
	}
	// чекаем на валидность данные 
	// создаём канал для нашей структуры которая будет означать что рабочий умер
	die := make(chan bool, 1) // канал буфер или получим ДеадЛок
	// увеличиваем счётчик waitGroup на один 
	wg.Add(1)
	// возвращаем структуры для Worker с заполненными филдами 
	return &Worker{id, die, wg}

}

func main() {
	// Создаём структуру WaitGroup
	var wg sync.WaitGroup
	// создаём новых воркеров через абстрактную структуру Worker
	w1 := NewWorker(1, &wg)
	w2 := NewWorker(2, &wg)
	// заставляем их выполнять работу и запихиваем информацию о их деятельности либо смерти в переменные
	res1 := w1.DoWork()
	res2 := w2.DoWork()

	for i:=0; i < 6; i++{
		// запускаем цикл для печати их работы
		fmt.Println(<-res1)
		fmt.Println(<-res2)
		
		//OR

		// select {
		// case r1 := <- res1:
		// 	fmt.Println(r1)
		// case r2 := <- res2:
		// 	fmt.Println(r2)
		// }
	}
	// убиваем рабочих и с помощью кейса в DoWork узнаём об их смерти соответсвенно waitgroup понижается на 1
	// и wg.Wait перестаёт работать и горутина main завершается
	w1.Die()
	w2.Die()
	// тут заставляем ждать
	wg.Wait()
	// программа завершилась
	fmt.Println("End Program")
}