package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) ([]int){
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	for i, v := range slice {
		slice[i] = f(v)
	}
}

func mapArray(f func(a int) int, array [3]int) {
	for i, v := range array {
		array[i] = f(v)
	}
}

func main() {
	intsSlice := []int{1, 6, 2, 3, 4, 5}
	initSlice2 := []int{0,1,2,3}


	fmt.Println(initSlice2)
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	newSlice := intsSlice[1:3]
	mapSlice(square, newSlice)
	fmt.Println(newSlice)
	fmt.Println(intsSlice)

	intsSlice = double(intsSlice)
	fmt.Println(intsSlice)
}
