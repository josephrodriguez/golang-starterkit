package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println(len("Hello, World!"))
	fmt.Println("Hello, World!"[0])
	fmt.Println((true && false) || (false && true) || !(false))

	var x string
	x = "Goodbye, World!"
	fmt.Println(x)
	x = "Hello again!"
	fmt.Println(x)

	var y string = "hello"
	var z string = "world"

	fmt.Println(y == z)

	var (
		a = 5
		b = 10
		c = 15
	)

	fmt.Println("Sum:", a+b+c)

	d := 5
	d += 1

	fmt.Println("x is now", d)

	var i int = 1

	//Control structures
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	for i := 1; i <= 5; i++ {
		switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		default:
			fmt.Println("Other number")
		}
	}

	i = 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}

	var arr [10]int

	for i := range len(arr) {
		arr[i] = i * 10
	}

	fmt.Println(arr[6])

	arr1 := [4]int{34, 56, 345, 4}
	fmt.Println(arr1)

	var map1 
}
