package main

import (
	"fmt"
	"study_golang/letcode"
	"time"
)

func main() {
	l1 := []int{1, 3}
	l2 := []int{2}
	start := time.Now().UnixNano()
	for i:=0; i < 20000; i++ {
		letcode.FindMedianSortedArrays2(l1, l2)
	}
	fmt.Println(time.Now().UnixNano()-start)

	start = time.Now().UnixNano()
	for i:=0; i < 20000; i++ {
		letcode.FindMedianSortedArrays1(l1, l2)
	}
	fmt.Println(time.Now().UnixNano()-start)
}
