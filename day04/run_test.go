package day04

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	cases := []struct {
		input    string
		expected []passportData
	}{
		{
			input: "byr:b",
			expected: []passportData{
				{byr: "b"},
			},
		},
		{
			input: "byr:b iyr:d\neyr:f\n\nhgt:h\nhcl:j\npid:l\n\ncid:n",
			expected: []passportData{
				{byr: "b", iyr: "d", eyr: "f"},
				{hgt: "h", hcl: "j", pid: "l"},
				{cid: "n"},
			},
		},
	}

	for _, test := range cases {
		reader := strings.NewReader(test.input)
		actual := parseFile(reader)
		assert.Equal(t, test.expected, actual)
	}
}

func Test_passportData_isValid(t *testing.T) {
	cases := []struct {
		passport passportData
		expected bool
	}{
		{passportData{byr: "abc", iyr: "abc", eyr: "abc", hgt: "abc", hcl: "abc", ecl: "abc", pid: "abc", cid: "abc"}, true},
		{passportData{byr: "abc", iyr: "abc", eyr: "abc", hgt: "abc", hcl: "abc", ecl: "abc", pid: "abc", cid: ""}, true},
		{passportData{byr: "abc", iyr: "abc", eyr: "abc", hgt: "abc", hcl: "abc", ecl: "abc", pid: "", cid: ""}, false},
	}

	for _, test := range cases {
		actual := test.passport.isValid()
		assert.Equal(t, test.expected, actual)
	}
}

func Test_passportData_isByrValid(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{"", false},
		{"abc", false},
		{"2002", true},
		{"2003", false},
	}

	for _, test := range cases {
		passport := passportData{byr: test.value}
		assert.Equal(t, test.expected, passport.isByrValid(), test.value)
	}
}

func Test_passportData_isIyrValid(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{"", false},
		{"abc", false},
		{"2010", true},
		{"2023", false},
	}

	for _, test := range cases {
		passport := passportData{iyr: test.value}
		assert.Equal(t, test.expected, passport.isIyrValid(), test.value)
	}
}

func Test_passportData_isEyrValid(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{"", false},
		{"abc", false},
		{"2020", true},
		{"2033", false},
	}

	for _, test := range cases {
		passport := passportData{eyr: test.value}
		assert.Equal(t, test.expected, passport.isEyrValid(), test.value)
	}
}

func Test_passportData_isHgtValid(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{"", false},
		{"abc", false},
		{"150cm", true},
		{"193cm", true},
		{"194cm", false},
		{"150cn", false},
		{"59in", true},
		{"76in", true},
		{"77in", false},
	}

	for _, test := range cases {
		passport := passportData{hgt: test.value}
		assert.Equal(t, test.expected, passport.isHgtValid(), test.value)
	}
}

func Test_passportData_isHclValid(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{"", false},
		{"abc", false},
		{"#123abc", true},
		{"#123abz", false},
	}

	for _, test := range cases {
		passport := passportData{hcl: test.value}
		assert.Equal(t, test.expected, passport.isHclValid(), test.value)
	}
}

func Test_passportData_isEclValid(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{"", false},
		{"abc", false},
		{"amb", true},
		{"blu", true},
		{"oth", true},
	}

	for _, test := range cases {
		passport := passportData{ecl: test.value}
		assert.Equal(t, test.expected, passport.isEclValid(), test.value)
	}
}

func Test_passportData_isPidValid(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{"", false},
		{"abc", false},
		{"123", false},
		{"987654321", true},
	}

	for _, test := range cases {
		passport := passportData{pid: test.value}
		assert.Equal(t, test.expected, passport.isPidValid(), test.value)
	}
}

func Test_passportData_isStrictlyValid(t *testing.T) {
	cases := []struct {
		id       int
		passport passportData
		expected bool
	}{
		{1, passportData{eyr: "1972", cid: "100", hcl: "#18171d", ecl: "amb", hgt: "170", pid: "186cm", iyr: "2018", byr: "1926"}, false},
		{2, passportData{iyr: "2019", hcl: "#602927", eyr: "1967", hgt: "170cm", ecl: "grn", pid: "012533040", byr: "1946"}, false},
		{3, passportData{hcl: "dab227", iyr: "2012", ecl: "brn", hgt: "182cm", pid: "021572410", eyr: "2020", byr: "1992", cid: "277"}, false},
		{4, passportData{hgt: "59cm", ecl: "zzz", eyr: "2038", hcl: "74454a", iyr: "2023", pid: "3556412378", byr: "2007"}, false},
		{5, passportData{pid: "087499704", hgt: "74in", ecl: "grn", iyr: "2012", eyr: "2030", byr: "1980", hcl: "#623a2f"}, true},
		{6, passportData{eyr: "2029", ecl: "blu", cid: "129", byr: "1989", iyr: "2014", pid: "896056539", hcl: "#a97842", hgt: "165cm"}, true},
		{7, passportData{hcl: "#888785", hgt: "164cm", byr: "2001", iyr: "2015", cid: "88", pid: "545766238", ecl: "hzl", eyr: "2022"}, true},
		{8, passportData{iyr: "2010", hgt: "158cm", hcl: "#b6652a", ecl: "blu", byr: "1944", eyr: "2021", pid: "093154719"}, true},
	}

	for _, test := range cases {
		actual := test.passport.isStrictlyValid()
		assert.Equal(t, test.expected, actual, test.id)
	}
}
