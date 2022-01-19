package main

import (
	"fmt"
	"sort"
)

func main() {
	l := make([]int, 0, 5)
	l = append(l, 3)
	l = append(l, 2)
	sort.Ints(l)
	fmt.Println(l)
}
