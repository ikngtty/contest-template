package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Scanner struct {
	bufScanner *bufio.Scanner
}

func NewScanner() *Scanner {
	bufSc := bufio.NewScanner(os.Stdin)
	bufSc.Split(bufio.ScanWords)
	return &Scanner{bufSc}
}

func (sc *Scanner) ReadString() string {
	bufSc := sc.bufScanner
	bufSc.Scan()
	return bufSc.Text()
}

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

func main() {
	sc := NewScanner()

	n := sc.ReadInt()

	x := sc.ReadInt()
	y := sc.ReadInt()
	z := sc.ReadInt()

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = sc.ReadInt()
		b[i] = sc.ReadInt()
	}

	total := x + y + z
	for _, ai := range a {
		total += ai
	}
	for _, bi := range b {
		total += bi
	}

	fmt.Println(total)
}
