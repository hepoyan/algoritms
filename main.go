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

	array7 := []int{2, 1, 4, 3, 5, 10, 12}
	fmt.Println(scr.LongestIncreasingSubsequenceOptimal(array7))
	fmt.Println(scr.LongestIncreasingSubsequence(array7))

	fmt.Println(scr.LCS("aaab", "abaa"))

	fmt.Println(scr.RedDistance("abc", "adck"))

	c := []int{55, 10, 47, 5, 4, 50, 8, 61, 85, 87}
	weight := []int{95, 4, 60, 32, 23, 72, 80, 62, 65, 46}
	/*	c := []int{5, 11, 13}
		weight := []int{2, 4, 6}*/
	fmt.Println(scr.GreedyKnapsack(269, weight, c))

	fmt.Println(scr.Knapsack(269, weight, c))

	fmt.Println(scr.ExtKnapsack(269, weight, c))

	graph := make(scr.Graph, 8)
	scr.CreateVertex(0, graph)
	scr.CreateVertex(1, graph)
	scr.CreateVertex(2, graph)
	scr.CreateVertex(3, graph)
	scr.CreateVertex(4, graph)
	scr.CreateVertex(5, graph)
	scr.CreateVertex(6, graph)
	scr.CreateVertex(7, graph)
	scr.CreateVertex(8, graph)

	scr.CreateEdge(0, 1, graph)
	scr.CreateEdge(0, 4, graph)
	scr.CreateEdge(0, 7, graph)
	scr.CreateEdge(1, 2, graph)
	scr.CreateEdge(1, 3, graph)
	scr.CreateEdge(2, 3, graph)
	scr.CreateEdge(4, 5, graph)
	scr.CreateEdge(5, 6, graph)
	used := make([]bool, 8)
	for i2 := range used {
		used[i2] = false
	}
	fmt.Println(graph)
	fmt.Printf("DFS")
	//scr.Dfs(graph, 0, used)

	scr.Bfs(graph)

	graph1 := make(scr.Graph, 5)
	scr.CreateVertex(0, graph1)
	scr.CreateVertex(1, graph1)
	scr.CreateVertex(2, graph1)
	scr.CreateVertex(3, graph1)
	scr.CreateVertex(4, graph1)

	scr.CreateEdge(0, 1, graph1)
	scr.CreateEdge(0, 2, graph1)
	scr.CreateEdge(1, 2, graph1)
	scr.CreateEdge(2, 0, graph1)
	scr.CreateEdge(2, 3, graph1)

	/*	var stack = list.New()
		used1 := make([]bool, 5)
		for i2 := range used1 {
			used1[i2] = false
		}
		fmt.Println("----------------------")
		for v := 0; v < 5; v++ {
			if !used1[v] {
				scr.DfsForSort(graph1, v, used1, stack)
			}
		}
		for e := stack.Front(); e != nil; e = e.Next() {
			fmt.Printf("%d ", e.Value.(int)+1)
		}*/
	stack := scr.NewStack()
	used1 := make([]bool, 5)
	for i2 := range used1 {
		used1[i2] = false
	}
	fmt.Println("----------------------")
	for v := 0; v < 5; v++ {
		if !used1[v] {
			scr.DfsForSort(graph1, v, used1, stack)
		}
	}
	fmt.Println(stack)
	//	&{[4 3 2 1 5] 4}

	/*for e := stack.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value.(int)+1)
	}*/

}

/*
8 8
0 1
0 4
0 7
1 2
1 3
2 3
4 5
5 6
*/
