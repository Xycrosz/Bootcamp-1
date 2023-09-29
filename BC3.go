package main

import (
	"fmt"
	"strings"
)

type penghoetang struct {
	Nama string
	Umur int
}

func (s penghoetang) aweu() {
	fmt.Println("halo", s.Nama)
}
func (s penghoetang) icikiwir(a int) string {
	return strings.Split(s.Nama, "")[a-1]
}

func main() {
	bihom := penghoetang{
		Nama: "bihom",
		Umur: 20,
	}
	fmt.Println(bihom)

	fmt.Println("Umurnya bihom adalah:", bihom.Umur)
	hutang := map[string]string{
		"name":      "Asep",
		"alamat":    "ujungberung",
		"pekerjaan": "begal motor",
	}
	jumlah := map[string]int{
		"jumlah1": 10000,
		"tanggal": 28,
		"bulan":   9,
		"tahun":   2023,
	}
	fmt.Println(hutang)
	fmt.Println(jumlah)

	fmt.Println("Nama penghutangnya adalah:", hutang["name"], "dengan jumlah :", jumlah["jumlah1"])

	var s1 = penghoetang{"kamu", 21}
	s1.aweu()
	var s2 = penghoetang{
		Nama: "kamu",
		Umur: 18,
	}
	s2.aweu()
	var Nama = s1.icikiwir(2)
	fmt.Println("kamu dalah :", Nama)

	var numbera int = 9
	var numberbes *int = &numbera
	fmt.Println(numbera)
	fmt.Println(&numbera)
	fmt.Println(*numberbes)
	fmt.Println(numberbes)

	numbera = 1

	fmt.Println(numbera)
	fmt.Println(&numbera)
	fmt.Println(*numberbes)
	fmt.Println(numberbes)
}
