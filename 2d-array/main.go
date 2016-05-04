package main

import (
	"fmt"
	"log"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -(MaxInt - 1)

type Matrix struct {
	Dimension int
	Data      [][]int
}

func main() {
	matrix := NewMatrix(6)

	err := matrix.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", matrix.MaxHourglass())
}

func NewMatrix(dim int) *Matrix {
	m := new(Matrix)

	m.Dimension = dim

	m.Data = make([][]int, m.Dimension)

	for i := range m.Data {
		m.Data[i] = make([]int, m.Dimension)
	}

	return m
}

func (m *Matrix) Read() error {
	for y := 0; y < m.Dimension; y++ {
		for x := 0; x < m.Dimension; x++ {
			_, err := fmt.Scanf("%d", &m.Data[y][x])
			if err != nil {
				return err
			}
		}
	}
	return nil
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

func (m *Matrix) HourglassSum(x, y int) int {
	sum := 0
	sum += m.At(x, y)
	sum += m.At(x+1, y)
	sum += m.At(x+2, y)
	sum += m.At(x+1, y+1)
	sum += m.At(x, y+2)
	sum += m.At(x+1, y+2)
	sum += m.At(x+2, y+2)
	return sum
}

func (m *Matrix) MaxHourglass() int {
	max := MinInt

	for y := 0; y < m.Dimension-2; y++ {
		for x := 0; x < m.Dimension-2; x++ {
			sum := m.HourglassSum(x, y)
			if sum > max {
				max = sum
			}
		}
	}

	return max
}
