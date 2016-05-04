package main

import (
	"fmt"
	"log"
)

const Q1 = 1
const Q2 = 2

type Query struct {
	Type int
	X    int
	Y    int
}

type Sequence []int

type List struct {
	Last      int
	Sequences []Sequence
	Queries   []Query
}

func main() {
	list, err := ReadList()
	if err != nil {
		log.Fatal(err)
	}

	list.Run()
}

func ReadList() (*List, error) {
	var sequenceCount, queryCount int

	_, err := fmt.Scanf("%d %d", &sequenceCount, &queryCount)
	if err != nil {
		return nil, err
	}

	list := new(List)

	list.Last = 0

	list.Sequences = make([]Sequence, sequenceCount)
	for i := range list.Sequences {
		list.Sequences[i] = make([]int, 0)
	}

	list.Queries = make([]Query, queryCount)

	for i := 0; i < queryCount; i++ {
		q := Query{}
		_, err := fmt.Scanf("%d %d %d", &q.Type, &q.X, &q.Y)
		if err != nil {
			return nil, err
		}
		list.Queries[i] = q
	}

	return list, nil
}

func (list *List) Print() {
	fmt.Printf("Queries\n-------\n")

	for _, q := range list.Queries {
		fmt.Printf("Type: %d / X: %d / Y: %d\n", q.Type, q.X, q.Y)
	}

	fmt.Printf("Sequences\n---------\n")

	for _, s := range list.Sequences {
		fmt.Printf("%v\n", s)
	}
}

func (list *List) Run() {
	for _, q := range list.Queries {
		switch q.Type {
		case Q1:
			list.Query1(q.X, q.Y)
		case Q2:
			list.Query2(q.X, q.Y)
		}
	}
}

func (list *List) Query1(x, y int) {
	index := (x ^ list.Last) % len(list.Sequences)
	list.Sequences[index] = append(list.Sequences[index], y)
}

func (list *List) Query2(x, y int) {
	index := (x ^ list.Last) % len(list.Sequences)
	list.Last = list.Sequences[index][y%len(list.Sequences[index])]
	fmt.Printf("%d\n", list.Last)
}
