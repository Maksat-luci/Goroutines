// Если при запуске программы все горутины окажутся в состоянии ожидания
// runtime выкинет панику, с сообщении all goroutines are asleep - deadlock!
package main

import (
	"fmt"
	"time"
)
// Ball абстрактная структура мяча со значением удары
type Ball struct {
	hits int
}

func main() {
	// создаём канал для взаимодействия игроков
	table := make(chan *Ball)

	// Создаем пару игроков
	go player("ping", table)
	time.Sleep(2 * time.Second)

	go player("pong", table)

	table <- new(Ball) // Запуска мяча в игру
	time.Sleep(2 * time.Second)
	<-table // Конец игры, забираем мяч


}

func player(name string, table chan *Ball) {
	for {
		// Ждем когда мяч попал к игроку
		ball := <-table
		// увеличиваем счетчик ударов
		ball.hits++
		fmt.Println(name, ball.hits)
		// ждем немного 
		time.Sleep(100 * time.Millisecond)
		// Отправляем мяч обратно в канал 
		// Важно, программа заблокируется, пока другой игрок оттуда не прочитает
		table <- ball
	}
}
