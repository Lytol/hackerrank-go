package main

import (
	"fmt"
	"log"
)

func main() {
	data, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", Sum(data))
}

func Sum(list []int64) int64 {
	var sum int64 = 0

	for _, v := range list {
		sum += v
	}

	return sum
}

func readData() ([]int64, error) {
	var length int64

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return nil, err
	}

	data := make([]int64, length)

	for i := range data {
		_, err := fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}
