package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const MaxInt = int(^uint(0) >> 1)

func main() {
	sc := NewScanner()

	n := sc.ReadInt()

	fmt.Println(n)
}

// from [my library](https://github.com/ikngtty/go-contestlib)
// package math/simple

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Min returns the minimum value of the specified values.
func Min(values ...int) int {
	if len(values) == 0 {
		panic("no values")
	}

	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum value of the specified values.
func Max(values ...int) int {
	if len(values) == 0 {
		panic("no values")
	}

	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

// Pow returns base^exponent.
func Pow(base, exponent int) int {
	if exponent < 0 {
		panic(fmt.Sprintf("invalid exponent: %d", exponent))
	}

	if exponent == 0 {
		return 1
	}

	if exponent%2 == 0 {
		half := Pow(base, exponent/2)
		return half * half
	} else {
		return base * Pow(base, exponent-1)
	}
}

// Ceil returns ceil(divident/dividor).
func Ceil(divident, dividor int) int {
	if dividor == 0 {
		panic("divide by zero")
	}

	quo := divident / dividor
	rem := divident % dividor

	if rem != 0 {
		if (divident > 0 && dividor > 0) ||
			(divident < 0 && dividor < 0) {
			return quo + 1
		}
	}
	return quo
}

// EucDiv does Euclidean divison.
func EucDiv(divident, dividor int) (quo, rem int) {
	rem = EucRem(divident, dividor)
	quo = (divident - rem) / dividor
	return
}

// EucRem returns the remainder of Euclidean divison.
func EucRem(divident, dividor int) int {
	if dividor == 0 {
		panic("divide by zero")
	}

	rem := divident % dividor
	if rem < 0 {
		rem += Abs(dividor)
	}
	return rem
}

// package arrayutil

// MakeBoolArray returns a slice of the bool array.
func MakeBoolArray(length int, initVal bool) []bool {
	a := make([]bool, length)

	if initVal {
		for i := 0; i < length; i++ {
			a[i] = initVal
		}
	}

	return a
}

// Make2DBoolArray returns a slice of the two-dimensional bool array.
func Make2DBoolArray(xLen, yLen int, initVal bool) [][]bool {
	a := make([][]bool, xLen)
	for x := 0; x < xLen; x++ {
		a[x] = make([]bool, yLen)
	}

	if initVal {
		for x := 0; x < xLen; x++ {
			for y := 0; y < yLen; y++ {
				a[x][y] = initVal
			}
		}
	}

	return a
}

// Make2DByteArray returns a slice of the two-dimensional byte array.
func Make2DByteArray(xLen, yLen int) [][]byte {
	a := make([][]byte, xLen)
	for x := 0; x < xLen; x++ {
		a[x] = make([]byte, yLen)
	}
	return a
}

// MakeIntArray returns a slice of the int array.
func MakeIntArray(length int, initVal int) []int {
	a := make([]int, length)

	if initVal != 0 {
		for i := 0; i < length; i++ {
			a[i] = initVal
		}
	}

	return a
}

// Make2DIntArray returns a slice of the two-dimensional int array.
func Make2DIntArray(xLen, yLen int, initVal int) [][]int {
	a := make([][]int, xLen)
	for x := 0; x < xLen; x++ {
		a[x] = make([]int, yLen)
	}

	if initVal != 0 {
		for x := 0; x < xLen; x++ {
			for y := 0; y < yLen; y++ {
				a[x][y] = initVal
			}
		}
	}

	return a
}

// package sortutil

// ReverseInts sorts a reversely.
func ReverseInts(a []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
}

// ReverseStrings sorts a reversely.
func ReverseStrings(a []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(a)))
}

// package io

// Scanner reads data from standard I/O.
type Scanner struct {
	bufScanner *bufio.Scanner
}

// NewScanner returns a new Scanner.
func NewScanner() *Scanner {
	bufSc := bufio.NewScanner(os.Stdin)
	bufSc.Split(bufio.ScanWords)
	bufSc.Buffer(nil, 100000000)
	return &Scanner{bufSc}
}

// ReadString reads a string value.
func (sc *Scanner) ReadString() string {
	bufSc := sc.bufScanner
	bufSc.Scan()
	return bufSc.Text()
}

// ReadInt reads a int value.
func (sc *Scanner) ReadInt() int {
	bufSc := sc.bufScanner
	bufSc.Scan()
	text := bufSc.Text()

	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return num
}

// ReadFloat reads a float value.
func (sc *Scanner) ReadFloat() float64 {
	bufSc := sc.bufScanner
	bufSc.Scan()
	text := bufSc.Text()

	num, err := strconv.ParseFloat(text, 64)
	if err != nil {
		panic(err)
	}
	return num
}
