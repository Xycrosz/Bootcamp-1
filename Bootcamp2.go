package main

import "fmt"

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] < arr[j] {
				minIndex = j

			}
		}

	}
}
func main() {
	arr := []int{10, 7, 34, 97, 2, 43, 23, 13, 9, 1}
	selectionSort(arr)
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
