package main

import (
	"fmt"
)

func main() {
	fmt.Println("===================Dibawah Merupakan Tugas Membedakan Isi 2 Array========================")
	diffTwoArray()
	fmt.Println("=========================================================================================")
}

func diffTwoArray() {
	array1 := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	array2 := [20]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	differenceValue := []int{}
	sameValue := []int{}
	fmt.Println("Ini Merupakan Array Pertama : ", array1)
	fmt.Println("Ini Merupakan Array Kedua : ", array2)

	for a := 0; a < len(array1); a++ {
		for b := 0; b < len(array2); b++ {
			if array1[a] == array2[b] {
				sameValue = append(sameValue, array1[a])
				array1[a] = 0
				array2[b] = 0
			}
		}
	}

	fmt.Println("Nilai Yang Sama Antara Array1 dan Array2 Adalah ", sameValue)

	for i := 0; i < len(array1); i++ {
		if array1[i] != 0 {
			differenceValue = append(differenceValue, array1[i])
		}
	}
	for j := 0; j < len(array2); j++ {
		if array2[j] != 0 {
			differenceValue = append(differenceValue, array2[j])
		}
	}
	fmt.Println("Nilai Yang Berbeda Antara Array1 dan Array2 Adalah ", differenceValue)

}
