package main

import (
	"log"
	"go-sort-bench/pkg/sort"
)

func main() {
    arr := []int{5, 2, 9, 1, 5, 6}
    log.Println("Original array:", arr)

    sort.BubbleSort(arr)
    log.Println("Sorted array (Bubble Sort):", arr)
}