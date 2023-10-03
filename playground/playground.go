package main

import (
	"log"
)

func mutateSlice(path *[]int) {
	*path = append(*path, 5)
	*path = append(*path, 5)
	*path = append(*path, 5)

	*path = (*path)[:len(*path)-1]
}

func main() {
	myPath := make([]int, 0)
	mutateSlice(&myPath)

	log.Println(myPath)
}
