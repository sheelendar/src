package main

import (
	"fmt"
	"math"
)

/*
Given a sequence, find the length of the longest palindromic subsequence in it.
Given sequence is “BBABCBCAB”, then the output should be 7 as “BABCBAB” is the longest palindromic subsequence in it.
“BBBBB” and “BBCBB” are also palindromic subsequences of the given sequence, but not the longest ones.
*/

func main() {
	seq := "GEEKSFORGEEKS"
	size := len(seq)
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
	}

	fmt.Println("The length of the LPS is ", longestPalindromSubSequence(seq, 0, len(seq)-1, dp))
	//displayMetrix(dp)
	fmt.Println("The length of the LPS is ", longestPalindromSubSeqByReverseString(seq, len(seq)))
	fmt.Println("The length of the LPS is ", longestPalindromSubSeq(seq, len(seq)))

}
func longestPalindromSubSeqByReverseString(seq string, size int) int {
	str1 := seq
	rev := []byte(seq)
	for i := 0; i < size/2; i++ {
		rev[i], rev[size-1-i] = rev[size-1-i], rev[i]
	}
	str2 := string(rev)
	if size == 1 {
		return 1
	}
	dp := make([][]int, size+1)
	for i := 0; i <= size; i++ {
		dp[i] = make([]int, size+1)
	}

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[size][size]
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindromSubSeq(seq string, size int) int {
	if size == 1 {
		return 1
	}
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		dp[i][i] = 1
	}
	for k := 2; k <= size; k++ {
		for i := 0; i < size-k+1; i++ {
			j := i + k - 1

			if seq[i] == seq[j] && k == 2 {
				dp[i][j] = 2
			} else if seq[i] == seq[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = int(math.Max(float64(dp[i][j-1]), float64(dp[i+1][j])))
			}
		}
	}
	return dp[0][size-1]
}

func longestPalindromSubSequence(seq string, start, end int, dp [][]int) int {
	if start == end {
		dp[start][end] = 1
		return 1
	}
	if seq[start] == seq[end] && start+1 == end {
		dp[start][end] = 2
		return 2
	}
	if dp[start][end] != 0 {
		return dp[start][end]
	}
	count := 0
	if seq[start] == seq[end] {
		count = longestPalindromSubSequence(seq, start+1, end-1, dp) + 2
		dp[start][end] = count
		return count
	}
	count = int(math.Max(float64(longestPalindromSubSequence(seq, start, end-1, dp)), float64(longestPalindromSubSequence(seq, start+1, end, dp))))
	dp[start][end] = count
	return count

}
func displayMetrix(dp [][]int) {
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp); j++ {
			fmt.Print(dp[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
