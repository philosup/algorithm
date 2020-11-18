package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		var vps string
		fmt.Scanln(&vps)

		var stack []int
		for j := 0; j < len(vps); j++ {
			if vps[j] == '(' { //push
				stack = append(stack, 1)
			} else if vps[j] == ')' { //pop
				var stackSize = len(stack)
				if stackSize == 0 { //
					stack = append(stack, 1)
					break
				} else {
					stack = stack[0 : stackSize-1]
				}
			}
		}
		if len(stack) == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
