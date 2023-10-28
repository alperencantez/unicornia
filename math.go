// math.go
//
// Author: Alperen Cantez <alperen.cantez@outlook.com>
// Created on: 23-10-2023

package unicornia

// Checks if <number> is in between <start> and <end>.
func MathInRange[T Number](number T, start T, end T) bool {
	if number > start && number < end {
		return true
	}

	return false
}

// Returns the mean value in a given array. Returns 0 if array length is 0.
func MathFindMean[T Number](array []T) float32 {
	length := len(array)

	if length > 0 {
		sum := ArraySum(array)
		return float32(sum) / float32(length)
	}

	return 0
}
