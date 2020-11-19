package main

import "fmt"

func main() {
	var n, m, v int
	fmt.Scanf("%d %d %d", &n, &m, &v)

	nMap := make([][]int, n)
	for i := 0; i < m; i++ {
		var n1, n2 int
		fmt.Scanf("%d %d", &n1, &n2)

		nMap[n1][n2] = 1
		nMap[n2][n1] = 1
	}
}
