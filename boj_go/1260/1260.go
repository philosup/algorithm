//https://www.acmicpc.net/problem/1260
package main

import "fmt"

func DFS(arr [][]int, s int, n int, visited *[]int) {
	fmt.Printf("%d ", s)
	(*visited)[s] = 1
	for i := 1; i <= n; i++ {
		if arr[s][i] == 1 && (*visited)[i] == 0 {
			DFS(arr, i, n, visited)
		}
	}
}

func DFS2(arr [][]int, s int, n int) {
	visited := make([]int, n+1)
	stack := []int{}
	stack = append(stack, s)

	fmt.Printf("%d ", s)

	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		visited[pop] = 1

		notGoingAnywhere := true
		for i := 1; i <= n; i++ {
			if arr[pop][i] == 1 && visited[i] == 0 {
				visited[i] = 1
				fmt.Printf("%d ", i)
				stack = append(stack, i)
				notGoingAnywhere = false
				break
			}
		}
		if notGoingAnywhere {
			stack = stack[0 : len(stack)-1]
		}
	}
	fmt.Println()
}

func BFS(arr [][]int, s int, n int) {
	queue := []int{}
	visited := make([]int, n+1)

	queue = append(queue, s)
	for len(queue) > 0 {
		first := queue[0]
		fmt.Printf("%d ", first)
		visited[first] = 1
		queue = queue[1:]
		for i := 1; i <= n; i++ {
			if visited[i] != 1 && arr[first][i] == 1 {
				visited[i] = 1
				queue = append(queue, i)
			}
		}
	}
}

func PrintVisit(visited []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("%d ", visited[i])
	}
	fmt.Println()
}

func main() {
	var n, m, v int
	fmt.Scanf("%d %d %d", &n, &m, &v)

	nMap := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		nMap[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		var n1, n2 int
		fmt.Scanf("%d %d", &n1, &n2)

		nMap[n1][n2] = 1
		nMap[n2][n1] = 1
	}

	// visited := make([]int, n+1)
	// DFS(nMap, v, n, &visited)
	DFS2(nMap, v, n)

	BFS(nMap, v, n)
}
