package main

import "fmt"

func main() {
	arraysAndSlices()
}

func arraysAndSlices() {
	//Arrays
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	fmt.Println("Array:", a)
	fmt.Println("First element:", a[0])
	fmt.Println("Length:", len(a))

	var b [5]string = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("Array b:", b)

	//Slices
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", s)

	s = append(s, 6)
	fmt.Println("After appending 6:", s)
}
