package main

import (
	"fmt"
)

// Tham khảo giải thuật Sàng Eratosthenes https://vi.wikipedia.org/wiki/S%C3%A0ng_Eratosthenes
//
// Pseudocode
// Input: một số nguyên n > 1
// Cho A là một mảng boolean, được đánh số từ 2 đến n,
// khởi tạo bằng cách gán tất cả phần tử trong mảng là false.
//
// for i = 2, 3, 4,..., √n:
//   if A[i] is false:
//     for j = i^2, i^2+i, i^2+2i,..., n:
//       A[j]:= true
//
// Lúc này, tất cả i ví dụ như của A[i] nếu false đều là số nguyên tố.

func main() {
	var n int
	fmt.Print("Nhập giới hạn của chuỗi số nguyên tố bạn cần tìm, N < 100,000: n = ")
	fmt.Scanln(&n)

	// Tạo slice với giá trị là false
	bools := make([]bool, n)
	// implies up to the sqrt of limit
	for k := 2; k*k <= n; k++ {
		if bools[k] == false {
			for ix := k * k; ix < n; ix += k {
				bools[ix] = true
			}
		}
	}
	// Lọc lại trong slice vị trí nào là false sẽ là số nguyên tố
	primes := []int{}
	for ix := 2; ix < n; ix++ {
		if bools[ix] == false {
			primes = append(primes, ix)
		}
	}

	fmt.Println(primes)
}
