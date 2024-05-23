package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Make an Array containing pseudorandom numbers in [0, max)
// The function should simply create a slice with 'numItems' entries, use a loop to
// fill in the values, and return the slice. 'rand.Intn(max)' returns a value in the
// interval [0, max)
func makeRandomArray(numItems, max int) []int {
	// Initialize a pseudorandom number generator.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed CHANGED(rolf): Author's solution with rand.Seed() is deprecated.

	arr := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		arr[i] = random.Intn(max)
	}

	return arr
}

// Print at most numItems slices.
func printArray(arr []int, numItems int) {
	if len(arr) <= numItems {
		fmt.Println(arr)
	} else {
		fmt.Println(arr[:numItems])
	}
}

// Verify that the slice is sorted.
func checkSorted(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] <= arr[i] {
			fmt.Println("The array is sorted.")
		} else {
			fmt.Println("The array is NOT sorted.")
		}
	}
}

// Use bubble sort to sort the slice.
func bubbleSort(arr []int) {
	n := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			// if this pair is out of order
			if arr[i-1] > arr[i] {
				// swap them and remember something changed
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
	}
}

func main() {
	// Get the number of items and maximum item value.
	var num_items, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&num_items)
	fmt.Printf("Maximum value: ")
	fmt.Scanln(&max)

	// Make and display the unsorted array.
	arr := makeRandomArray(num_items, max)
	printArray(arr, 50)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(arr)
	printArray(arr, 50)

	// Verify that it's sorted.
	checkSorted(arr)
}
