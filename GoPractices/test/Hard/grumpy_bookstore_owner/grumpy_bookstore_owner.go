package main

import "fmt"

/*
https://leetcode.com/problems/grumpy-bookstore-owner/description/
There is a bookstore owner that has a store open for n minutes. Every minute, some number of customers enter the store.

	You are given an integer array customers of length n where customers[i] is the number of the customer that enters the store
	 at the start of the ith minute and all those customers leave after the end of that minute.

On some minutes, the bookstore owner is grumpy. You are given a binary array grumpy where grumpy[i]
is 1 if the bookstore owner is grumpy during the ith minute, and is 0 otherwise.

When the bookstore owner is grumpy, the customers of that minute are not satisfied, otherwise, they are satisfied.

The bookstore owner knows a secret technique to keep themselves not grumpy for minutes consecutive minutes,

	but can only use it once.

Return the maximum number of customers that can be satisfied throughout the day.

Example 1:

Input: customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], minutes = 3
Output: 16
Explanation: The bookstore owner keeps themselves not grumpy for the last 3 minutes.
The maximum number of customers that can be satisfied = 1 + 1 + 1 + 1 + 7 + 5 = 16.

Example 2:

Input: customers = [1], grumpy = [0], minutes = 1
Output: 1
*/
func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	minutes := 3
	fmt.Println(maxSatisfied(customers, grumpy, minutes))
}

func maxSatisfied(cus []int, g []int, minutes int) int {
	n := len(cus)
	if n == 0 {
		return 0
	}
	max := 0
	sum := 0
	for i := 0; i < n; i++ { // sliding window find out max sum when owner is grumpy with in window of given minutes
		if g[i] == 1 {
			sum = sum + cus[i] // add customer for given minutes period
		}
		if i >= minutes && g[i-minutes] == 1 { // remove last added customers at index (i-minutes) and calculate sum
			sum = sum - cus[i-minutes]
		}
		if max < sum {
			max = sum
		}
	}
	sum = 0
	// added all element when owner is not grumpy.
	for i := 0; i < n; i++ {
		if g[i] == 0 {
			sum = sum + cus[i]
		}
	}
	return sum + max
}
