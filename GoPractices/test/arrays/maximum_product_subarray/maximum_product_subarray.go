package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, -2, -3, 0, 7, -8, -2}
	fmt.Println("maximum product is", maxProductSubArray(arr, len(arr)))
}

func maxProductSubArray(arr []int, size int) int {
	max_product := -999999
	if size <= 0 {
		return 0
	}
	product := 1
	for i := 0; i < size; i++ {
		product = product * arr[i]
		max_product = int(math.Max(float64(product), float64(max_product)))
		if arr[i] == 0 {
			product = 1
		}
	}
	product = 1
	for i := size - 1; i >= 0; i-- {
		product = product * arr[i]
		max_product = int(math.Max(float64(product), float64(max_product)))
		if arr[i] == 0 {
			product = 1
		}
	}
	return max_product
}
