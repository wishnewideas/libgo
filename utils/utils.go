package utils

import (
	"fmt"
)

// FibonacciRecursive recursive implementation of Fibonacci
func FibonacciRecursive(n int) {
	arr := []int{}

	for i := 0; i < n; i++ {

		arr = append(arr, fibonacci(i))
	}
	printArr(arr)
}

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

// FactorialOf simple factorial
func FactorialOf(n int) (result int) {
	result = 1
	if n == 1 {
		return
	}

	return n * FactorialOf(n-1)
}

// FactorialOf Experimenting with recursion
// func factorialOf(n int) {
// 	arr := []int{}

// 	for i := 0; i < n; i++ {

// 		arr = append(arr, Factorial(i))
// 	}
// 	printArr(arr)
// }

// SelectionSort first sorting algorithm
func SelectionSort(unsortedArr []int) {
	arrLength := len(unsortedArr)
	unsortedArrStartIndex := 0
	indexOfLowestValue := 0
	lowestValChanged := false

	for ; unsortedArrStartIndex < arrLength-1; unsortedArrStartIndex++ {
		lowestValChanged = false
		indexOfLowestValue = unsortedArrStartIndex

		// find smallest value
		for i := unsortedArrStartIndex; i < arrLength; i++ {
			if unsortedArr[i] < unsortedArr[indexOfLowestValue] {
				indexOfLowestValue = i
				lowestValChanged = true
			}
		}

		// swap
		if lowestValChanged {
			unsortedArr = swapElementsIntArr(unsortedArr, unsortedArrStartIndex, indexOfLowestValue)
		}
		// redo all steps and incriment the starting index of the unordered section of the arr
	}

	printArr(unsortedArr)

}

// BubbleSort second sorting algorithm
func BubbleSort(arr []int) {
	len := len(arr)
	swapcounter := -1

	for swapcounter != 0 {
		swapcounter = 0

		for i := 0; i < len-1; i++ {
			if arr[i] > arr[i+1] {
				arr = swapElementsIntArr(arr, i, i+1)
				swapcounter++
			}
		}

	}

	printArr(arr)

}

// MergeSort third sorting algorithm, more complex and uses recursion
func MergeSort(arr []int, leftBound int, rightBound int) {

	mid := leftBound + (rightBound-1)/2

	if leftBound < rightBound {
		// first sort left half of the arr
		MergeSort(arr, leftBound, mid)
		// sort right half of the arr
		MergeSort(arr, mid+1, rightBound)
		// merge two halves together

	}
	merge(arr, leftBound, mid, rightBound)
}

func merge(arr []int, leftBound int, mid int, rightBound int) {
	var i, j, k int
	n1 := mid - leftBound + 1
	n2 := rightBound - mid
	var leftArr, rightArr []int

	for i = 0; i < n1; i++ {
		leftArr[i] = arr[leftBound+i]
	}

	for j = 0; j < n2; j++ {
		rightArr[j] = arr[mid+1+j]
	}

	i = 0
	j = 0
	k = leftBound

	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

func printArr(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		if i < length-1 {
			fmt.Print(arr[i], ", ")
		} else {
			fmt.Print(arr[i])
		}

	}
	fmt.Println()
}

func swapElementsIntArr(unsortedArr []int, index1 int, index2 int) []int {
	tmp := unsortedArr[index1]
	unsortedArr[index1] = unsortedArr[index2]
	unsortedArr[index2] = tmp

	return unsortedArr
}
