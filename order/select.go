package main

import "fmt"

func SelectOrder(l []int) {
	m, n := 0, len(l)-1

	for m < n {
		minIndex, maxIndex := m, n

		for i := m; i <= n; i++ {
			if l[i] < l[minIndex] {
				minIndex = i
				continue
			}

			if l[i] > l[maxIndex] {
				maxIndex = i
			}
		}

		l[m], l[minIndex] = l[minIndex], l[m]

		// 重新定位最大值坐标，防止重复交换
		if maxIndex == m {
			maxIndex = minIndex
		}

		l[n], l[maxIndex] = l[maxIndex], l[n]

		m++
		n--
	}
}

func main() {
	l := []int{5, 4, 3, 2, 1}
	SelectOrder(l)
	fmt.Println(l)
}
