package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	matrix, err := NewMatrix(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", matrix.AbsoluteDiagonalDifference())
}

type Matrix struct {
	Dimension int
	Data      [][]int
}

func NewMatrix(r io.Reader) (*Matrix, error) {
	m := new(Matrix)

	_, err := fmt.Fscanf(r, "%d", &m.Dimension)
	if err != nil {
		return nil, err
	}

	m.Data = make([][]int, m.Dimension)

	for i := range m.Data {
		m.Data[i] = make([]int, m.Dimension)
	}

	for y := 0; y < m.Dimension; y++ {
		for x := 0; x < m.Dimension; x++ {
			_, err := fmt.Fscanf(r, "%d", &m.Data[y][x])
			if err != nil {
				return nil, err
			}
		}
	}

	return m, nil
}

func (m *Matrix) Print() {
	for y := 0; y < m.Dimension; y++ {
		for x := 0; x < m.Dimension; x++ {
			fmt.Printf("%d", m.At(x, y))

			if x != m.Dimension-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func (m *Matrix) At(x, y int) int {
	return m.Data[y][x]
}

func (m *Matrix) Diagonal() int {
	sum := 0
	for i := 0; i < m.Dimension; i++ {
		sum += m.At(i, i)
	}
	return sum
}

func (m *Matrix) ReverseDiagonal() int {
	sum := 0
	for i := 0; i < m.Dimension; i++ {
		sum += m.At(m.Dimension-i-1, i)
	}
	return sum
}

func (m *Matrix) AbsoluteDiagonalDifference() int {
	diff := m.Diagonal() - m.ReverseDiagonal()
	if diff < 0 {
		return -diff
	} else {
		return diff
	}
}
