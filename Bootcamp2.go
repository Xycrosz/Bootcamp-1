package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {

			fmt.Println(arr)
			fmt.Println("Element yang ditukar yaitu ", arr[j], "dengan", arr[j+1])
			//elemeent di bawah
			if arr[j] > arr[j+1] {
				//yang jadi dari gede ke kecil ato kecil ke gede itu diatas tanda lebih darinya:)))))))))))))
				arr[j], arr[j+1] = arr[j+1], arr[j]

			} else {
				fmt.Println("tidak ada pennukaran ")
			}
		}
	}
}

func main() {
	arr := []int{20, 14, 17, 47, 3, 53, 73, 63, 29, 1}
	bubbleSort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Print("\n")

	names := [...]string{"Eko", "Julianto", "Farhab", "Ebid", "Dangdut", "Aweu"}
	slice := names[:]
	slice1 := append(slice, "cikini", "surabaya", "bandung")

	fmt.Println(slice1)
	fmt.Println(len(arr))
}
