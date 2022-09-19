package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Line []int

type Dots struct {
	count  int
	places []int
}

func NewDots(count int, place int) *Dots {

	return &Dots{count: count, places: []int{place}}
}

// Поиск минимального количества точек по заданным отрезкам, чтобы каждый отрезок содержал как минимум одну точку
func dotsOnLines() {
	var n int
	fmt.Scan(&n)
	var list []Line
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		sc.Scan()
		list = append(list, makeLine(sc.Text()))
	}
	listSort(list)

	dots := NewDots(1, list[0][1])
	for _, v := range list {
		if v[0] > dots.places[len(dots.places)-1] {
			dots.places = append(dots.places, v[1])
			dots.count++
		}
	}
	fmt.Println(dots.count)
	fmt.Println(sliceToString(dots.places))
}

func makeLine(input string) Line {
	var line Line
	nums := strings.Split(input, " ")
	for _, v := range nums {
		number, _ := strconv.Atoi(v)
		line = append(line, number)
	}
	return line
}

func listSort(list []Line) {
	sort.Slice(list, func(i, j int) bool {
		return list[i][1] < list[j][1]
	})
}

func sliceToString(list []int) string {
	var stringSlice []string
	for _, v := range list {
		stringSlice = append(stringSlice, strconv.Itoa(v))
	}

	return strings.Join(stringSlice, " ")
}
