package slices

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func Filter[T constraints.Ordered](nums []int, cb func(T) bool) {
	result := make([]T, 0, len(nums))
	for _, num := range nums {
		if cb(num) {
			result = append(result, num)
		}
	}
	fmt.Println(result)
}
