package main

import (
	"fmt"
	"log"
)

type StringList []string

func main() {
	data, queries, err := ReadData()
	if err != nil {
		log.Fatal(err)
	}

	for _, query := range queries {
		fmt.Printf("%d\n", Occurrences(data, query))
	}
}

func ReadData() ([]string, []string, error) {
	var dataLength, queriesLength int

	_, err := fmt.Scanf("%d", &dataLength)
	if err != nil {
		return nil, nil, err
	}

	data := make([]string, dataLength)

	for i := 0; i < len(data); i++ {
		_, err := fmt.Scanf("%s", &data[i])
		if err != nil {
			return nil, nil, err
		}
	}

	_, err = fmt.Scanf("%d", &queriesLength)
	if err != nil {
		return nil, nil, err
	}

	queries := make([]string, queriesLength)

	for i := 0; i < len(queries); i++ {
		_, err := fmt.Scanf("%s", &queries[i])
		if err != nil {
			return nil, nil, err
		}
	}

	return data, queries, nil
}

func Occurrences(data []string, query string) int {
	count := 0

	for _, str := range data {
		if str == query {
			count++
		}
	}

	return count
}
