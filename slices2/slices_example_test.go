package slices2_test

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mkch/gg/slices2"
)

func ExampleFilter() {
	isEven := func(n int) bool { return n%2 == 0 }
	evens := slices2.Filter([]int{1, 2, 3, 4, 5}, isEven)
	fmt.Println(evens)
	// Output:
	// [2 4]
}

func ExampleMap() {
	s1 := []int{1, 2, 3}
	f := func(n int) string { return "number" + strconv.Itoa(n) }
	s2 := slices2.Map(s1, f)
	fmt.Println(s2)
	// Output:
	// [number1 number2 number3]
}

func ExampleReduce() {
	s := []int{1, 2, 3}
	sumReducer := func(a, n, i int) int { return a + n }
	sum := slices2.Reduce(s, sumReducer, 0)
	fmt.Println(sum)
	// Output:
	// 6
}

func Example_mapReduce() {
	rawData := []string{"Alice:100", "Bob:90"}
	sum := slices2.Reduce(
		slices2.Map(rawData, func(s string) int {
			n, _ := strconv.Atoi(s[strings.IndexRune(s, ':')+1:])
			return n
		}),
		func(sum, n, _ int) int {
			return sum + n
		},
		0)
	fmt.Printf("Total score: %v", sum)
	// Output:
	// Total score: 190
}

func ExampleFill() {
	s := make([]string, 3)
	slices2.Fill(s, "go!")
	fmt.Println(s)
	// Output:
	// [go! go! go!]
}

func ExampleRepeat() {
	s := slices2.Repeat("go!", 3)
	fmt.Println(s)
	// Output:
	// [go! go! go!]
}
