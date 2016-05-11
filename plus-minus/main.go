package main

import (
	"fmt"
	"log"
)

type List []int

func main() {
	list, err := ReadList()
	if err != nil {
		log.Fatal(err)
	}

	plus, minus, zero := list.PlusMinusZero()

	fmt.Printf("%f\n", float64(plus)/float64(len(list)))
	fmt.Printf("%f\n", float64(minus)/float64(len(list)))
	fmt.Printf("%f\n", float64(zero)/float64(len(list)))
}

func ReadList() (List, error) {
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

func (l List) PlusMinusZero() (int, int, int) {
	plus, minus, zero := 0, 0, 0

	for i := 0; i < len(l); i++ {
		switch {
		case l[i] > 0:
			plus++
		case l[i] < 0:
			minus++
		case l[i] == 0:
			zero++
		}
	}

	return plus, minus, zero
}
