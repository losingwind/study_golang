package main

func ShellOrder(l []int) {
	for d := len(l) / 2; d > 0; d /= 2 {
		for i := 1; i*d < len(l); i++ {
			tmp := l[i*d]

			for j := i - 1; j >= 0; j-- {

			}
		}
	}
}
