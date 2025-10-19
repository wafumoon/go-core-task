package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := Generator(25)

	for number := range numbers {
		fmt.Println(number)
	}
}

func Generator(count int) <-chan int {
	channel := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			channel <- rand.Intn(10000)
		}
		close(channel)
	}()
	return channel
}
