// Implementing several sorting and searching algo in go
package main

import (
	"fmt"
	"math/rand"
)

func linearSearch(arr []int, val int) int {
	for index, num := range arr {
		if num == val {
			return index
		}
	}
	return -1
}

func binarySearch(arr []int, val int) int {
	low := 0
	high := len(arr) - 1

	for low < high {
		mid := (low + high) / 2

		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			j--
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minVal := i
		for j := i; j < len(arr); j++ {
			if arr[minVal] > arr[j] {
				minVal = j
			}
		}
		arr[minVal], arr[i] = arr[i], arr[minVal]
	}
	return arr
}

func bubbleSort(arr []int) []int {
	sort := false
	size := len(arr)

	for !sort {
		swap := false
		for i := 0; i < size-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
			}
		}
		if !swap {
			sort = true
		}
		size--
	}
	return arr
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivot := rand.Intn(len(arr))

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i, _ := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	return arr
}

func factorial(val int) int {
	//testing recursion
	if val == 0 {
		return 1
	}
	return val * (factorial(val - 1))
}

func dpFibo(n int) int {
	//fibonnacci until nth sequence
	first := 0
	second := 1

	if n < 0 {
		fmt.Println("Invalid value.")
	} else if n == 1 {
		return first
	} else if n == 2 {
		return second
	} else {
		for i := 2; i < n; i++ {
			res := first + second
			first = second
			second = res
		}
	}
	return second
}

func main() {
	fmt.Println(factorial(12))
	fmt.Println("The 10th fibbonacci number is", dpFibo(15))

	list := []int{95, 78, 46, 58, 45, 86, 12, 312, 99, 251, 320}
	fmt.Println("Value found on index", linearSearch(list, 251))
	fmt.Println("Sorted list:", quickSort(list))
	fmt.Println("Value found on index", binarySearch(list, 86))
}
