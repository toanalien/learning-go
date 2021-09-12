package main

import (
	"fmt"
	"math"
)

// ứng dụng hiển thị lịch

// tìm thứ của ngày đầu tiên của tháng
// tham khảo: https://thefullsnack.com/posts/phuc-tap-hoa-datepicker.html

func zeller(ngay, thang, nam int) int {
	if thang < 3 {
		thang += 10
	} else {
		thang -= 2
	}
	nam %= 100
	century := nam / 100
	return ((13*thang-1)/5 + nam/4 + century/4 + ngay + nam - 2*century) % 7
}

// tìm năm nhuận
// tham khảo: https://vi.wikipedia.org/wiki/nam_nhuan

func namNhuan(nam int) int {
	if nam%400 == 0 || ((nam%4 == 0) && (nam%100 != 0)) {
		return 1
	} else {
		return 0
	}
}

// tham khảo: http://www.dispersiondesign.com/articles/time/number_of_days_in_a_month

func soNgayCuaThang(thang, nam int) int {
	if thang == 2 {
		return 28 + int(namNhuan(nam))
	} else {
		return 31 - (thang-1)%7%2
	}
}

func arrLich(thang, nam int) (result []int) {
	var startIndex = zeller(1, thang, nam)
	var endIndex = soNgayCuaThang(thang, nam)
	for k, _ := range result {
		result[k] = 0
	}
	for i := 0; i < endIndex+startIndex; i++ {
		if i < startIndex {
			result = append(result, 0)
		} else {
			result = append(result, (i-startIndex)+1)
		}
	}
	return
}

func inLich(ngay int, lich []int) {
	arrThu := []string{"Chủ Nhật", "Thứ 2", "Thứ 3", "Thứ 4", "Thứ 5", "Thứ 6", "Thứ 7"}
	var result string
	for _, v := range arrThu {
		result += fmt.Sprintf("%8s |", v)
	}
	result += "\n"
	row := int(math.Ceil(float64(len(lich)) / 7))
	for i := 0; i < row; i++ {
		for j := 0; j < 7; j++ {
			pos := i*7 + j
			if i*7+j < len(lich) && lich[pos] != 0 {
				if lich[pos] == ngay {
					result += fmt.Sprintf("%8s |", fmt.Sprintf("[%d]", lich[pos]))
				} else {
					result += fmt.Sprintf("%8d |", lich[pos])
				}

			} else {
				result += fmt.Sprintf("%8s |", "")
			}
		}
		result += "\n"
	}
	fmt.Println(result)
}

func main() {
	ngay := 12
	thang := 9
	nam := 2021
	inLich(ngay, arrLich(thang, nam))
}
