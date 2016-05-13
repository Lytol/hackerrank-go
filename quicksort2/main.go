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

	partition(list)
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

func partition(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	pivot := list[0]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 1; i < len(list); i++ {
		if list[i] <= pivot {
			left = append(left, list[i])
		} else {
			right = append(right, list[i])
		}
	}

	sorted := append(append(partition(left), pivot), partition(right)...)

	Print(sorted)

	return sorted
}

func Divide(list []int, pivot int) ([]int, []int, []int) {
	left, equal, right := make([]int, 0), make([]int, 0), make([]int, 0)

	for i := 0; i < len(list); i++ {
		switch {
		case list[i] == list[pivot]:
			equal = append(equal, list[i])
		case list[i] < list[pivot]:
			left = append(left, list[i])
		case list[i] > list[pivot]:
			right = append(right, list[i])
		}
	}

	return left, equal, right
}
