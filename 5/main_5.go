package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(CheckBoth(a, b))

}

func CheckBoth(slice1 []int, slice2 []int) (bool, []int) {

	result := []int{}

	newSlice := append([]int{}, slice1...)
	newSlice = append(newSlice, slice2...)

	for _, value := range newSlice {
		if contains(slice1, value) && contains(slice2, value) && !contains(result, value) {
			result = append(result, value)
		}
	}
	return len(result) != 0, result

}

func contains(slice []int, target int) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}
