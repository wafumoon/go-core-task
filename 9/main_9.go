package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []int{
		15,
		25,
		44,
		43,
		15,
		86,
		11,
		111,
		1,
		19,
		50,
	}

	channelIn := make(chan int)
	channelOut := make(chan float64)

	go func() {
		for _, number := range numbers {
			channelIn <- number
		}
		close(channelIn)
	}()

	go func() {
		for number := range channelIn {
			channelOut <- math.Pow(float64(number), 3)
		}
		close(channelOut)
	}()

	// go func() {
	// 	//number := math.Pow(float64(<-channelIn), 3)
	// 	channelOut <- math.Pow(float64(<-channelIn), 3)

	// }()

	for result := range channelOut {
		fmt.Println(result)
	}
}
