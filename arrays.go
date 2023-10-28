// arrays.go
//
// Author: Alperen Cantez <alperen.cantez@outlook.com>
// Created on: 22-10-2023

package unicornia

import (
	"math/rand"
	"time"
)

// Returns sum of all elements in an array of integers.
func ArraySum[T Number](array []T) T {
	var arraySum T

	for _, value := range array {
		arraySum += value
	}

	return arraySum
}

// Returns index of the element in a given array or -1 if not found.
func ArrayIndexOf[E comparable](array []E, element E) int {
	for idx, e := range array {
		if e == element {
			return idx
		}
	}

	return -1
}

// Returns the first index found of <element> in <array>.
// Returns -1 if array is empty or <element> is not found in <array>.
func ArrayFindIndex[E comparable](array []E, element E) int {
	if len(array) < 1 {
		return -1
	}

	for idx, el := range array {
		if element == el {
			return idx
		}
	}

	return -1
}

// Fills elements of array with value from start up to, but not including end.
func ArrayFillInRange[E any](items []E, value E, start int, end int) {
	if start < 0 {
		start = 0
	}

	if end >= len(items) {
		end = len(items)
	}

	for i := start; i < end; i++ {
		items[i] = value
	}
}

// Eliminates duplicate elements in an array and returns it.
func ArrayUnique[E comparable](array []E) []E {
	uniqueElements := make(map[E]bool)
	res := []E{}

	for _, item := range array {
		if !uniqueElements[item] {
			uniqueElements[item] = true
			res = append(res, item)
		}
	}

	return res
}

// Concatenation is done by creating new array with the provided values.
func ArrayConcat[E any](arrayOne []E, arrayTwo []E) []E {
	result := append([]E{}, arrayOne...)
	result = append(result, arrayTwo...)

	return result
}

// Creates a slice of array with <count> elements dropped from the start.
func ArraySlice[E any](array []E, count int) []E {
	result := []E{}

	for i := count; i < len(array); i++ {
		result = append(result, array[i])
	}

	return result
}

// Opposite of filtering. Chooses elements in an array of integers where the condition is not met.
func ArrayIntReject(array []int, condition func(int) bool) []int {
	var res []int

	for _, el := range array {
		if !condition(el) {
			res = append(res, el)
		}
	}

	return res
}

// Very similar to <unicornia.ArrayIntReject()>. Works for strings.
func ArrayStringReject(array []string, condition func(string) bool) []string {
	var res []string

	for _, el := range array {
		if !condition(el) {
			res = append(res, el)
		}
	}

	return res
}

// Shuffles an array.
func ArrayShuffle[E any](array []E) []E {
	rand.NewSource(time.Now().UnixNano())

	length := len(array)

	if length < 2 {
		return array
	}

	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}

	return array
}

// Checks whether <target> is in <array> or not.
func ArrayIncludes[E comparable](array []E, target E) bool {
	for _, el := range array {
		if el == target {
			return true
		}
	}

	return false
}

// Returns all the elements in an array except the first one.
func ArrayTail[E any](array []E) []E {
	if len(array) > 0 {
		return array[1:]
	}

	return []E{}
}

// Reverses the order of a given array.
func ArrayReverse[E any](array []E) []E {
	if len(array) < 2 {
		return array
	}

	res := []E{}

	for i := len(array) - 1; i >= 0; i-- {
		res = append(res, array[i])
	}

	return res
}

// Sorts an array of integers by descending or ascending order.
// Param <order> is optional and defaults to "asc".
func ArrayIntSort(array []int, order ...string) []int {
	if len(array) < 2 {
		return array
	}

	for i := 0; i <= len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	if len(order) > 0 && order[0] == "desc" {
		return ArrayReverse(array)
	}

	return array
}

// Returns intersecting elements in given two arrays. <arrayOne> and <arrayTwo> are type sensitive meaning both must have the same type.
func ArrayIntersection[E comparable](arrayOne []E, arrayTwo []E) []E {
	res := []E{}

	for i := 0; i < len(arrayOne); i++ {
		for j := 0; j < len(arrayTwo); j++ {
			if arrayTwo[j] == arrayOne[i] {
				res = append(res, arrayOne[i])
			}
		}
	}

	return res
}
