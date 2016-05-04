package main

import (
	"fmt"
	"log"
)

func main() {
	length, err := readLength()
	if err != nil {
		log.Fatal(err)
	}

	values, err := readValues(length)
	if err != nil {
		log.Fatal(err)
	}

	printReverse(values)
}

func readLength() (int, error) {
	var value int

	_, err := fmt.Scanf("%d", &value)
	if err != nil {
		return -1, err
	}

	return value, err
}

func readValues(length int) ([]int, error) {
	var err error

	values := make([]int, length)

	for i := range values {
		_, err = fmt.Scanf("%d", &values[i])
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func printReverse(values []int) {
	for i := len(values) - 1; i >= 0; i-- {
		fmt.Printf("%d", values[i])

		if i != 0 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}
