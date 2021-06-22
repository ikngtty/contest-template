package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := NewScanner()

	n := sc.ReadInt()

	fmt.Println(n)
}

// from [my library](https://github.com/ikngtty/go-contestlib)
// package math/modular

// ExtendedEuclidean does the extended Euclidean algorithm.
func ExtendedEuclidean(a, b int) (gcd, x, y int) {
	if b == 0 {
		gcd = a
		x = 1
		y = 0
		return
	}
	q, r := EucDiv(a, b)
	gcd, s, t := ExtendedEuclidean(b, r)
	x = t
	y = s - q*t
	return
}

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

	answer := 1
	for i := 0; i < exponent; i++ {
		answer *= base
	}
	return answer
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
	if dividor == 0 {
		panic("divide by zero")
	}

	quo = divident / dividor
	rem = divident % dividor
	if rem < 0 {
		if dividor > 0 {
			quo -= 1
			rem += dividor
		} else {
			quo += 1
			rem -= dividor
		}
	}
	return
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
