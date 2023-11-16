package main

import "fmt"

func containsDuplicate(nums []int) bool {
	// Create a map to simulate a set where we store unique numbers.
	hashset := make(map[int]bool)

	// Iterate over the slice of integers.
	for _, n := range nums {
		// Check if the number is already present in the set.
		if hashset[n] {
			// If a duplicate is found, return true.
			return true
		}
		// Add the number to the set.
		hashset[n] = true
	}
	// If the loop completes without finding duplicates, return false.
	return false
}

// 1. Create a map to simulate a set where we store unique numbers.
// 2. Iterate over the slice of integers.
// 3. Check if the number is already present in the set.
// 4. If a duplicate is found, return true.
// 5. Add the number to the set.
// 6. If the loop completes without finding duplicates, return false.

func containsDuplicate2(nums []int) bool {
	hashset := make(map[int]bool)

	for _, n := range nums {
		if hashset[n] {
			return true
		}
		hashset[n] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate2(nums))
}
