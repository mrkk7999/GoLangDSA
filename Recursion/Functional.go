package main

import "fmt"

func SumFuncRec(N int) int {
	if N == 0 {
		return 0
	}
	return N + SumFuncRec(N-1)
}
func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(SumFuncRec(N))
}
