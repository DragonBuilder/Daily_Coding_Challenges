// This problem was asked by Stripe.

// Given an array of integers, find the first missing positive integer in linear time and constant space.
// In other words, find the lowest positive integer that does not exist in the array.
// The array can contain duplicates and negative numbers as well.

// For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

// You can modify the input array in-place.
package main

func sort(L []int) {
	// O(n^2)
	for i := 0; i < len(L); i++ {
		for j := 1; j < len(L)-i; j++ {
			if L[j] < L[j-1] {
				temp := L[j]
				L[j] = L[j-1]
				L[j-1] = temp
			}
		}
	}
}

// current solution is O(n^2)
func findLowestPositive(L []int) int {
	isFirstPositive := true
	// O(n ^ 2)
	sort(L)

	// O(n)
	for i := range L {
		if L[i] < 1 {
			continue
		}
		if isFirstPositive {
			isFirstPositive = false
			continue
		}
		ep := L[i] - 1
		en := L[i] + 1

		if i > 0 {
			if L[i-1] != ep {
				return ep
			}
		}

		if i < len(L)-1 {
			if L[i+1] != en {
				return en
			}
		}
	}
	return L[len(L)-1] + 1
}
