package scr

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
