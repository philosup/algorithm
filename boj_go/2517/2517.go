//https://www.acmicpc.net/problem/2517
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func p2517_1(n int) { //시간초과
	c := make([]int, n)

	for i := 0; i < n; i++ {
		var v int
		fmt.Scanln(&v)
		c[i] = v

		if i == 0 {
			fmt.Println(1)
			continue
		}
		min := i + 1
		for j := i - 1; j >= 0; j-- {
			if c[j] < v {
				min--
			}
		}
		fmt.Println(min)
	}
}

func LowerBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func p2517_2(n int) { //시간초과
	c := make([]int, 0, n)

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())

		lowerIndex := LowerBound(c, v)
		if lowerIndex == len(c) {
			c = append(c, v)
		} else {
			c = append(c[:lowerIndex+1], c[lowerIndex:]...)
			c[lowerIndex] = v
		}

		predictStr := strconv.Itoa(len(c) - lowerIndex)
		writer.WriteString(predictStr + "\n")
	}
}

func updateSegmentTree(value int, node int, start int, end int, _tree *[]int) int {
	tree := (*_tree)
	if start == end {
		tree[node] = 1
		return tree[node]
	}
	mid := (start + end) / 2
	if value <= mid {
		tree[node] = updateSegmentTree(value, 2*node, start, mid, _tree) + tree[2*node+1]
	} else {
		tree[node] = tree[2*node] + updateSegmentTree(value, 2*node+1, mid+1, end, _tree)
	}

	return tree[node]
}

func search(value int, node int, start int, end int, tree []int) int {
	if start == end {
		return tree[node]
	}
	mid := (start + end) / 2
	if value <= mid {
		return search(value, 2*node, start, mid, tree) + tree[2*node+1]
	} else {
		return search(value, 2*node+1, mid+1, end, tree)
	}
}

func segmentTree(n int) {
	pos := make([]int, 0, n)
	tree := make([]int, 4*n)
	normalMap := make(map[int]int)

	for i := 0; i < n; i++ {
		var v int
		fmt.Scanln(&v)
		pos = append(pos, v)
	}

	sorted := make([]int, n)
	copy(sorted, pos)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	for i := 0; i < n; i++ {
		normalMap[sorted[i]] = i + 1
	}
	for i := 0; i < n; i++ {
		updateSegmentTree(normalMap[pos[i]], 1, 1, n, &tree)
		fmt.Println(search(normalMap[pos[i]], 1, 1, n, tree))
	}
}

func segmentTreeWithBufio(n int) {
	pos := make([]int, 0, n)
	tree := make([]int, 4*n)
	normalMap := make(map[int]int)

	// readLineInt(n, &pos)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		pos = append(pos, v)
	}

	sorted := make([]int, n)
	copy(sorted, pos)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	for i := 0; i < n; i++ {
		normalMap[sorted[i]] = i + 1
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for i := 0; i < n; i++ {
		updateSegmentTree(normalMap[pos[i]], 1, 1, n, &tree)
		predict := search(normalMap[pos[i]], 1, 1, n, tree)
		predictStr := strconv.Itoa(predict)
		writer.WriteString(predictStr + "\n")
	}
}

func main() {
	var n int
	fmt.Scanln(&n)

	segmentTreeWithBufio(n)
	// p2517_2(n)
}
