package sorted

import (
	"cmp"
	"reflect"
	"testing"
)

func TestBisectRight(t *testing.T) {
	tests := []struct {
		name        string // Test case name
		slice       []int  // Input sorted slice
		target      int    // Target value to find
		expectedIdx int    // Expected index
	}{
		{
			name:        "EmptySlice", // Empty slice
			slice:       []int{},
			target:      5,
			expectedIdx: 0, // In an empty slice, the insertion point for any value is index 0
		},
		{
			name:        "TargetSmallerThanAll", // Target value is smaller than all elements
			slice:       []int{10, 20, 30, 40, 50},
			target:      5,
			expectedIdx: 0, // Should be inserted at the beginning
		},
		{
			name:        "TargetLargerThanAll", // Target value is larger than all elements
			slice:       []int{10, 20, 30, 40, 50},
			target:      55,
			expectedIdx: 5, // Should be inserted at the end (slice length)
		},
		{
			name:        "TargetInMiddleExact", // Target value exists in the middle (non-duplicate)
			slice:       []int{10, 20, 30, 40, 50},
			target:      30,
			expectedIdx: 3, // Should be inserted after 30, at the position of 40
		},
		{
			name:        "TargetInMiddleBetween", // Target value does not exist, between two elements in the middle
			slice:       []int{10, 20, 30, 40, 50},
			target:      35,
			expectedIdx: 3, // Should be inserted after 30, before 40, at the position of 40
		},
		{
			name:        "TargetAtStartExact", // Target value is the first element (non-duplicate)
			slice:       []int{10, 20, 30, 40, 50},
			target:      10,
			expectedIdx: 1, // Should be inserted after 10, at the position of 20
		},
		{
			name:        "TargetAtEndExact", // Target value is the last element (non-duplicate)
			slice:       []int{10, 20, 30, 40, 50},
			target:      50,
			expectedIdx: 5, // Should be inserted after 50, at the end of the slice
		},
		{
			name:        "DuplicatesAtStart", // Target value appears repeatedly at the beginning
			slice:       []int{10, 10, 10, 20, 30},
			target:      10,
			expectedIdx: 3, // Should be inserted after all 10s, at the position of the first 20
		},
		{
			name:        "DuplicatesInMiddle", // Target value appears repeatedly in the middle
			slice:       []int{10, 20, 20, 20, 30},
			target:      20,
			expectedIdx: 4, // Should be inserted after all 20s, at the position of the first 30
		},
		{
			name:        "DuplicatesAtEnd", // Target value appears repeatedly at the end
			slice:       []int{10, 20, 30, 30, 30},
			target:      30,
			expectedIdx: 5, // Should be inserted after all 30s, at the end of the slice
		},
		{
			name:        "AllElementsAreTarget", // All elements in the slice are equal to the target value
			slice:       []int{10, 10, 10, 10},
			target:      10,
			expectedIdx: 4, // Should be inserted after all 10s, at the end of the slice
		},
		{
			name:        "TargetBetweenDuplicates", // Target value does not exist, between duplicate elements
			slice:       []int{10, 10, 20, 20, 30, 30},
			target:      15,
			expectedIdx: 2, // Should be inserted between 10 and 20, at the position of the first 20
		},
		{
			name:        "SingleElement_Less", // Single element slice, target value is less than it
			slice:       []int{10},
			target:      5,
			expectedIdx: 0, // Should be inserted at the beginning
		},
		{
			name:        "SingleElement_Equal", // Single element slice, target value is equal to it
			target:      10,
			slice:       []int{10},
			expectedIdx: 1, // Should be inserted at the end
		},
		{
			name:        "SingleElement_Greater", // Single element slice, target value is greater than it
			slice:       []int{10},
			target:      15,
			expectedIdx: 1, // Should be inserted at the end
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotIdx := BisectRight(tc.slice, tc.target)
			if gotIdx != tc.expectedIdx {
				t.Errorf("BinarySearchRight(%v, %v) = %d, expected %d", tc.slice, tc.target, gotIdx, tc.expectedIdx)
			}
		})
	}
}

func TestBisectRightFunc(t *testing.T) {
	tests := []struct {
		name        string             // Test case name
		slice       []int              // Input sorted slice
		target      int                // Target value to find
		cmp         func(int, int) int // Comparison function
		expectedIdx int                // Expected index
	}{
		{
			name:        "Func_EmptySlice",
			slice:       []int{},
			target:      5,
			cmp:         cmp.Compare[int], // Use standard comparison
			expectedIdx: 0,
		},
		{
			name:        "Func_TargetSmallerThanAll",
			slice:       []int{10, 20, 30, 40, 50},
			target:      5,
			cmp:         cmp.Compare[int],
			expectedIdx: 0,
		},
		{
			name:        "Func_TargetLargerThanAll",
			slice:       []int{10, 20, 30, 40, 50},
			target:      55,
			cmp:         cmp.Compare[int],
			expectedIdx: 5,
		},
		{
			name:        "Func_TargetInMiddleExact",
			slice:       []int{10, 20, 30, 40, 50},
			target:      30,
			cmp:         cmp.Compare[int],
			expectedIdx: 3,
		},
		{
			name:        "Func_TargetInMiddleBetween",
			slice:       []int{10, 20, 30, 40, 50},
			target:      35,
			cmp:         cmp.Compare[int],
			expectedIdx: 3,
		},
		{
			name:        "Func_Duplicates",
			slice:       []int{10, 20, 20, 20, 30},
			target:      20,
			cmp:         cmp.Compare[int],
			expectedIdx: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotIdx := BisectRightFunc(tc.slice, tc.target, tc.cmp)
			if gotIdx != tc.expectedIdx {
				t.Errorf("BinarySearchRightFunc(%v, %v) = %d, expected %d", tc.slice, tc.target, gotIdx, tc.expectedIdx)
			}
		})
	}
}

// TestInsert tests the Insert function.
func TestInsert(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		e    int
		want []int
	}{
		{"empty slice, insert 5", []int{}, 5, []int{5}},
		{"insert 5 at beginning", []int{10, 20, 30}, 5, []int{5, 10, 20, 30}},
		{"insert 35 at end", []int{10, 20, 30}, 35, []int{10, 20, 30, 35}},
		{"insert 25 in middle", []int{10, 20, 30}, 25, []int{10, 20, 25, 30}},
		{"insert 20 (duplicate) in middle", []int{10, 20, 30}, 20, []int{10, 20, 20, 30}},
		{"insert into single element slice", []int{10}, 5, []int{5, 10}},
		{"insert into single element slice (after)", []int{10}, 15, []int{10, 15}},
		{"insert duplicate into single element slice", []int{10}, 10, []int{10, 10}},
		{"insert into slice with duplicates", []int{10, 20, 20, 30}, 20, []int{10, 20, 20, 20, 30}},
		{"insert into slice with duplicates (before)", []int{10, 20, 20, 30}, 15, []int{10, 15, 20, 20, 30}},
		{"insert into slice with duplicates (after)", []int{10, 20, 20, 30}, 25, []int{10, 20, 20, 25, 30}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Insert(tt.s, tt.e)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert(%v, %v) = %v, want %v", tt.s, tt.e, got, tt.want)
			}
		})
	}
}

func TestInsertFunc(t *testing.T) {
	// Custom comparison function for reverse order (for testing purposes)
	cmpReverseInt := func(a, b int) int {
		if a < b {
			return 1 // a comes after b in reverse order
		} else if a > b {
			return -1 // a comes before b in reverse order
		}
		return 0 // equal
	}

	// Custom struct and comparison for testing InsertFunc with non-ordered types
	type Item struct {
		ID   int
		Name string
	}
	cmpItemID := func(a, b Item) int {
		if a.ID < b.ID {
			return -1
		} else if a.ID > b.ID {
			return 1
		}
		return 0
	}

	tests := []struct {
		name string
		s    []int
		e    int
		cmp  func(a, b int) int
		want []int
	}{
		{"reverse: empty, insert 5", []int{}, 5, cmpReverseInt, []int{5}},
		{"reverse: insert 5 (largest in reverse) at beginning", []int{30, 20, 10}, 5, cmpReverseInt, []int{30, 20, 10, 5}}, // 5 is largest in reverse order
		{"reverse: insert 35 (smallest in reverse) at end", []int{30, 20, 10}, 35, cmpReverseInt, []int{35, 30, 20, 10}},   // 35 is smallest in reverse order
		{"reverse: insert 25 in middle", []int{30, 20, 10}, 25, cmpReverseInt, []int{30, 25, 20, 10}},
		{"reverse: insert 20 (duplicate)", []int{30, 20, 10}, 20, cmpReverseInt, []int{30, 20, 20, 10}},
		{"reverse: insert into single element", []int{10}, 5, cmpReverseInt, []int{10, 5}},
		{"reverse: insert into single element (after)", []int{10}, 15, cmpReverseInt, []int{15, 10}},
	}

	testsStruct := []struct {
		name string
		s    []Item
		e    Item
		cmp  func(a, b Item) int
		want []Item
	}{
		{"struct: empty, insert {2}", []Item{}, Item{ID: 2, Name: "B"}, cmpItemID, []Item{{ID: 2, Name: "B"}}},
		{"struct: insert {1} at beginning", []Item{{ID: 2}, {ID: 4}, {ID: 6}}, Item{ID: 1, Name: "A"}, cmpItemID, []Item{{ID: 1, Name: "A"}, {ID: 2}, {ID: 4}, {ID: 6}}},
		{"struct: insert {7} at end", []Item{{ID: 2}, {ID: 4}, {ID: 6}}, Item{ID: 7, Name: "G"}, cmpItemID, []Item{{ID: 2}, {ID: 4}, {ID: 6}, {ID: 7, Name: "G"}}},
		{"struct: insert {3} in middle", []Item{{ID: 2}, {ID: 4}, {ID: 6}}, Item{ID: 3, Name: "C"}, cmpItemID, []Item{{ID: 2}, {ID: 3, Name: "C"}, {ID: 4}, {ID: 6}}},
		{"struct: insert {4} (duplicate ID)", []Item{{ID: 2}, {ID: 4}, {ID: 6}}, Item{ID: 4, Name: "D"}, cmpItemID, []Item{{ID: 2}, {ID: 4, Name: "D"}, {ID: 4}, {ID: 6}}}, // Note: The second {ID: 4} will be inserted before the first one by BinarySearchFunc behavior.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertFunc(tt.s, tt.e, tt.cmp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertFunc(%v, %v) = %v, want %v", tt.s, tt.e, got, tt.want)
			}
		})
	}

	for _, tt := range testsStruct {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertFunc(tt.s, tt.e, tt.cmp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertFunc(%v, %v) = %v, want %v", tt.s, tt.e, got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		e    int
		want []int
	}{
		{"empty slice, delete 5", []int{}, 5, []int{}},
		{"element not present", []int{10, 20, 30}, 5, []int{10, 20, 30}},
		{"delete 20 (once)", []int{10, 20, 30}, 20, []int{10, 30}},
		{"delete 10 (once, start)", []int{10, 20, 30}, 10, []int{20, 30}},
		{"delete 30 (once, end)", []int{10, 20, 30}, 30, []int{10, 20}},
		{"delete 20 (multiple)", []int{10, 20, 20, 20, 30}, 20, []int{10, 30}},
		{"delete 10 (multiple, start)", []int{10, 10, 20, 30}, 10, []int{20, 30}},
		{"delete 30 (multiple, end)", []int{10, 20, 30, 30}, 30, []int{10, 20}},
		{"delete all elements", []int{10, 10, 10}, 10, []int{}},
		{"delete the only element", []int{10}, 10, []int{}},
		{"delete element from single element (not present)", []int{10}, 5, []int{10}},
		{"delete element not matching duplicates", []int{10, 10, 20, 20}, 15, []int{10, 10, 20, 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Delete(tt.s, tt.e)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete(%v, %v) = %v, want %v", tt.s, tt.e, got, tt.want)
			}
		})
	}
}

func TestDeleteFunc(t *testing.T) {
	// Custom comparison function for reverse order
	cmpReverseInt := func(a, b int) int {
		if a < b {
			return 1
		} else if a > b {
			return -1
		}
		return 0
	}

	tests := []struct {
		name string
		s    []int
		e    int
		cmp  func(a, b int) int
		want []int
	}{
		{"reverse: empty slice, delete 5", []int{}, 5, cmpReverseInt, []int{}},
		{"reverse: element not present", []int{30, 20, 10}, 5, cmpReverseInt, []int{30, 20, 10}},
		{"reverse: delete 20 (once)", []int{30, 20, 10}, 20, cmpReverseInt, []int{30, 10}},
		{"reverse: delete 30 (once, start in reverse)", []int{30, 20, 10}, 30, cmpReverseInt, []int{20, 10}},
		{"reverse: delete 10 (once, end in reverse)", []int{30, 20, 10}, 10, cmpReverseInt, []int{30, 20}},
		{"reverse: delete 20 (multiple)", []int{30, 20, 20, 20, 10}, 20, cmpReverseInt, []int{30, 10}},
		{"reverse: delete 30 (multiple, start in reverse)", []int{30, 30, 20, 10}, 30, cmpReverseInt, []int{20, 10}},
		{"reverse: delete 10 (multiple, end in reverse)", []int{30, 20, 10, 10}, 10, cmpReverseInt, []int{30, 20}},
		{"reverse: delete all elements", []int{10, 10, 10}, 10, cmpReverseInt, []int{}},
		{"reverse: delete the only element", []int{10}, 10, cmpReverseInt, []int{}},
		{"reverse: delete element from single element (not present)", []int{10}, 5, cmpReverseInt, []int{10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteFunc(tt.s, tt.e, tt.cmp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFunc(%v, %v) = %v, want %v", tt.s, tt.e, got, tt.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		e    int
		want []int
	}{
		{"empty slice, find 5", []int{}, 5, nil},
		{"element not present", []int{10, 20, 30}, 5, nil},
		{"find 20 (once)", []int{10, 20, 30}, 20, []int{20}},
		{"find 10 (once, start)", []int{10, 20, 30}, 10, []int{10}},
		{"find 30 (once, end)", []int{10, 20, 30}, 30, []int{30}},
		{"find 20 (multiple)", []int{10, 20, 20, 20, 30}, 20, []int{20, 20, 20}},
		{"find 10 (multiple, start)", []int{10, 10, 20, 30}, 10, []int{10, 10}},
		{"find 30 (multiple, end)", []int{10, 20, 30, 30}, 30, []int{30, 30}},
		{"find from slice with only identical elements", []int{10, 10, 10}, 10, []int{10, 10, 10}},
		{"find element not matching duplicates", []int{10, 10, 20, 20}, 15, nil},
		{"find the only element", []int{10}, 10, []int{10}},
		{"find non-existent in single element", []int{10}, 5, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Find(tt.s, tt.e)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find(%v, %v) = %v, want %v", tt.s, tt.e, got, tt.want)
			}
		})
	}
}

func TestFindFunc(t *testing.T) {
	// Custom comparison function for reverse order
	cmpReverseInt := func(a, b int) int {
		if a < b {
			return 1
		} else if a > b {
			return -1
		}
		return 0
	}

	tests := []struct {
		name string
		s    []int
		e    int
		cmp  func(a, b int) int
		want []int
	}{
		{"reverse: empty slice, find 5", []int{}, 5, cmpReverseInt, nil},
		{"reverse: element not present", []int{30, 20, 10}, 5, cmpReverseInt, nil},
		{"reverse: find 20 (once)", []int{30, 20, 10}, 20, cmpReverseInt, []int{20}},
		{"reverse: find 30 (once, start in reverse)", []int{30, 20, 10}, 30, cmpReverseInt, []int{30}},
		{"reverse: find 10 (once, end in reverse)", []int{30, 20, 10}, 10, cmpReverseInt, []int{10}},
		{"reverse: find 20 (multiple)", []int{30, 20, 20, 20, 10}, 20, cmpReverseInt, []int{20, 20, 20}},
		{"reverse: find 30 (multiple, start in reverse)", []int{30, 30, 20, 10}, 30, cmpReverseInt, []int{30, 30}},
		{"reverse: find 10 (multiple, end in reverse)", []int{30, 20, 10, 10}, 10, cmpReverseInt, []int{10, 10}},
		{"reverse: find from slice with only identical elements", []int{10, 10, 10}, 10, cmpReverseInt, []int{10, 10, 10}},
		{"reverse: find element not matching duplicates", []int{30, 30, 20, 20}, 15, cmpReverseInt, nil},
		{"reverse: find the only element", []int{10}, 10, cmpReverseInt, []int{10}},
		{"reverse: find non-existent in single element", []int{10}, 5, cmpReverseInt, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindFunc(tt.s, tt.e, tt.cmp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindFunc(%v, %v) = %v, want %v", tt.s, tt.e, got, tt.want)
			}
		})
	}
}

func TestPrefix(t *testing.T) {
	tests := []struct {
		name       string
		s          []int
		wantPrefix []int
	}{
		{"empty slice", []int{}, nil},
		{"single element", []int{10}, []int{10}},
		{"all identical", []int{10, 10, 10}, []int{10, 10, 10}},
		{"distinct elements", []int{10, 20, 30}, []int{10}},
		{"prefix run", []int{10, 10, 20, 30}, []int{10, 10}},
		{"prefix run at end", []int{10, 20, 30, 30}, []int{10}},    // First run is just 10
		{"prefix run in middle", []int{10, 20, 20, 30}, []int{10}}, // First run is just 10
		{"slice starting with zero", []int{0, 0, 1, 2}, []int{0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefix := Prefix(tt.s)
			if !reflect.DeepEqual(gotPrefix, tt.wantPrefix) {
				t.Errorf("Prefix(%v) = %v, want %v", tt.s, gotPrefix, tt.wantPrefix)
			}
		})
	}
}

func TestPrefixFunc(t *testing.T) {
	// Custom comparison function for reverse order
	cmpReverseInt := func(a, b int) int {
		if a < b {
			return 1
		} else if a > b {
			return -1
		}
		return 0
	}

	// Custom struct and comparison for testing SplitPrefixRunFunc
	type Item struct {
		ID   int
		Name string
	}
	// This comparison considers items "equal" if their ID is the same.
	cmpItemIDEqual := func(a, b Item) int {
		if a.ID < b.ID {
			return -1 // a is "less than" b
		} else if a.ID > b.ID {
			return 1 // a is "greater than" b
		}
		return 0 // a is "equal" to b based on ID
	}

	tests := []struct {
		name       string
		s          []int // Sorted according to cmpFunc
		cmp        func(a, b int) int
		wantPrefix []int
	}{
		{"reverse: empty slice", []int{}, cmpReverseInt, nil},
		{"reverse: single element", []int{10}, cmpReverseInt, []int{10}},
		{"reverse: all identical", []int{10, 10, 10}, cmpReverseInt, []int{10, 10, 10}},
		{"reverse: distinct elements", []int{30, 20, 10}, cmpReverseInt, []int{30}},
		{"reverse: prefix run", []int{30, 30, 20, 10}, cmpReverseInt, []int{30, 30}},
		{"reverse: prefix run at end", []int{30, 20, 10, 10}, cmpReverseInt, []int{30}}, // First run is 30
	}

	testsStruct := []struct {
		name       string
		s          []Item // Sorted by ID
		cmp        func(a, b Item) int
		wantPrefix []Item
	}{
		{"struct: empty slice", []Item{}, cmpItemIDEqual, nil},
		{"struct: single element", []Item{{ID: 1, Name: "A"}}, cmpItemIDEqual, []Item{{ID: 1, Name: "A"}}},
		{"struct: all identical ID", []Item{{ID: 1, Name: "A"}, {ID: 1, Name: "B"}, {ID: 1, Name: "C"}}, cmpItemIDEqual, []Item{{ID: 1, Name: "A"}, {ID: 1, Name: "B"}, {ID: 1, Name: "C"}}},
		{"struct: distinct ID", []Item{{ID: 1}, {ID: 2}, {ID: 3}}, cmpItemIDEqual, []Item{{ID: 1}}},
		{"struct: prefix run by ID", []Item{{ID: 1, Name: "A"}, {ID: 1, Name: "B"}, {ID: 2}, {ID: 3}}, cmpItemIDEqual, []Item{{ID: 1, Name: "A"}, {ID: 1, Name: "B"}}},
		{"struct: prefix run at end by ID", []Item{{ID: 1}, {ID: 2, Name: "C"}, {ID: 2, Name: "D"}}, cmpItemIDEqual, []Item{{ID: 1}}}, // First run is {ID: 1}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefix := PrefixFunc(tt.s, tt.cmp)
			if !reflect.DeepEqual(gotPrefix, tt.wantPrefix) {
				t.Errorf("PrefixFunc(%v) = %v, want %v", tt.s, gotPrefix, tt.wantPrefix)
			}
		})
	}

	for _, tt := range testsStruct {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefix := PrefixFunc(tt.s, tt.cmp)
			if !reflect.DeepEqual(gotPrefix, tt.wantPrefix) {
				t.Errorf("PrefixFunc(%v) = %v, want %v", tt.s, gotPrefix, tt.wantPrefix)
			}
		})
	}
}
