package main

import "testing"

func TestGenerator(t *testing.T) {
	count := 5
	ch := Generator(count)

	received := 0
	for range ch {
		received++
	}

	if received != count {
		t.Errorf("Есть: %d, надо: %d", received, count)
	}
}

func TestGeneratorZero(t *testing.T) {
	ch := Generator(0)

	received := 0
	for range ch {
		received++
	}

	if received != 0 {
		t.Errorf("Есть: %d, надо: 0", received)
	}
}

func TestGeneratorOne(t *testing.T) {
	ch := Generator(1)

	received := 0
	for num := range ch {
		if num < 0 || num >= 10000 {
			t.Errorf("Число %d вне диапазона", num)
		}
		received++
	}

	if received != 1 {
		t.Errorf("Есть: %d, надо: 1", received)
	}
}

func TestGeneratorMultiple(t *testing.T) {
	count := 10
	ch := Generator(count)

	received := 0
	for num := range ch {
		if num < 0 || num >= 10000 {
			t.Errorf("Число %d вне диапазона", num)
		}
		received++
	}

	if received != count {
		t.Errorf("Есть: %d, надо: %d", received, count)
	}
}
