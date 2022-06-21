package main

import "fmt"
// создаём свой типа данных стринг, плюс в том что мы можем его модернизировать добавляя к нему методы
type myString string

// Создаём метод для нашего типа myString 
func (*myString) process(i int) {
	fmt.Println("process struct ", i)
}
// Обычная функция 
func procces(i int) {
	fmt.Println("procces: ", i)
}

func main() {
	fmt.Println("Start goroutine")
	// можно запустить функцию 
	go procces(1)
	// можно запустить анонимную фукнцию
	go func() {
		fmt.Println("Anon func")
	}()
	
	// Можем запустить много горутин в цикле
	for i:=0; i < 10; i++{
		go procces(i)
	}
	// Создаём обьект Mystring
	myStr := new(myString)
	go myStr.process(1) // Запуск горутины исполняемым методом типа myString

	// Нужно дождаться завершения выполнения
	//Тут происходят часто первые ошибки запуска
	// если главная функция завершится до того как горутины успеют что-то сделать то все работающие горутины умрут, 
	//  поэтому мы должны дождаться выполнения горутин,
	// и замедлим функцию мейн с помошью fmt.Scanln() метод типа fmt  который ждёт ввода в консоль типа input() по голенговски
	fmt.Scanln()
}