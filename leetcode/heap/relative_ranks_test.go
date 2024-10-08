package heap

import (
	"reflect"
	"testing"
)

func TestFindRelativeRanks(t *testing.T) {
	testCases := []struct {
		name     string
		score    []int
		expected []string
	}{
		{name: "1", score: []int{5, 4, 3, 2, 1}, expected: []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
		{name: "2", score: []int{10, 3, 8, 9, 4}, expected: []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := FindRelativeRanks(tt.score)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
