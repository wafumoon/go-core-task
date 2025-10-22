package main

import "testing"

func TestStringIntMap(t *testing.T) {
	m := NewStringIntMap()

	m.Add("test", 42)
	if value, _ := m.Get("test"); value != 42 {
		t.Errorf("Есть: %d, надо: 42", value)
	}

	if !m.Exists("test") {
		t.Error("Должен существовать")
	}

	m.Remove("test")
	if m.Exists("test") {
		t.Error("Должен быть удален")
	}
}

func TestCopy(t *testing.T) {
	original := NewStringIntMap()
	original.Add("a", 1)
	original.Add("b", 2)

	copy := original.Copy()

	if !copy.Exists("a") || !copy.Exists("b") {
		t.Error("Копия должна содержать те же элементы")
	}

	original.Add("c", 3)
	if copy.Exists("c") {
		t.Error("Копия не должна меняться при изменении оригинала")
	}
}

func TestGetNonExistent(t *testing.T) {
	m := NewStringIntMap()

	value, exists := m.Get("nonexistent")
	if exists {
		t.Error("Не должен существовать")
	}
	if value != 0 {
		t.Errorf("Есть: %d, надо: 0", value)
	}
}

func TestMultipleOperations(t *testing.T) {
	m := NewStringIntMap()

	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	if !m.Exists("one") || !m.Exists("two") {
		t.Error("Элементы должны существовать")
	}

	m.Remove("two")
	if m.Exists("two") {
		t.Error("Элемент должен быть удален")
	}

	if value, _ := m.Get("three"); value != 3 {
		t.Errorf("Есть: %d, надо: 3", value)
	}
}
