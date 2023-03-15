package main

import (
	"algoritms/scr"
	"fmt"
)

func main() {
	fmt.Println(scr.GetPrimeNumbers(11))
	fmt.Println(scr.GetPrimeNumbersOpt(11))

	array := []int{10, 6, 2, 1, 5, 8, 3, 4, 778, 96}
	scr.BubbleSort(array)

	array1 := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	fmt.Println(array)
	array1 = scr.MergeSort(array1)

	fmt.Println(array1)
	array2 := []int{-10, 10, 8, 9, 2, 0, 0, 0, -6}
	fmt.Println(scr.CountingSort(array2))

	array3 := []int{10, 317, 15, 7, 2, 2873, 4045}
	scr.RadixSort(array3)
	fmt.Println(array3)

	array4 := []int{5, -7, 2, 14, 3, -9, 1, 0}
	scr.QuickSort(array4)
	fmt.Println(array4)
}
