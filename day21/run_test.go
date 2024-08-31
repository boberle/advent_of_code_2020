package day21

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  AllergenGroups
	}{
		{"one item", "abc (contains def)", AllergenGroups{"def": []*Group{{"abc"}}}},
		{"two items", "abc def (contains ghi, jkl)", AllergenGroups{
			"ghi": []*Group{{"abc", "def"}},
			"jkl": []*Group{{"abc", "def"}},
		}},
		{"multiline", "abc def (contains ghi, jkl)\nABC (contains DEF)", AllergenGroups{
			"ghi": []*Group{{"abc", "def"}},
			"jkl": []*Group{{"abc", "def"}},
			"DEF": []*Group{{"ABC"}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			_, got := parseFile(reader)
			assert.Equal(t, tt.want, got)
		})
	}
}
