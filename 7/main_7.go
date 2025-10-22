package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	newMergedInt := Merge(ch1, ch2, ch3)

	ch4 := make(chan string)
	ch5 := make(chan string)
	ch6 := make(chan string)

	newMergedString := Merge(ch4, ch5, ch6)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		ch2 <- 15
		ch2 <- 555
		close(ch2)
	}()

	go func() {
		ch4 <- "1515151"
		ch4 <- "aaaa"
		close(ch4)
	}()

	go func() {
		ch3 <- 14
		ch3 <- 14
		close(ch3)
	}()

	go func() {
		ch6 <- "14114141414"
		close(ch6)
	}()

	go func() {
		ch5 <- "ууууу"
		close(ch5)
	}()

	for value := range newMergedInt {
		fmt.Println(value)
	}

	for value := range newMergedString {
		fmt.Println(value)
	}
}

func Merge[T any](channels ...<-chan T) <-chan T {
	out := make(chan T)

	var wg sync.WaitGroup

	output := func(channel <-chan T) {
		defer wg.Done()
		for value := range channel {
			out <- value
		}
	}

	wg.Add(len(channels))
	for _, channel := range channels {
		go output(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
