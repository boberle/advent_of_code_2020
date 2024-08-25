package day17

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Cube
	}{
		{"test 1", ".#.\n..#\n###", []Cube{{1, 0, 0}, {2, 1, 0}, {0, 2, 0}, {1, 2, 0}, {2, 2, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			got := parseFile(reader)
			want := map[Cube]struct{}{}
			for _, cube := range tt.want {
				want[cube] = struct{}{}
			}
			assert.Equal(t, want, got)
		})
	}
}
