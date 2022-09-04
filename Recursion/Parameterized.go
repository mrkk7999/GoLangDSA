package main

import "fmt"

func SumParamRec(N, sum int) int {
	if N < 1 {
		fmt.Println(sum)
		return sum
	}
	return SumParamRec(N-1, sum+N)
}
func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(SumParamRec(N, 0))
}
