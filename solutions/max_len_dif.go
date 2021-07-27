package main

import (
	"fmt"
	"math"
)

func main() {
	l1 := []string{"r", "df"}
	l2 := []string{"eert", "dsfdsf", "sdf"}
	fmt.Println(max_len_def(l1, l2))
}
dfgfd
func max_len_def(a1 []string, a2 []string) int {
	if (len(a1) == 0 && len(a2) == 0) || (len(a1) == 0 || len(a2) == 0) {
		return -1
	}
	res := (math.Abs(float64(len(a1)) - float64(len(a2))))
	return int(res)
}
