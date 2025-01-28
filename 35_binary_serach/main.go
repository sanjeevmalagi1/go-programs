package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	target := 5

	var mid int
	l, r := 0, len(arr)-1

	for l <= r {

		mid = (l + r) / 2

		if target == arr[mid] {
			break
		} else if target < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	fmt.Println(mid)

}
