package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Chương trình sinh ra một số nguyên dương ngẫu nhiên x thoả 0 <= x <= 100
// Nhận input từ người dùng là 1 số dạng int, lưu vào y
// Nếu y > x: In ra Số bạn đoán lớn hơn X
// Nếu y < x: In ra Số bạn đoán lớn hơn X
// Nếu y == x: In ra Bạn đã đoán đúng
// Lặp lại chương trình đến khi nào user nhập đúng
//
// Tham khảo cách dùng Rand ở đây https://pkg.go.dev/math/rand
func main() {
	rand.Seed(time.Now().UnixNano())
	var x, y int
	min := 0
	max := 100
	fmt.Println("Chương trình đoán số ngẫu nhiên. Nhấn Control + C để thoát.")
	for {
		x = rand.Intn(max-min+1) + min

		fmt.Print("Số bạn đoán là? ")
		fmt.Scanln(&y)

		if y > x {
			fmt.Println("Số bạn đoán lớn hơn X")
		} else if y < x {
			fmt.Println("Số bạn đoán nhỏ hơn X")
		} else {
			fmt.Println("Bạn đã đoán đúng")
			os.Exit(0)
		}
	}
}
