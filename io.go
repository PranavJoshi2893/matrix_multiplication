package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func LoadFlatMatrix(filename string, n int) ([][]int, error) {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const maxBuffer = 1024 * 1024
	buf := make([]byte, maxBuffer)
	scanner.Buffer(buf, maxBuffer)
	scanner.Split(bufio.ScanWords)

	row, col := 0, 0
	for scanner.Scan() {
		if row >= n {
			break
		}
		v, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return nil, err
		}
		matrix[row][col] = v
		col++
		if col == n {
			col = 0
			row++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if row != n {
		return nil, fmt.Errorf("expected %d numbers, got %d", n*n, row*n+col)
	}

	return matrix, nil
}

func WriteFlatMatrix(filename string, n int, rng *rand.Rand) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	first := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !first {
				w.WriteByte(' ')
			}
			first = false
			fmt.Fprint(w, rng.Intn(RAND_MAX))
		}
	}
	return w.Flush()
}

func WriteFlatResult(filename string, m [][]int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	first := true
	for i := range m {
		for j := range m[i] {
			if !first {
				w.WriteByte(' ')
			}
			first = false
			fmt.Fprint(w, m[i][j])
		}
	}
	return w.Flush()
}
