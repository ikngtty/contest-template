package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var x, y, z int
	fmt.Scanf("%d %d %d", &x, &y, &z)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &a[i], &b[i])
	}

	total := x + y + z
	for _, ai := range a {
		total += ai
	}
	for _, bi := range b {
		total += bi
	}

	fmt.Println(total)
}
