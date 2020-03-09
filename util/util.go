package util

import "fmt"

// SumArray returns the sum of all elements in a int array
func SumArray(data []int) int {
	size := len(data)
	total := 0

	for index := 0; index < size; index++ {
		total += data[index]
	}
	return total

}

// SequentialSearch returns true if the value exists in the list
func SequentialSearch(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if data[i] == value {
			return true
		}
	}
	return false
}

//BinarySearch traverse a list, looking for the passed value, in a Binary Search algo
func BinarySearch(data []int, value int) bool {
	size := len(data)
	var mid int
	low := 0
	high := size - 1
	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value { //if the value is right in the middle, returns true
			return true
		} else {
			if data[mid] < value {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

	}
	return false
}

//RotateArray rotates elements in a array k times
//Given a list, you need to rotate its elements K number of times. For example, a list [10,20,30,40,50,60] rotate by
//2 positions to [30,40,50,60,10,20]
func RotateArray(data []int, k int) []int {
	n := len(data)
	ReverseArray(data, 0, k-1)
	ReverseArray(data, k, n-1)
	ReverseArray(data, 0, n-1)
	return data

}

//ReverseArray changes positions in a given list
func ReverseArray(data []int, start int, end int) {
	i := start
	j := end

	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

//MaxSubArraySum returns max sum of elements in a subarray
//Given a list of positive and negative integers, find a contiguous subarray whose sum (sum of elements) is maximum
func MaxSubArraySum(data []int) int {
	size := len(data)
	maxSoFar := 0
	maxEndingHere := 0
	for i := 0; i < size; i++ {
		maxEndingHere = maxEndingHere + data[i]
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}

//Factorial calculates Factorials: N! = N* (N-1). It is a good example of recursive calls
func Factorial(i int) int {
	//Terminal condition
	if i <= 1 {
		return 1
	}
	//Body, Recursive Expanstion
	return i * Factorial(i-1)
}

//TOHUtil performs a Hanoi Tower algorith
func TOHUtil(num int, from string, to string, temp string) {
	if num < 1 {
		return
	}
	TOHUtil(num-1, from, temp, to)
	fmt.Println("Move disk", num, "from peg", from, "to peg", to)
	TOHUtil(num-1, temp, to, from)
}
