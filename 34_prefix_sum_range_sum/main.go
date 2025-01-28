package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 7}

	left := 1
	right := 3

	prefixSum := []int{arr[0]}

	for i := 1; i < len(arr); i++ {
		prefixSum = append(prefixSum, prefixSum[i-1]+arr[i])
	}
	fmt.Println(prefixSum)

	rangeSum := prefixSum[right] - prefixSum[left]

	fmt.Println(rangeSum)
}
