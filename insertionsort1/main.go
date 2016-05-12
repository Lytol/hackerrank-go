package main

import (
	"fmt"
	"log"
)

func main() {
	list, err := readList()
	if err != nil {
		log.Fatal(err)
	}

	InsertionSort(list)
}

func readList() ([]int, error) {
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

func InsertionSort(list []int) {
	value := list[len(list)-1]

	for i := len(list) - 1; i > 0; i-- {
		list[i] = list[i-1]

		if list[i] < value {
			list[i] = value
			Print(list)
			return
		}
		Print(list)
	}

	list[0] = value
	Print(list)
}

func Print(list []int) {
	for i := 0; i < len(list); i++ {
		if i != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", list[i])
	}
	fmt.Printf("\n")
}
