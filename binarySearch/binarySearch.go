package main

import (
	"cmp"
	"log"
)

func Search[K cmp.Ordered](arr []K, val K) bool {
	return search(arr, val, 0, len(arr))
}

func search[K cmp.Ordered](arr []K, val K, lo int, hi int) bool {
	for ok := true; ok; ok = (lo < hi) {
		mid := lo + (hi-lo)/2
		midVal := arr[mid]

		if midVal == val {
			return true
		} else if midVal > val {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return false
}

func main() {
	testSearch := Search([]int{0, 1, 2, 3, 4, 5, 6}, -5)
	log.Println(testSearch)
}
