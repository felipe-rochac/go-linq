package main

import (
	"fmt"
	"golinq/src/linq"
)

func main() {
	linq := linq.From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	result := linq.Where(func(i int) bool {
		return i%2 == 0
	}).Select(func(i int) any {
		return fmt.Sprintf("Number: %d", i)
	}).Take(0, 3)
	fmt.Println(result)
}
