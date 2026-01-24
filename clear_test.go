package gg

import "testing"

func TestClearSafe(t *testing.T) {
	var n = 1
	clearSafe(&n)
	if n != 0 {
		t.Errorf("Clear failed")
	}
	var str = "hello"
	clearSafe(&str)
	if str != "" {
		t.Errorf("Clear failed")
	}
	var s = []int{1, 2, 3}
	clearSafe(&s)
	if s != nil {
		t.Errorf("Clear failed")
	}
	var m = map[string]int{"a": 1}
	clearSafe(&m)
	if m != nil {
		t.Errorf("Clear failed")
	}
	var p = &n
	clearSafe(&p)
	if p != nil {
		t.Errorf("Clear failed")
	}

	type MyStruct struct {
		A int
		B string
		C []float64
		D map[int]int
		E chan int
	}
	var st = MyStruct{
		A: 10,
		B: "test",
		C: []float64{1.1, 2.2},
		D: map[int]int{1: 1},
		E: make(chan int),
	}
	clearSafe(&st)
	if st.A != 0 || st.B != "" || st.C != nil || st.D != nil || st.E != nil {
		t.Errorf("Clear failed")
	}

	var arr = [5]int{1, 2, 3, 4, 5}
	clearSafe(&arr)
	for i, v := range arr {
		if v != 0 {
			t.Errorf("Clear failed at index %d: got %d, want 0", i, v)
		}
	}
}

func TestClear(t *testing.T) {
	var n = 1
	Clear(&n)
	if n != 0 {
		t.Errorf("Clear2 failed")
	}
	var str = "hello"
	Clear(&str)
	if str != "" {
		t.Errorf("Clear2 failed")
	}
	var s = []int{1, 2, 3}
	Clear(&s)
	if s != nil {
		t.Errorf("Clear2 failed")
	}
	var m = map[string]int{"a": 1}
	Clear(&m)
	if m != nil {
		t.Errorf("Clear2 failed")
	}
	var p = &n
	Clear(&p)
	if p != nil {
		t.Errorf("Clear2 failed")
	}

	type MyStruct struct {
		A int
		B string
		C []float64
		D map[int]int
		E chan int
	}
	var st = MyStruct{
		A: 10,
		B: "test",
		C: []float64{1.1, 2.2},
		D: map[int]int{1: 1},
		E: make(chan int),
	}
	Clear(&st)
	if st.A != 0 || st.B != "" || st.C != nil || st.D != nil || st.E != nil {
		t.Errorf("Clear2 failed")
	}

	var arr = [5]int{1, 2, 3, 4, 5}
	Clear(&arr)
	for i, v := range arr {
		if v != 0 {
			t.Errorf("Clear2 failed at index %d: got %d, want 0", i, v)
		}
	}
}

const SmallN = 1024 * 10
const LargeN = 1024 * 100

func BenchmarkClearSafeSmall(b *testing.B) {
	var n [SmallN]int
	for b.Loop() {
		clearSafe(&n)
	}
}
func BenchmarkClearSmall(b *testing.B) {
	var n [SmallN]int
	for b.Loop() {
		Clear(&n)
	}
}

func BenchmarkClearSafeLarge(b *testing.B) {
	var n [LargeN]int
	for b.Loop() {
		clearSafe(&n)
	}
}
func BenchmarkClearLarge(b *testing.B) {
	var n [LargeN]int
	for b.Loop() {
		Clear(&n)
	}
}
