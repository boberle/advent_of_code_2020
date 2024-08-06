package day02

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	inputText := "1-3 a: abc\n1-2 b: def"
	reader := strings.NewReader(inputText)

	expectedPolicies := []policy{
		{1, 3, 'a'},
		{1, 2, 'b'},
	}
	expectedPasswords := []password{"abc", "def"}
	actualPolicies, actualPasswords := parseFile(reader)
	assert.Equal(t, expectedPolicies, actualPolicies)
	assert.Equal(t, expectedPasswords, actualPasswords)
}

func Test_isPasswordConformToPolicyAtOfficialToboggan(t *testing.T) {
	cases := []struct {
		policy   policy
		password password
		expected bool
	}{
		{policy{1, 3, 'a'}, "abcde", true},
		{policy{1, 3, 'b'}, "cdefg", false},
		{policy{2, 9, 'c'}, "ccccccccc", true},
	}

	for i, test := range cases {
		actual := isPasswordConformPart1(test.policy, test.password)
		if actual != test.expected {
			t.Errorf("%d: expected %v, got %v", i, test.expected, actual)
		}
	}
}

func Test_isPasswordConformToPolicyAtSledRental(t *testing.T) {
	cases := []struct {
		policy   policy
		password password
		expected bool
	}{
		{policy{1, 3, 'a'}, "abcde", true},
		{policy{1, 3, 'b'}, "cdefg", false},
		{policy{2, 9, 'c'}, "ccccccccc", false},
	}

	for i, test := range cases {
		actual := isPasswordConformPart2(test.policy, test.password)
		if actual != test.expected {
			t.Errorf("%d: expected %v, got %v", i, test.expected, actual)
		}
	}
}
