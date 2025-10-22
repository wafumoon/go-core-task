package main

import (
	"sync"
	"testing"
	"time"
)

func TestCustomWaitGroup_AddDoneWait(t *testing.T) {
	wg := NewCustomWaitGroup()
	var counter int
	var mu sync.Mutex

	wg.Add(1)

	go func() {
		defer wg.Done()
		mu.Lock()
		counter++
		mu.Unlock()
	}()

	wg.Wait()

	mu.Lock()
	if counter != 1 {
		t.Errorf("Expected counter = 1, got %d", counter)
	}
	mu.Unlock()
}

func TestCustomWaitGroup_MultipleGoroutines(t *testing.T) {
	wg := NewCustomWaitGroup()
	var counter int
	var mu sync.Mutex
	const numGoroutines = 5

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	mu.Lock()
	if counter != numGoroutines {
		t.Errorf("Expected counter = %d, got %d", numGoroutines, counter)
	}
	mu.Unlock()
}

func TestCustomWaitGroup_Reuse(t *testing.T) {
	wg := NewCustomWaitGroup()
	var results []int
	var mu sync.Mutex

	wg.Add(2)

	go func() {
		defer wg.Done()
		mu.Lock()
		results = append(results, 1)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		results = append(results, 2)
		mu.Unlock()
	}()

	wg.Wait()

	if len(results) != 2 {
		t.Errorf("First use: expected 2 results, got %d", len(results))
	}

	time.Sleep(10 * time.Millisecond)

	results = nil
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(val int) {
			defer wg.Done()
			mu.Lock()
			results = append(results, val)
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	if len(results) != 3 {
		t.Errorf("Reuse: expected 3 results, got %d", len(results))
	}
}

func TestCustomWaitGroup_ConcurrentAccess(t *testing.T) {
	wg := NewCustomWaitGroup()
	var counter int
	var wgMain sync.WaitGroup
	const numWriters = 10

	wg.Add(numWriters)
	wgMain.Add(numWriters)

	for i := 0; i < numWriters; i++ {
		go func() {
			defer wgMain.Done()
			defer wg.Done()
			time.Sleep(5 * time.Millisecond)
			counter++
		}()
	}

	wg.Wait()

	if counter != numWriters {
		t.Errorf("Expected counter = %d, got %d", numWriters, counter)
	}

	wgMain.Wait()
}

func TestCustomWaitGroup_NegativeCounter(t *testing.T) {
	wg := NewCustomWaitGroup()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic on negative counter")
		}
	}()

	wg.Add(-1)
}

func TestCustomWaitGroup_ZeroAdd(t *testing.T) {
	wg := NewCustomWaitGroup()

	wg.Add(0)
	wg.Wait()
}

func TestCustomWaitGroup_RapidSequence(t *testing.T) {
	wg := NewCustomWaitGroup()
	const iterations = 100
	ch := make(chan int, iterations)

	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		go func(val int) {
			defer wg.Done()
			ch <- val
		}(i)
	}

	wg.Wait()
	close(ch)

	count := 0
	for range ch {
		count++
	}

	if count != iterations {
		t.Errorf("Expected %d values, got %d", iterations, count)
	}
}

func TestCustomWaitGroup_WaitWithoutAdd(t *testing.T) {
	wg := NewCustomWaitGroup()

	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait should return immediately when no goroutines are added")
	}
}

func TestCustomWaitGroup_AddAfterWait(t *testing.T) {
	wg := NewCustomWaitGroup()

	wg.Add(1)
	go func() {
		defer wg.Done()
	}()

	wg.Wait()

	wg.Add(1)
	done := make(chan bool)

	go func() {
		defer wg.Done()
		done <- true
	}()

	wg.Wait()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Error("Second Wait should complete after second Add/Done")
	}
}
