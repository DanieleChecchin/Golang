package main

import "fmt"

func main() {

	ages := [4]int{15, 18, 20, 22}

	fmt.Println(ages, len(ages))

	names := [3]string{"Alice", "Bob", "Charlie"}

	fmt.Println(names, len(names))

	var numbers = []int{1, 2, 3}
	numbers[0] = 7
	numbers = append(numbers, 8, 9)

	fmt.Println(numbers, len(numbers))
}
