package main

import "fmt"

func main() {

	//ЗАДАНИЕ 2.1
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//ЗАДАНИЕ 2.2
	fmt.Println(sliceExample(originalSlice))

	//ЗАДАНИЕ 2.3
	fmt.Println(addElements(originalSlice, 25))

	//ЗАДАНИЕ 2.4
	copiedSlice := copySlice(originalSlice)
	originalSlice = append(originalSlice, 11)
	fmt.Println("original:", originalSlice)
	fmt.Println("copy:    ", copiedSlice)

	//ЗАДАНИЕ 2.5
	removedSlice := removeElement(originalSlice, 55)
	fmt.Println("original:", originalSlice)
	fmt.Println("removed: ", removedSlice)
}

func sliceExample(slice []int) []int {

	result := []int{}

	for _, value := range slice {
		if value%2 == 0 {
			result = append(result, value)
		}
	}
	return result
}

func addElements(slice []int, value int) []int {
	return append(slice, value)
}

func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func removeElement(slice []int, value int) []int {

	if value >= len(slice) {
		value = len(slice) - 1
	}

	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:value])
	copy(newSlice[value:], slice[value+1:])
	return newSlice

}
