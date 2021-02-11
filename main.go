package main

import (
	"new_module/welcomepackage"
)

func main() {

	unsorted := []int{0, 1, 3, 2, 6, 5, 4, 563, 2345, 1234, 99, 23, 7, 3, 3, 4, 34, 34, 5, 6, 6}

	//welcomepackage.SelectionSort(unsorted)

	welcomepackage.BubbleSort(unsorted)

	welcomepackage.FibonacciRecursive(17)
}
