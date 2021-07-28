package main

import "fmt"

func main() {
	fmt.Println(solution(10))
}

func solution(n int) int {
	res := 0

	for i := 0; i <= n-1; i++ {
		if i%3 == 0 || i%5 == 0 {
			res += i
		}
	}

	return res
}
