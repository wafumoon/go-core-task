package main

import (
	"testing"
)

func TestDefineType(t *testing.T) {
	testCases := []interface{}{
		42,
		3.14,
		"hello",
		true,
		1 + 2i,
		[]int{1, 2, 3},
	}

	for _, input := range testCases {
		defineType(input)
	}
}

func TestTransformToString(t *testing.T) {
	input := []interface{}{42, "test", true}
	result := transformToString(input)
	expected := "42testtrue"

	if result != expected {
		t.Errorf("Есть: %q, надо: %q", result, expected)
	}
}

func TestStringToRunes(t *testing.T) {
	result := stringToRunes("hello")
	expected := []rune{'h', 'e', 'l', 'l', 'o'}

	if len(result) != len(expected) {
		t.Errorf("Есть: %d, надо: %d", len(result), len(expected))
	}

	if len(result) > 0 && result[0] != 'h' {
		t.Errorf("Есть: %c, надо: 'h'", result[0])
	}
}

func TestHashRunesWithSalt(t *testing.T) {
	runes := []rune("test")
	result := hashRunesWithSalt(runes, "salt")

	if len(result) != 64 {
		t.Errorf("Есть: %d, надо: 64", len(result))
	}
}

func TestFull(t *testing.T) {
	variables := []interface{}{42, "hello", true}

	str := transformToString(variables)
	runes := stringToRunes(str)
	hash := hashRunesWithSalt(runes, "go-2024")

	if str == "" {
		t.Error("пустая строка")
	}
	if len(runes) == 0 {
		t.Error("пустой слайс рун")
	}
	if len(hash) != 64 {
		t.Error("ошибка хэша")
	}
}
