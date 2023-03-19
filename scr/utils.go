package scr

import (
	"container/list"
	"fmt"
	"sort"
)

// GetPrimeNumbers O(n * sqrt(n))
func GetPrimeNumbers(n int) []int {
	var resArr []int
	if n == 1 {
		return resArr
	}
	if n == 2 {
		resArr = append(resArr, 2)
		return resArr
	}

	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			resArr = append(resArr, i)
		}
	}
	return resArr
}

// GetPrimeNumbersOpt O(n * log(log n))
func GetPrimeNumbersOpt(n int) []int {
	var isPrimeList []bool

	isPrimeList = append(isPrimeList, false)
	isPrimeList = append(isPrimeList, false)

	for i := 2; i < n+1; i++ {
		isPrimeList = append(isPrimeList, true)
	}

	for i := 2; i < n; i++ {
		if isPrimeList[i] == true {
			for j := i * i; j < n+1; j += i {
				isPrimeList[j] = false
			}
		}
	}

	var resArr []int

	for i := 0; i < len(isPrimeList); i++ {
		if isPrimeList[i] {
			resArr = append(resArr, i)
		}
	}
	return resArr
}

// Sort algorithms

// BubbleSort O(n^2)
func BubbleSort(inputList []int) {
	for i := 0; i < len(inputList); i++ {
		isSorted := true
		for j := 0; j < len(inputList)-1; j++ {
			if inputList[j] > inputList[j+1] {
				inputList[j], inputList[j+1] = inputList[j+1], inputList[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
	//	return inputList
}

// MergeSort O(log n! = n*log n)
func merge(leftList []int, rightList []int) []int {
	var res []int
	i := 0
	j := 0
	for i < len(leftList) && j < len(rightList) {
		if leftList[i] < rightList[j] {
			res = append(res, leftList[i])
			i++
		} else {
			res = append(res, rightList[j])
			j++
		}
	}
	for i < len(leftList) {
		res = append(res, leftList[i])
		i++
	}
	for j < len(rightList) {
		res = append(res, rightList[j])
		j++
	}
	return res
}

func MergeSort(inputList []int) []int {
	if len(inputList) <= 1 {
		return inputList
	}
	leftList := MergeSort(inputList[:len(inputList)/2])
	rightList := MergeSort(inputList[len(inputList)/2:])
	return merge(leftList, rightList)
}

func minElement(inputList []int) int {
	min := inputList[0]
	for i := 0; i < len(inputList); i++ {
		if min > inputList[i] {
			min = inputList[i]
		}
	}
	return min
}

func maxElement(inputList []int) int {
	max := inputList[0]
	for i := 0; i < len(inputList); i++ {
		if max < inputList[i] {
			max = inputList[i]
		}
	}
	return max
}

// CountingSort O(n + k) k = maxElement - minElement | don't work in double or string only working integers
func CountingSort(inputList []int) []int {
	if len(inputList) == 0 {
		return inputList
	}

	min := minElement(inputList)
	max := maxElement(inputList)
	counts := make([]int, max-min+1)
	res := make([]int, len(inputList))

	for i := 0; i < len(inputList); i++ {
		counts[inputList[i]-min]++
	}
	k := 0
	for i := 0; i < len(counts); i++ {
		for counts[i] != 0 {
			res[k] = i + min
			k++
			counts[i]--
		}
	}
	return res
}

// RadixSort O(d*(n+b)) b = 10 d = log b (k)  only list element  >= 0
/* 10 317 7 2 2728 4045
I	0 7 5 7 2 8 5
II	1 1 2 0 0 2 4
III	0 3 0 0 0 7 0
IV	0 0 0 0 0 2 4
I  -> counting sor
II -> counting sor
III-> counting sor
VI -> counting sor
*/
func RadixSort(inputList []int) {
	if len(inputList) == 0 {
		return
	}
	max := maxElement(inputList)
	for exp := 1; max/exp > 0; exp *= 10 {
		countingSOrtForDigit(inputList, exp)
	}
}

func countingSOrtForDigit(inputList []int, exps int) {

	count := make([]int, 10)
	for i := 0; i < len(inputList); i++ {
		count[inputList[i]/exps%10]++
	}

	countP := make([]int, 10)
	for i := 1; i < len(countP); i++ {
		countP[i] = countP[i-1] + count[i-1]
	}

	var index int

	sorted := make([]int, len(inputList))
	for i := 0; i < len(inputList); i++ {
		index = countP[inputList[i]/exps%10]
		countP[inputList[i]/exps%10] = countP[inputList[i]/exps%10] + 1
		sorted[index] = inputList[i]
	}

	for i := 0; i < len(inputList); i++ {
		inputList[i] = sorted[i]
	}

}

// QuickSort Disadvantages O(n^2), Summary O(n log n).
// 5 -7 2 14 3 -9 1 [0] - k
// index = 0
// v[i] < pivot
// swap(v[index],v[i])
// ++index
// I Step
// 5 -7 2 14 3 -9 1 0
// -7 -9 2 14 3 5 1 0
// -7 -9 0 14 3 5 1 2
// ...
func QuickSort(inputList []int) {
	quickSortImpl(inputList, 0, len(inputList))
}

func partition(inputList []int, left int, right int) int {
	pivot := inputList[right-1]
	pivotIndex := left
	for i := left; i < right-1; i++ {
		if pivot > inputList[i] {
			inputList[i], inputList[pivotIndex] = inputList[pivotIndex], inputList[i]
			pivotIndex++
		}
	}
	inputList[pivotIndex], inputList[right-1] = inputList[right-1], inputList[pivotIndex]
	return pivotIndex
}

func quickSortImpl(inputList []int, l int, r int) { //[l,r)
	if r-l <= 1 {
		return
	}
	pi := partition(inputList, l, r)
	quickSortImpl(inputList, l, pi)
	quickSortImpl(inputList, pi+1, r)
}

//Search algorithms

// BinarySearchRecursive O(log n)
func BinarySearchRecursive(inputSortedList []int, targetElement int) (bool, int) {
	return binarySearchImp(inputSortedList, 0, len(inputSortedList), targetElement)
}

func binarySearchImp(inputList []int, l int, r int, targetElement int) (bool, int) {
	if l == r {
		return false, -1
	}
	mid := (l + r) / 2
	if targetElement == inputList[mid] {
		return true, mid
	}
	if targetElement < inputList[mid] {
		return binarySearchImp(inputList, l, mid, targetElement)
	} else {
		return binarySearchImp(inputList, mid+1, r, targetElement)
	}
}

func BinarySearchIterative(inputSortedList []int, lookingFor int) (bool, int) {
	l := 0
	r := len(inputSortedList) - 1

	var mid int
	var midValue int
	for l <= r {
		mid = l + (r-l)/2
		midValue = inputSortedList[mid]
		if midValue == lookingFor {
			return true, mid
		} else if midValue > lookingFor {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false, -1
}

// FibRecursive approx O(2^n)
func FibRecursive(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

// FibRecursiveOptimal  O(n)
func FibRecursiveOptimal(n int) int {
	fb := make([]int, n+1)
	if n == 1 || n == 0 {
		fb[n] = 1
		return fb[n]
	}
	if fb[n] != 0 {
		return fb[n]
	}
	fb[n] = FibRecursive(n-1) + FibRecursive(n-2)
	return fb[n]
}

// FibIterative O(n)
func FibIterative(n int) []int {
	fb := make([]int, n+1)
	fb[0] = 1
	fb[1] = 1
	for i := 2; i <= n; i++ {
		fb[i] = fb[i-1] + fb[i-2]
	}
	return fb
}

func maximum(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Dynamic Programming

// LongestIncreasingSubList find the longest substring element (not separately)
// example 3, 1, 2, 5, 0   result is 3 |(1 2 5)
func LongestIncreasingSubList(inputList []int) int {
	var subList []int
	if len(inputList) == 0 {
		return 0
	}
	maxLen := 1
	current := 1
	for i := 0; i < len(inputList)-1; i++ {
		if inputList[i] < inputList[i+1] {
			subList = append(subList, inputList[i], inputList[i+1])
			current++
		} else {
			maxLen = maximum(maxLen, current)
			current = 1
		}

	}
	maxLen = maximum(maxLen, current)
	return maxLen
}

// Dynamic Programming
//we have a board, we need to reach from the upper left corner to the lower right corner.
//In the columns are written units and must collect the maximum
// 3 1 2 77
// 1 5 1 1
// 4 4 4 4
// dp[i][j] shows how many points A[i][j] will score by landing on the box
// dp[i][j] = max(dp[i - 1][j] ,dp[i][j - 1]) + A[i][j]

func MaxScore(a [][]int) int {
	n := len(a)
	m := len(a[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = a[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + a[i][0]
	}

	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + a[0][j]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = maximum(dp[i-1][j], dp[i][j-1]) + a[i][j]
		}
	}

	return dp[n-1][m-1]
}

// Dynamic Programming

// LongestIncreasingSubsequence (LIS) find the longest substring element (separately)
// example 2, 1, 4, 3, 5, 10, 12  result is 5 |(1, 4, 5,10, 12)
// dp[i] = max(dp[i] + 1)
//
//	  0 <=j < i
//	inputList[i] > inputList[j]
//
// result max(dp[i])
// O(n^2)
func LongestIncreasingSubsequence(input []int) int {
	n := len(input)
	dp := make([]int, n)
	for index := range dp {
		dp[index] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if input[i] > input[j] {
				dp[i] = maximum(dp[j]+1, dp[i])
			}
		}
	}
	return maxElement(dp)
}

// LongestIncreasingSubsequenceOptimal O(nlog n)
func LongestIncreasingSubsequenceOptimal(inputList []int) int {
	res := make([]int, len(inputList))
	res[0] = 1
	dp := make([]int, 1)
	dp[0] = inputList[0]
	for i := 1; i < len(inputList); i++ {
		if inputList[i] > dp[len(dp)-1] {
			dp = append(dp, inputList[i])
			res[i] = len(dp)
		} else {
			k := -1
			for j := len(dp) - 2; j > -1; j-- {
				if dp[j] < inputList[i] {
					k = j
					break
				}
			}
			dp[k+1] = inputList[i]
			res[i] = k + 2
		}
	}
	return maxElement(res)
}

// LCS
// a a b a second s1
// a b a second d s2
// result is  3 | a b second
// dp[i][j]
// dp[n][m]
//
//	s1[i] = s2[j] dp[i - 1][j - 1] + 1
//
// dp[i][j] ||
//
//	!= max(dp[i - 1][j], dp[i][j - 1])
//
// res is dp[n][m]
// example
//
//	""a d second
//	0 1 2 3
//
// i"" 0 0 0 0 0
//
//	a	1 0 1 1 1
//	b	2 0 1 1 2
//	second	3 0 1 1 2
//
// O(n*m)
func LCS(s1, s2 string) int {
	n := len(s1)
	m := len(s2)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maximum(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

// we have 2 different strings,
// the minimum number of steps through which
// we can get the second from the first
//
//	s[i] delete	dp[i - 1][j] + 1
//
// dp[i][j] = ||  t[j] add      dp[i][j - 1] + 1
//
//	s[i] = t[j]	dp[i - 1][j - 1]
//	!= change 	dp[i - 1][j - 1] +1
//
// example  a b second
//
//	a d second k
//
// a b second -> change b
// res a d second
// res -> add k
// a d second k
// use in correct the error when writing something on the phone

func RedDistance(s1, s2 string) int {
	n := len(s1)
	m := len(s2)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < n+1; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < m+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			case1 := dp[i-1][j] + 1
			case2 := dp[i][j-1] + 1
			case3 := dp[i-1][j-1]
			if s1[i-1] != s2[j-1] {
				case3++
			}
			dp[i][j] = case1
			if case2 < dp[i][j] {
				dp[i][j] = case2
			}
			if case3 < dp[i][j] {
				dp[i][j] = case3
			}
		}
	}

	return dp[n][m]
}

// GreedyKnapsack
// We are given N items where each item has some weight and profit associated with it.
// We are also given a bag with capacity W, [i.e., the bag can hold at most W weight in it].
// The target is to put the items into the bag such that the sum of profits associated with them is the maximum possible.
// O(n * W)
func GreedyKnapsack(knapsack int, w []int, c []int) float64 {
	n := len(w)
	cw := make([]Pair, n)
	for i := 0; i < n; i++ {
		cw[i] = Pair{Wight: w[i], Price: c[i]}
	}
	sort.Slice(cw, func(i, j int) bool {
		return cw[i].Price*cw[j].Wight > cw[j].Price*cw[i].Wight
	})
	var cost float64 = 0
	i := 0
	for knapsack > 0 && n != i {
		if cw[i].Wight <= knapsack {
			cost += float64(cw[i].Price)
			knapsack -= cw[i].Wight
		} else {
			cost += float64(cw[i].Price*knapsack) / float64(cw[i].Wight)
			break
		}
		i++
	}
	return cost
}

// Knapsack Note: The constraint here is we can either put an item completely into the bag or cannot put it at all [It is not possible to put a part of an item into the bag].
// dp[i] = max{dp[i - w[j]] != 1, dp[i - w[j]] + c[j]}
//
//	1<j<n
//
// O(n * W)
func Knapsack(knapsack int, w []int, c []int) int {

	dp := make([]int, knapsack+1)
	for i := 0; i < knapsack; i++ {
		dp[i] = -1
	}
	n := len(w)
	dp[0] = 0

	for i := 0; i < n; i++ {
		for j := knapsack; j >= 0; j-- {
			if j-w[i] >= 0 && dp[j-w[i]] != -1 {
				dp[j] = maximum(dp[j], dp[j-w[i]]+c[i])
			}
		}
	}

	return maxElement(dp)
}

// ExtKnapsack use each element as much as we want
// O(n * W)
func ExtKnapsack(knapsack int, w []int, c []int) int {

	dp := make([]int, knapsack+1)
	for i := 0; i < knapsack; i++ {
		dp[i] = -1
	}
	n := len(w)
	dp[0] = 0
	for j := 0; j < n; j++ {
		for i := w[j]; i <= knapsack; i++ {
			if dp[i-w[j]] != -1 {
				dp[i] = maximum(dp[i], dp[i-w[j]]+c[j])
			}
		}

	}
	//fmt.Println(dp)
	return maxElement(dp)
}

// NodeDfs
// G = (V,E)
// V - a set of vertices E -  a set of edges
// undirected  graph is a set of objects (called vertices or nodes) that are connected together, where all the edges are bidirectional.
// undirected  graph can be represented as a symmetric matrix
// count vertices is n => count edges = n(n - 1)/2
// matrix representation is not good,

func CreateVertex(value int, graph Graph) {
	if graph[value] == nil {
		graph[value] = append(graph[value], value)
	} else {
		fmt.Println("vertex already created.")
	}
}

func CreateEdge(node1 int, node2 int, graph Graph) {
	graph[node1] = append(graph[node1], node2)
	graph[node2] = append(graph[node2], node1)
}

//DFS & BFS

// Dfs  - Depth First Traversal (or Search) for a graph is similar to the Depth First Traversal of a tree.
// The only catch here is, that, unlike trees, graphs may contain cycles (a node may be visited twice).
// To avoid processing a node more than once, use a boolean visited array.
// A graph can have more than one DFS traversal.
// In algorithm has 3 case (observed, observable,unobserved,)
// O(V + E)

func Dfs(graph Graph, v int, used []bool) {
	used[v] = true
	/*	for u := 0; u < len(graph[v]); u++ {
		if !used[graph[v][u]] {
			Dfs(graph, graph[v][u], used)
		}
	}*/
	for _, u := range graph[v] {
		if !used[u] {
			Dfs(graph, u, used)
		}
	}
	fmt.Println(v + 1)
}

// Bfs The breadth-first search (BFS) algorithm is used to search a tree or graph data structure for a node that meets a set of criteria.
// It starts at the tree’s root or graph and searches/visits all nodes at the current depth level before moving on to the nodes at the next depth level.
// Breadth-first search can be used to solve many problems in graph theory.
// O(V + E)

func Bfs(graph Graph) {
	fmt.Printf("BFS  ")
	used := make([]bool, len(graph))
	for i := range used {
		used[i] = false
	}
	used[0] = true
	queue := list.New()
	queue.PushBack(0)

	for queue.Len() > 0 {
		v := queue.Front() // First element

		queue.Remove(v)
		for _, u := range graph[(v.Value.(int))] {
			if !used[u] {
				used[u] = true
				queue.PushBack(u)
			}
		}
		fmt.Println(v.Value.(int) + 1)
	}
}

//Directed acyclic graph (DAC)
//Topology sort in Graph
//Topological sorting for Directed Acyclic Graph (DAG) is a linear ordering of vertices such that for every directed edge uv, vertex u comes before v in the ordering.
//Topological Sorting for a graph is not possible if the graph is not a DAG.

func DfsForSort(graph Graph, v int, used []bool, sorted *Stack) {
	used[v] = true
	for _, u := range graph[v] {
		if !used[u] {
			DfsForSort(graph, u, used, sorted)
		}

	}

	sorted.Push(v + 1)

}
