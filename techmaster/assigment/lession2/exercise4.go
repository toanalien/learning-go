package main

import (
	"fmt"
	"sort"
)

type nhanVien struct {
	ten        string
	heSoLuong  float64
	tienTroCap float64
}

func elInSlice(a float64, list []float64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func findSecond(arrNhanVien []nhanVien) (second float64) {
	var arrLuong []float64
	for _, s := range arrNhanVien {
		luong := s.heSoLuong*1500000 + s.tienTroCap
		if !elInSlice(luong, arrLuong) {
			arrLuong = append(arrLuong, luong)
		}
	}
	second = arrLuong[len(arrLuong)-2]
	return
}

func main() {
	arrNhanVien := []nhanVien{
		{"Toán", 1, 2},
		{"Quân", 5, 3},
		{"Chinh", 4, 2},
		{"Nam", 2, 4},
		{"Mạnh", 2, 4},
	}

	// sort tên theo bảng chứ cái tăng dần
	sort.Slice(arrNhanVien, func(i, j int) bool {
		return arrNhanVien[i].ten < arrNhanVien[j].ten
	})

	fmt.Println(arrNhanVien)

	// sort theo lương giảm dần
	sort.Slice(arrNhanVien, func(i, j int) bool {
		luong1 := arrNhanVien[i].heSoLuong*1500000 + arrNhanVien[i].tienTroCap
		luong2 := arrNhanVien[j].heSoLuong*1500000 + arrNhanVien[j].tienTroCap
		return luong1 > luong2
	})

	fmt.Println(arrNhanVien)

	// danh sách nhân viên có lương cao thứ 2
	// 1. tìm mức lương cao thứ 2
	luongCaoThu2 := findSecond(arrNhanVien)

	// 2. in ra danh sách nhân viên có lương cao thứ 2
	var arrNhanVienLuongCao2 []nhanVien
	for _, n := range arrNhanVien {
		luong := n.heSoLuong*1500000 + n.tienTroCap
		if luong == luongCaoThu2 {
			arrNhanVienLuongCao2 = append(arrNhanVienLuongCao2, n)
		}
	}

	fmt.Println(arrNhanVienLuongCao2)
}
