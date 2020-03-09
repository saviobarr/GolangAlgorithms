package main

import (
	"fmt"

	"github.com/saviobarr/book_ds_and_algos_in_go/util"
)

func main() {
	data := []int{10, 20, 30, 40, 50, 60}

	fmt.Println(util.SumArray(data))
	fmt.Println(util.SequentialSearch(data, 11))

	fmt.Println(util.BinarySearch(data, 10))

	fmt.Println(util.RotateArray(data, 2))

	data = []int{1, -2, 3, 4, -4, 6, -14, 8, 2}
	fmt.Println("Max sub array sum: ", util.MaxSubArraySum(data))
}
