package sorted_test

import (
	"cmp"
	"fmt"

	"github.com/mkch/gg/slices2/sorted"
)

func ExampleBisectRight() {
	type Int int
	var s = []Int{1, 2, 2, 3}
	i := sorted.BisectRight(s, 2)
	fmt.Printf("The latest insert position of 2 in %v is %v", s, i)
	// Output:
	// The latest insert position of 2 in [1 2 2 3] is 3
}

func ExampleBisectRightFunc() {
	type Entry struct {
		ID   int
		Name string
	}
	cmpName := func(a, b Entry) int {
		return cmp.Compare(a.Name, b.Name)
	}
	var entries = []Entry{
		{1, "a"},
		{2, "b"},
		{3, "b"},
		{4, "c"},
	}
	var toInsert = Entry{5, "b"}
	i := sorted.BisectRightFunc(entries, toInsert, cmpName)
	fmt.Printf("The latest insert position of %v in %v is %v", toInsert, entries, i)
	// Output:
	// The latest insert position of {5 b} in [{1 a} {2 b} {3 b} {4 c}] is 3
}

func ExampleInsert() {
	type Int int
	var s = []Int{1, 2, 3}
	s = sorted.Insert(s, 2)
	fmt.Println(s)
	// Output:
	// [1 2 2 3]
}

func ExampleInsertFunc() {
	type Entry struct {
		ID   int
		Name string
	}
	cmpName := func(a, b Entry) int {
		return cmp.Compare(a.Name, b.Name)
	}
	var entries = []Entry{
		{1, "a"},
		{2, "b"},
		{3, "b"},
		{4, "c"},
	}
	var toInsert = Entry{5, "b"}
	entries = sorted.InsertFunc(entries, toInsert, cmpName)
	fmt.Println(entries)
	// Output:
	// [{1 a} {5 b} {2 b} {3 b} {4 c}]
}

func ExampleDelete() {
	type Int int
	var s = []Int{1, 2, 2, 3}
	s = sorted.Delete(s, 2)
	fmt.Println(s)
	// Output:
	// [1 3]
}

func ExampleDeleteFunc() {
	type Entry struct {
		ID   int
		Name string
	}
	cmpName := func(a, b Entry) int {
		return cmp.Compare(a.Name, b.Name)
	}
	var entries = []Entry{
		{1, "a"},
		{2, "b"},
		{3, "b"},
		{4, "c"},
	}
	var toInsert = Entry{Name: "b"}
	// All entries with Name "b" will be deleted.
	entries = sorted.DeleteFunc(entries, toInsert, cmpName)
	fmt.Println(entries)
	// Output:
	// [{1 a} {4 c}]
}

func ExampleFind() {
	var s = []int{1, 2, 2, 3}
	s2 := sorted.Find(s, 2)
	fmt.Println(s2)
	// Output:
	// [2 2]
}

func ExampleFindFunc() {
	type Entry struct {
		ID   int
		Name string
	}
	cmpName := func(a, b Entry) int {
		return cmp.Compare(a.Name, b.Name)
	}
	var entries = []Entry{
		{1, "a"},
		{3, "b"},
		{4, "b"},
		{2, "c"},
	}
	var target = Entry{Name: "b"}
	found := sorted.FindFunc(entries, target, cmpName)
	fmt.Println(found)
	// Output:
	// [{3 b} {4 b}]
}

func ExamplePrefix() {
	var s = []int{1, 1, 1, 2, 2, 3}
	for {
		prefix := sorted.Prefix(s)
		if prefix == nil {
			break
		}
		s = s[len(prefix):]
		fmt.Println(prefix)
	}
	// Output:
	// [1 1 1]
	// [2 2]
	// [3]
}

func ExamplePrefixFunc() {
	type Entry struct {
		ID   int
		Name string
	}
	cmpName := func(a, b Entry) int {
		return cmp.Compare(a.Name, b.Name)
	}
	var entries = []Entry{
		{1, "a"},
		{3, "b"},
		{4, "b"},
		{2, "c"},
	}
	for {
		prefix := sorted.PrefixFunc(entries, cmpName)
		if prefix == nil {
			break
		}
		entries = entries[len(prefix):]
		fmt.Println(prefix)
	}
	// Output:
	// [{1 a}]
	// [{3 b} {4 b}]
	// [{2 c}]
}
