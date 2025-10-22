package main

import (
	"testing"
)

func TestMergeInt(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	merged := Merge(ch1, ch2)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	count := 0
	for range merged {
		count++
	}

	if count != 4 {
		t.Errorf("Есть: %d, надо: 4", count)
	}
}

func TestMergeString(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	merged := Merge(ch1, ch2)

	go func() {
		ch1 <- "a"
		ch1 <- "b"
		close(ch1)
	}()

	go func() {
		ch2 <- "c"
		ch2 <- "d"
		close(ch2)
	}()

	count := 0
	for range merged {
		count++
	}

	if count != 4 {
		t.Errorf("Есть: %d, надо: 4", count)
	}
}

func TestMergeSingleChannel(t *testing.T) {
	ch := make(chan int)

	merged := Merge(ch)

	go func() {
		ch <- 42
		close(ch)
	}()

	value := <-merged
	if value != 42 {
		t.Errorf("Есть: %d, надо: 42", value)
	}
}

func TestMergeEmpty(t *testing.T) {
	merged := Merge[int]()

	count := 0
	for range merged {
		count++
	}

	if count != 0 {
		t.Errorf("Есть: %d, надо: 0", count)
	}
}

func TestMergeWithEmptyChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	merged := Merge(ch1, ch2)

	go func() {
		close(ch1)
	}()

	go func() {
		close(ch2)
	}()

	count := 0
	for range merged {
		count++
	}

	if count != 0 {
		t.Errorf("Есть: %d, надо: 0", count)
	}
}
