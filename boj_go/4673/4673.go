package main

import "fmt"

func main() {
	d := func(n int) int {
		sum := n

		for i := n; i > 0; i /= 10 {
			sum += i % 10
		}
		return sum
	}

	notSelfNum := make(map[int]bool)
	for i := 1; i <= 10000; i++ {
		notSelfNum[d(i)] = true
		if !notSelfNum[i] {
			fmt.Printf("%d\n", i)
		}
	}
}
