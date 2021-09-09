package main

import (
	"fmt"
	"os"
	"sort"
)

func max2Numbers(arr []int) (second int) {
	if len(arr) < 2 {
		fmt.Println("Dữ liệu đầu vào chưa đúng, mảng phải có số phần tử >=2")
		os.Exit(1)
	}
	sort.Ints(arr)
	second = arr[len(arr)-2]
	return
}

func main() {
	var arr = []int{2, 1, 3, 4}
	fmt.Println(max2Numbers(arr))
}
