package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Arrow int

const (
	UP Arrow = iota
	LEFT
	UPLEFT
)

type IndexAddDel struct {
	X []int
	Y []int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// X := []string{"h", "e", "l", "ee", "e", "a", "r", "sdsfwe"}

	// Y := []string{"h", "f", "ee", "s", "e", "a", "r", "fwefewa"}

	args := os.Args[1:]

	XByte, err := ioutil.ReadFile(args[0])
	check(err)

	X := strings.Split(string(XByte), "\n")

	YByte, err := ioutil.ReadFile(args[1])
	check(err)

	Y := strings.Split(string(YByte), "\n")

	xLength := len(X)
	yLength := len(Y)

	b := make([][]Arrow, xLength+1)
	for i := range b {
		b[i] = make([]Arrow, yLength+1)
	}

	c := make([][]int, xLength+1)
	for i := range c {
		c[i] = make([]int, yLength+1)
	}

	for i := 0; i < xLength+1; i++ {
		c[i][0] = 0
	}

	for i := 0; i < yLength+1; i++ {
		c[0][i] = 0
	}

	for i := 1; i < xLength+1; i++ {
		for j := 1; j < yLength+1; j++ {
			if X[i-1] == Y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				b[i][j] = UPLEFT
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = UP
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = LEFT
			}
		}
	}

	_, index := getLCS(b, X, xLength, yLength)
	// fmt.Println("-----------")
	// fmt.Println("common line")
	// fmt.Println("-----------")
	// for _, str := range ans {
	// 	fmt.Println(str)
	// }

	fmt.Println("-----------")
	fmt.Println("deleted line")
	fmt.Println("-----------")

	for i := 0; i < xLength; i++ {
		searchIn := sort.SearchInts(index.X, i)
		if searchIn >= len(index.X) || index.X[searchIn] != i {
			fmt.Println(X[i])
		}
	}

	fmt.Println("-----------")
	fmt.Println("inserted line")
	fmt.Println("-----------")

	for i := 0; i < yLength; i++ {
		searchIn := sort.SearchInts(index.Y, i)
		if searchIn >= len(index.Y) || index.Y[searchIn] != i {
			fmt.Println(Y[i])
		}
	}
}

func printLCS(b [][]Arrow, X []string, i int, j int) {
	if i == 0 || j == 0 {
		return
	}

	if b[i][j] == UPLEFT {
		printLCS(b, X, i-1, j-1)
		fmt.Printf(X[i-1])
	} else if b[i][j] == UP {
		printLCS(b, X, i-1, j)
	} else {
		printLCS(b, X, i, j-1)
	}
}

func getLCS(b [][]Arrow, X []string, i int, j int) ([]string, IndexAddDel) {
	if i == 0 || j == 0 {
		return []string{}, IndexAddDel{}
	}

	ans := []string{}
	index := IndexAddDel{}

	if b[i][j] == UPLEFT {
		ans, index = getLCS(b, X, i-1, j-1)
		ans = append(ans, X[i-1])
		index.X = append(index.X, i-1)
		index.Y = append(index.Y, j-1)
		return ans, index
	} else if b[i][j] == UP {
		return getLCS(b, X, i-1, j)
	} else {
		return getLCS(b, X, i, j-1)
	}
}
