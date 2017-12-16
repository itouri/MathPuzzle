package main

// OK! 23'41"65

import (
	"fmt"
)

func digitSum(in int) int {
	sum := 0
	for in != 0 {
		sum += in % 10
		in /= 10
	}
	return sum
}

func main() {
	found := []int{}
	pre := 8
	x := 13
	for len(found) < 7 {
		if (x % digitSum(x)) == 0 {
			found = append(found, x)
		}
		tmp := x + pre
		pre = x
		x = tmp
	}
	fmt.Printf("%#v\n", found)
}
