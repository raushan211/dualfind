package main

import "fmt"

func main() {
	fmt.Println("No. of digits")
	var num int
	fmt.Scanln(&num)

	input := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Println("enter digits")
		var digits int
		fmt.Scanln(&digits)
		input[i] = digits

	}
	fmt.Println("input: ", input)
	dualFind(input)
	//fmt.Println("output: ", output)
}

func dualFind(arr []int) {
	var count int
	for i := 0; i < len(arr); i++ {
		count = 0
		for j := 0; j < len(arr); j++ {
			if i != j {
				if arr[i] == arr[j] {
					count++

				}
			}
		}

		if count == 0 {
			fmt.Println(arr[i])
		}
	}
}
