package main

import (
	"fmt"
	"sync"
)

func calculateSum(row []int, index int, ch chan<- [2]int, wg *sync.WaitGroup) {
	defer wg.Done()

	res := 0
	for _, element := range row {
		res += element
	}

	ch <- [2]int{index, res}
	// return res
}

func main() {
	slices := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	ch := make(chan [2]int, len(slices))
	var wg sync.WaitGroup

	for i, slice := range slices {
		wg.Add(1)
		go calculateSum(slice, i, ch, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(ch) // Close the channel after all goroutines finish

	rowSums := make(map[int]int)
	for i := 0; i < len(slices); i++ {
		result := <-ch

		rowSums[result[0]] = result[1]
	}

	fmt.Printf("%s", rowSums)

}
