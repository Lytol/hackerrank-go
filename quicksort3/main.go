package main

import (
	"fmt"
	"log"
)

func main() {
	list, err := Read()
	if err != nil {
		log.Fatal(err)
	}

	Quicksort(list)
}

func Read() ([]int, error) {
	var length int

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return nil, err
	}

	list := make([]int, length)

	for i := 0; i < len(list); i++ {
		_, err := fmt.Scanf("%d", &list[i])
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func Print(lists ...[]int) {
	for i := 0; i < len(lists); i++ {
		list := lists[i]

		for j := 0; j < len(list); j++ {
			fmt.Printf("%d", list[j])

			if j == len(list)-1 && i == len(lists)-1 {
				fmt.Printf("\n")
			} else {
				fmt.Printf(" ")
			}
		}
	}
}

func Quicksort(list []int) {
	quicksort(list, 0, len(list)-1)
}

func quicksort(list []int, low, high int) {
	if low < high {
		pivot := partition(list, low, high)

		Print(list)

		quicksort(list, low, pivot-1)
		quicksort(list, pivot+1, high)
	}
}

func partition(list []int, low, high int) int {
	pivot := list[high]
	index := low

	for i := low; i < high; i++ {
		if list[i] <= pivot {
			list[i], list[index] = list[index], list[i]
			index++
		}
	}

	list[index], list[high] = list[high], list[index]

	return index
}
