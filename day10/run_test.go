package day10

import (
	"reflect"
	"testing"
)

func Test_countDifferences(t *testing.T) {
	tests := []struct {
		name           string
		numbers        *[]int
		want1DiffCount int
		want3DiffCount int
	}{
		{"case 1", &[]int{0, 3, 4, 5}, 2, 1},
		{"case 2", &[]int{0, 3, 4, 7, 10}, 1, 3},
		{"case 3", &[]int{3, 4, 7, 10, 13}, 1, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := countDifferences(tt.numbers)
			if got != tt.want1DiffCount {
				t.Errorf("countDifferences() got = %v, want1DiffCount %v", got, tt.want1DiffCount)
			}
			if got1 != tt.want3DiffCount {
				t.Errorf("countDifferences() got1 = %v, want1DiffCount %v", got1, tt.want3DiffCount)
			}
		})
	}
}

func Test_getIgnoredNumbers(t *testing.T) {
	tests := []struct {
		name    string
		numbers *[]int
		want    *[]int
	}{
		{"nothing to ignore", &[]int{0, 1, 2}, &[]int{}},
		{"ignore gap 1", &[]int{0, 1, 3}, &[]int{2}},
		{"ignore gap 2", &[]int{0, 1, 4}, &[]int{2, 3}},
		{"ignore gaps 1 and 2", &[]int{0, 1, 4, 5, 7}, &[]int{2, 3, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := map[int]struct{}{}
			for _, x := range *tt.want {
				want[x] = struct{}{}
			}
			if got := getIgnoredNumbers(tt.numbers); !reflect.DeepEqual(got, &want) {
				t.Errorf("getIgnoredNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPathsAt(t *testing.T) {
	tests := []struct {
		name           string
		ignoredNumbers []int
		n              int
		want           int
	}{
		{"case 1", []int{2, 4, 7, 8, 11, 15, 22}, 30, 51184},
		{"case 2", []int{2, 4, 7, 8}, 10, 4},
		{"case 3", []int{2, 4, 7, 8}, 15, 96},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ignoredNumbers := map[int]struct{}{}
			for _, n := range tt.ignoredNumbers {
				ignoredNumbers[n] = struct{}{}
			}
			cache := map[int]int{}
			if got := countPathsAt(&ignoredNumbers, tt.n, &cache); got != tt.want {
				t.Errorf("countPathsAt() = %v, want %v", got, tt.want)
			}
			if _, found := cache[3]; !found {
				t.Errorf("no cache found for value 3")
			}
		})
	}
}
