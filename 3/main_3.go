package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (sim *StringIntMap) Add(key string, value int) {
	sim.data[key] = value
}

func (sim *StringIntMap) Remove(key string) {
	delete(sim.data, key)
}

func (sim *StringIntMap) Copy() *StringIntMap {

	newSim := NewStringIntMap()

	for key, value := range sim.data {
		newSim.data[key] = value
	}

	return newSim
}

func (sim *StringIntMap) Exists(key string) bool {
	for k, _ := range sim.data {
		if key == k {
			return true
		}
	}
	return false
}

func (sim *StringIntMap) Get(key string) (int, bool) {

	if sim.Exists(key) {
		return sim.data[key], true
	}
	return sim.data[key], false
}

func main() {

	//ЗАДАНИЕ 3.0
	test := NewStringIntMap()

	//ЗАДАНИЕ 3.1
	test.Add("test1", 15)
	fmt.Println(test.Get("test1"))

	//ЗАДАНИЕ 3.2
	test.Add("test2", 44)
	test.Remove("test1")
	fmt.Println(test.data)

	//ЗАДАНИЕ 3.3
	newTest := test.Copy()
	newTest.Add("newTest", 44444)
	fmt.Println("test:", test.data)
	fmt.Println("newTest:", newTest.data)

	//ЗАДАНИЕ 3.4
	fmt.Println(newTest.Exists("test2"))

	//ЗАДАНИЕ 3.5
	fmt.Println(newTest.Get("test2"))
	fmt.Println(newTest.Get("test2555"))
}
