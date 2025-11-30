package main

import "fmt"

func main() {
	arraysAndSlices()
	slices()
	average()
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

func average() {
	trasactions := []float64{100.5, 200.75, 50.25, 300.0}

	total := 0.0
	for _, amount := range trasactions {
		total += amount
	}

	average := total / float64(len(trasactions))
	fmt.Println("Average transaction amount:", average)
}

func slices() {
	slice := make([]int, 5, 10)

	fmt.Println("Slice length:", len(slice))
	fmt.Println("Slice capacity:", cap(slice))

	slice = append(slice, 1, 2, 3)
	fmt.Println("Slice after appending:", slice)
	fmt.Println("Slice length after appending:", len(slice))
	fmt.Println("Slice capacity after appending:", cap(slice))

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice1, slice3)

}
