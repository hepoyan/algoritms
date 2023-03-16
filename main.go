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

	array5 := []int{-9, -9, 0, 78, 800, 1880, 630, 2200, 2200}
	t, i := scr.BinarySearchRecursive(array5, 2200)
	t1, i1 := scr.BinarySearchIterative(array5, -9)
	fmt.Println(t, i)
	fmt.Println(t1, i1)

	fmt.Println(scr.FibIterative(10))
	fmt.Println(scr.FibRecursive(10))
	fmt.Println(scr.FibRecursiveOptimal(10))

	array6 := []int{5, 1, 3, 10, 4, 2, 7, -3, 8, 9}
	fmt.Println(scr.LongestIncreasingSubList(array6))

	matrix := [][]int{
		{3, 1, 2, 1},
		{1, 5, 1, 1},
		{4, 4, 4, 4},
	}
	fmt.Printf("max sum ")
	fmt.Println(scr.MaxScore(matrix))
}
