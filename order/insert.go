package main

import "fmt"

func InsertOrder(l []int) {
	for i := 1; i < len(l); i++ {
		tmp := l[i]
		j := i - 1
		for ; j >= 0; j-- {
			if l[j] > tmp {
				l[j+1] = l[j]
				l[j] = tmp
			}
		}
	}
}

func main() {
	l := []int{5, 4, 3, 2, 1}
	InsertOrder(l)
	fmt.Println(l)
}
