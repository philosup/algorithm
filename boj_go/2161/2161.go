package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var stack = make([]int, n)
	for i := 0; i < n; i++ {
		stack[i] = i + 1
	}

	if n == 1 {
		fmt.Printf("%d", stack[0])
		return
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", stack[0])
		stackSize := len(stack)
		stack = stack[1:stackSize]
		stackSize--
		if stackSize == 1 {
			fmt.Printf("%d", stack[0])
			break
		} else {
			stack = append(stack[1:stackSize], stack[0])
		}
	}
}
