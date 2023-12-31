package main

/*
Given a string str, find the length of the longest substring without repeating characters.
Example:
    For “ABDEFGABEF”, the longest substring are “BDEFGA” and “DEFGAB”, with length 6.
    For “BBBB” the longest substring is “B”, with length 1.
*/
import (
	"fmt"
	"math"
	"strings"
)

func main() {
	s := "wewkew"
	fmt.Println(countMaxSubStringLenth(s))
}
func countMaxSubStringLenth(s string) int {
	size := len(s)
	if size <= 0 {
		return 0
	}
	if size == 1 {
		return 1
	}
	maxCount := 0
	subS := ""

	for i := 0; i < size; i++ {
		a := s[i]
		if strings.Contains(subS, string(a)) {
			index := strings.Index(subS, string(a))
			subS = subS[index+1 : len(subS)]
		}
		subS = fmt.Sprintf("%s%s", subS, string(a))
		maxCount = int(math.Max(float64(maxCount), float64(len(subS))))
	}
	return maxCount
}

//pwwkew

//a=p
//sub="a"
//maxCount=1

//a=w
//subS=aw
//maxCount=2

//a=w
//subS=w
//maxCount=2

//a=k
//subS=wk
//maxCount=2

//a=e
//subS=wke
//maxCount=3

//a=w
//subS=ke
//maxCount=3

// max(3,2)
