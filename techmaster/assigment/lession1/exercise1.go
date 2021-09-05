package main

import (
	"fmt"
	"math"
)

// Giải phương trình bậc 2
// Công thức tham khảo https://freetuts.net/cach-giai-phuong-trinh-bac-hai-mot-an-3450.html

func main() {
	var a, b, c, x1, x2, delta float64
	fmt.Println("Phương trình bậc 2 có dạng: ax^2 + bx + c = 0")
	fmt.Println("Nhập giá trị a, b, c của phương trình")
	fmt.Print("a = ")
	fmt.Scanln(&a)
	fmt.Print("b = ")
	fmt.Scanln(&b)
	fmt.Print("c = ")
	fmt.Scanln(&c)

	delta = math.Pow(b, 2) - (4 * a * c)

	if delta > 0 {
		x1 = (-b + math.Sqrt(delta)) / (2 * a)
		x2 = (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Println("Phương trình có 2 nghiệm phân biệt là: x1 = ", x1, " và x2 = ", x2)
	} else if delta == 0 {
		x1 = -b / (2 * a)
		x2 = x1
		fmt.Println("Phương trình cso nghiệm kép là: x1 = ", x1, " và x2 = ", x2)
	} else if delta < 0 {
		fmt.Println("Phương trình vô nghiệm")
	}
}
