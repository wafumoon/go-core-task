package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	newSlice := OnlyInFirstSlice(slice1, slice2)
	fmt.Println(newSlice)

}

func OnlyInFirstSlice(slice1 []string, slice2 []string) []string {

	bothSlice := append([]string{}, slice1...)
	bothSlice = append(bothSlice, slice2...)

	newSlice := []string{}

	for _, value := range bothSlice {
		if contains(slice1, value) && !contains(slice2, value) {
			newSlice = append(newSlice, value)
		}
	}

	return newSlice

}

func contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}
