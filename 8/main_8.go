package main

import (
	"fmt"
	"time"
)

type CustomWaitGroup struct {
	sem     chan struct{}
	counter int
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		sem: make(chan struct{}, 1),
	}
}

func (wg *CustomWaitGroup) Add(delta int) {
	// Захватываем "мьютекс" для безопасного изменения counter
	wg.sem <- struct{}{}
	defer func() { <-wg.sem }()

	wg.counter += delta

	if wg.counter < 0 {
		panic("negative WaitGroup counter")
	}

	// Если счетчик стал 0, разблокируем Wait
	if wg.counter == 0 {
		// Пробуждаем все ждущие горутины
		select {
		case wg.sem <- struct{}{}: // Освобождаем если нужно
		default:
		}
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	for {
		wg.sem <- struct{}{}
		if wg.counter == 0 {
			<-wg.sem
			return
		}
		<-wg.sem

		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	wg := NewCustomWaitGroup()

	fmt.Println("Запускаем горутины")
	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Горутина %d начала работу\n", id)
			time.Sleep(time.Duration(id) * time.Second)
			fmt.Printf("Горутина %d завершила работу\n", id)
		}(i)
	}

	fmt.Println("Ожидаем завершения всех горутин...")
	wg.Wait()
	fmt.Println("Все горутины завершили работу!")

	fmt.Println("\nПовторное использование WaitGroup:")
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Дополнительная горутина 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Дополнительная горутина 2")
	}()

	wg.Wait()
	fmt.Println("Все дополнительные горутины завершены!")
}
