package main

import "fmt"

func main() {
	prices := []int{4, 2, 1, 4, 7, 6}
	l, r := 0, 1
	maxP := 0

	for r < len(prices) {
		// check if its profitable
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			if profit > maxP {
				maxP = profit
			}
		} else {
			l = r
		}

		r += 1
	}

	fmt.Printf("left: %d right: %d \n", l, r)

	fmt.Println(maxP)
}
