package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Make a slice containing pseudorandom numbers in [0, max)
// The function should simply create a slice with 'numItems' entries, use a loop to
// fill in the values, and return the slice. 'rand.Intn(max)' returns a value in the
// interval [0, max)
func makeRandomSlice(numItems, max int) []int {
	// Initialize a pseudorandom number generator.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

	slice := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		slice[i] = random.Intn(max)
	}

	return slice
}

// Print at most numItems slices.
func printSlice(slice []int, numItems int) {
	if len(slice) < numItems {
		numItems = len(slice)
	}

	for i := 0; i < numItems; i++ {
		fmt.Print(slice[i], " ")
	}
}

// Verify that the slice is sorted.
func checkSorted(slice []int) {
	for i := 1; i <= len(slice); i++ {
		if slice[i-1] < slice[i] {
			fmt.Println("The slice is sorted.")
		} else {
			fmt.Println("The slice is NOT sorted.")
			
		}
	}
}

// Use bubble sort to sort the slice.
func bubbleSort(slice []int) []int {
	n := len(slice)

	for true {
		var swapped = false
		for i := 1; i < n; i++ {
			// if this pair is out of order
			if slice[i-1] > slice[i] {
				// swap them and remember something changed
				slice[i-1], slice[i] = slice[i], slice[i-1]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
	return slice
}

func main() {
	mySlice := makeRandomSlice(100, 1000)
	bubbleSort(mySlice)
	printSlice(mySlice, 50)
}
