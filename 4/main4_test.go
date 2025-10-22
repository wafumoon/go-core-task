package main

import "testing"

func TestOnlyInFirstSlice(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := OnlyInFirstSlice(slice1, slice2)
	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	if len(result) != len(expected) {
		t.Errorf("Есть: %d, надо: %d", len(result), len(expected))
	}

	for i, val := range expected {
		if result[i] != val {
			t.Errorf("На позиции %d: есть %s, надо %s", i, result[i], val)
		}
	}
}

func TestOnlyInFirstSliceEmptySecond(t *testing.T) {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{}

	result := OnlyInFirstSlice(slice1, slice2)

	if len(result) != 3 {
		t.Errorf("Есть: %d, надо: 3", len(result))
	}
}

func TestOnlyInFirstSliceEmptyFirst(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{"a", "b"}

	result := OnlyInFirstSlice(slice1, slice2)

	if len(result) != 0 {
		t.Errorf("Есть: %d, надо: 0", len(result))
	}
}

func TestOnlyInFirstSliceNoCommon(t *testing.T) {
	slice1 := []string{"x", "y", "z"}
	slice2 := []string{"a", "b", "c"}

	result := OnlyInFirstSlice(slice1, slice2)

	if len(result) != 3 {
		t.Errorf("Есть: %d, надо: 3", len(result))
	}
}

func TestOnlyInFirstSliceAllCommon(t *testing.T) {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"a", "b", "c"}

	result := OnlyInFirstSlice(slice1, slice2)

	if len(result) != 0 {
		t.Errorf("Есть: %d, надо: 0", len(result))
	}
}

func TestContains(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}

	if !contains(slice, "banana") {
		t.Error("Должен найти banana")
	}

	if contains(slice, "grape") {
		t.Error("Не должен найти grape")
	}
}
