package main

import "fmt"

func main() {
	var input []int = []int{100000, 5, 7, 9, 44, 33, 4, 6, 882222226, 453, 244, 234}
	output := MergeSort(input)
	fmt.Println(output)
}

func MergeSort(input []int) []int {
	length := len(input)
	if length < 2 {
		return input
	}

	middle := length / 2
	return merge(MergeSort(input[:middle]), MergeSort(input[middle:]))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) != 0 {
		result = append(result, left...)
	}
	if len(right) != 0 {
		result = append(result, right...)
	}
	return result
}
