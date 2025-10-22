package main

import "testing"

func TestCheckBoth(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, result := CheckBoth(a, b)

	if !hasIntersection {
		t.Error("Должен быть true")
	}

	if len(result) != 2 {
		t.Errorf("Есть: %d, надо: 2", len(result))
	}

	if !contains(result, 3) || !contains(result, 64) {
		t.Error("Должны содержать 3 и 64")
	}
}

func TestCheckBothNoIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	hasIntersection, result := CheckBoth(a, b)

	if hasIntersection {
		t.Error("Должен быть false")
	}

	if len(result) != 0 {
		t.Errorf("Есть: %d, надо: 0", len(result))
	}
}

func TestCheckBothEmptyFirst(t *testing.T) {
	a := []int{}
	b := []int{1, 2, 3}

	hasIntersection, result := CheckBoth(a, b)

	if hasIntersection {
		t.Error("Должен быть false")
	}

	if len(result) != 0 {
		t.Errorf("Есть: %d, надо: 0", len(result))
	}
}

func TestCheckBothEmptySecond(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{}

	hasIntersection, result := CheckBoth(a, b)

	if hasIntersection {
		t.Error("Должен быть false")
	}

	if len(result) != 0 {
		t.Errorf("Есть: %d, надо: 0", len(result))
	}
}

func TestCheckBothBothEmpty(t *testing.T) {
	a := []int{}
	b := []int{}

	hasIntersection, result := CheckBoth(a, b)

	if hasIntersection {
		t.Error("Должен быть false")
	}

	if len(result) != 0 {
		t.Errorf("Есть: %d, надо: 0", len(result))
	}
}

func TestCheckBothDuplicates(t *testing.T) {
	a := []int{1, 1, 2, 2, 3}
	b := []int{1, 2, 2, 4}

	hasIntersection, result := CheckBoth(a, b)

	if !hasIntersection {
		t.Error("Должен быть true")
	}

	if len(result) != 2 {
		t.Errorf("Есть: %d, надо: 2", len(result))
	}

	if !contains(result, 1) || !contains(result, 2) {
		t.Error("Должны содержать 1 и 2")
	}
}

func TestContains(t *testing.T) {
	slice := []int{10, 20, 30}

	if !contains(slice, 20) {
		t.Error("Должен найти 20")
	}

	if contains(slice, 40) {
		t.Error("Не должен найти 40")
	}
}
