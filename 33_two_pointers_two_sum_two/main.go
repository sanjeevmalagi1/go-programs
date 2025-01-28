package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	target := 6

	l := 0
	r := len(arr) - 1

	for l < r {
		// get sum
		sum := arr[l] + arr[r]
		fmt.Println(sum)

		if sum == target {
			break
		} else if sum < target {
			l += 1
		} else {
			r -= 1
		}

	}

	fmt.Printf("res[0]: %d res[1]: %d\n", arr[l], arr[r])

}
