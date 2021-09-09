package main

import (
	"fmt"
)

func elementInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func removeDuplicates(arr []int) (returnArr []int) {
	for _, s := range arr {
		if !elementInSlice(s, returnArr) {
			returnArr = append(returnArr, s)
		}
	}
	return
}

func main() {
	var arr = []int{1, 2, 5, 2, 6, 2, 5}
	fmt.Println(removeDuplicates(arr))
}
