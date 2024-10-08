package playground

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{name: "1", arr: []int{64, 25, 12, 22, 11}, expected: []int{11, 12, 22, 25, 64}},
		{name: "2", arr: []int{2, 3, 4, 2}, expected: []int{2, 2, 3, 4}},
		{name: "3", arr: []int{1}, expected: []int{1}},
		{name: "4", arr: []int{}, expected: []int{}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectionSort(tt.arr, len(tt.arr))
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{name: "1", arr: []int{64, 25, 12, 22, 11}, expected: []int{11, 12, 22, 25, 64}},
		{name: "2", arr: []int{2, 3, 4, 2}, expected: []int{2, 2, 3, 4}},
		{name: "3", arr: []int{1}, expected: []int{1}},
		{name: "4", arr: []int{}, expected: []int{}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := BubbleSort(tt.arr, len(tt.arr))
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{name: "1", arr: []int{23, 1, 10, 5, 2}, expected: []int{1, 2, 5, 10, 23}},
		{name: "2", arr: []int{7, 7, 7, 1}, expected: []int{1, 7, 7, 7}},
		{name: "3", arr: []int{1}, expected: []int{1}},
		{name: "4", arr: []int{}, expected: []int{}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertionSort(tt.arr, len(tt.arr))
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{name: "1", arr: []int{23, 1, 10, 5, 2}, expected: []int{1, 2, 5, 10, 23}},
		{name: "2", arr: []int{7, 7, 7, 1}, expected: []int{1, 7, 7, 7}},
		{name: "3", arr: []int{1}, expected: []int{1}},
		{name: "4", arr: []int{}, expected: []int{}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSort(tt.arr)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{name: "1", arr: []int{23, 1, 10, 5, 2}, expected: []int{1, 2, 5, 10, 23}},
		{name: "2", arr: []int{7, 7, 7, 1}, expected: []int{1, 7, 7, 7}},
		{name: "3", arr: []int{1}, expected: []int{1}},
		{name: "4", arr: []int{}, expected: []int{}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := QuickSort(tt.arr)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
