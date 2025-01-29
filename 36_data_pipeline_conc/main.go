package main

import (
	"fmt"
	"sync"
)

func generateRandomNumbers(ch chan<- int) {
	arr := []int{12, 3, 13, 3, 1}

	for _, ele := range arr {
		ch <- ele
	}

	close(ch)
}

func filterEvenNumbers(ch <-chan int, evenInCh chan<- int) {
	for ele := range ch {
		if ele%2 == 1 { // Filtering even numbers
			evenInCh <- ele
		}
	}

	close(evenInCh)
}

func sqrNumbers(inChat <-chan int, outCh chan<- int) {
	for ele := range inChat {
		outCh <- ele * ele
	}

	close(outCh)
}

func sumNumbers(inChat <-chan int, outCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0

	for ele := range inChat {
		sum += ele
	}

	outCh <- sum

	close(outCh)
}

func main() {
	var wg sync.WaitGroup
	randomInCh := make(chan int)
	evenInCh := make(chan int)
	sqrCh := make(chan int)
	sumCh := make(chan int)

	wg.Add(1)
	go generateRandomNumbers(randomInCh)

	go filterEvenNumbers(randomInCh, evenInCh)

	go sqrNumbers(evenInCh, sqrCh)

	go sumNumbers(sqrCh, sumCh, &wg)

	for ele := range sumCh {
		fmt.Println(ele)
	}

	wg.Wait()
}
