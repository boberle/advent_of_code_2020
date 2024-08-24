package day16

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTicket_setInvalidFields(t *testing.T) {
	survey := getSampleSurvey()
	survey.setTicketInvalidFields()
	expectedSums := []int{0, 4, 55, 12}
	expectedTicketValidity := []bool{true, false, false, false}
	for i, ticket := range survey.tickets {
		got := ticket.sumInvalidValues()
		if got != expectedSums[i] {
			t.Errorf("ticket %d invalid sum mistmatch, expected: %d, got %d\n", i, expectedSums[i], got)
		}
		if ticket.isValid != expectedTicketValidity[i] {
			t.Errorf("ticket %d validity mismatch, expected %v, got %v\n", i, expectedTicketValidity[i], ticket.isValid)
		}
	}
}

func Test_ParseFile(t *testing.T) {
	input := "departure class: 1-3 or 5-7\nrow: 6-11 or 33-44\nseat: 13-40 or 45-50\n\nyour ticket:\n7,1,14\n\nnearby tickets:\n7,3,47\n40,4,50\n55,2,20\n38,6,12\n"
	expectedSurvey := getSampleSurvey()
	expectedTicket := Ticket{fields: []Field{{value: 7}, {value: 1}, {value: 14}}}

	gotSurvey, gotTicket := parseFile(strings.NewReader(input))
	assert.Equal(t, expectedSurvey, gotSurvey)
	assert.Equal(t, expectedTicket, gotTicket)
}

func getSampleSurvey() Survey {
	surveyedFields := []SurveyedField{
		{name: "departure class", range1: Range{1, 3}, range2: Range{5, 7}},
		{name: "row", range1: Range{6, 11}, range2: Range{33, 44}},
		{name: "seat", range1: Range{13, 40}, range2: Range{45, 50}},
	}
	tickets := []Ticket{
		{fields: []Field{{value: 7}, {value: 3}, {value: 47}}},
		{fields: []Field{{value: 40}, {value: 4}, {value: 50}}},
		{fields: []Field{{value: 55}, {value: 2}, {value: 20}}},
		{fields: []Field{{value: 38}, {value: 6}, {value: 12}}},
	}
	survey := Survey{
		fields:  surveyedFields,
		tickets: tickets,
	}
	return survey
}

func TestSurvey_findFieldNames(t *testing.T) {
	surveyedFields := []SurveyedField{
		{name: "class", range1: Range{0, 1}, range2: Range{4, 19}},
		{name: "row", range1: Range{0, 5}, range2: Range{8, 19}},
		{name: "seat", range1: Range{0, 13}, range2: Range{16, 19}},
	}
	tickets := []Ticket{
		{fields: []Field{{value: 3}, {value: 9}, {value: 18}}},
		{fields: []Field{{value: 15}, {value: 1}, {value: 5}}},
		{fields: []Field{{value: 5}, {value: 14}, {value: 9}}},
	}
	survey := Survey{
		fields:  surveyedFields,
		tickets: tickets,
	}
	survey.setTicketInvalidFields()
	got := survey.findFieldName()
	expected := map[string]int{"row": 0, "class": 1, "seat": 2}
	assert.Equal(t, got, expected)
}

func TestSurvey_findFieldNames_otherOrder(t *testing.T) {
	surveyedFields := []SurveyedField{
		{name: "class", range1: Range{0, 1}, range2: Range{4, 19}},
		{name: "row", range1: Range{0, 5}, range2: Range{8, 19}},
		{name: "seat", range1: Range{0, 13}, range2: Range{16, 19}},
	}
	tickets := []Ticket{
		{fields: []Field{{value: 18}, {value: 3}, {value: 9}}},
		{fields: []Field{{value: 5}, {value: 15}, {value: 1}}},
		{fields: []Field{{value: 9}, {value: 5}, {value: 14}}},
	}
	survey := Survey{
		fields:  surveyedFields,
		tickets: tickets,
	}
	survey.setTicketInvalidFields()
	got := survey.findFieldName()
	expected := map[string]int{"row": 1, "class": 2, "seat": 0}
	assert.Equal(t, got, expected)
}

func TestTicket_getDepartureValues(t *testing.T) {
	surveyedFields := []SurveyedField{
		{name: "departure_class", range1: Range{0, 1}, range2: Range{4, 19}},
		{name: "row", range1: Range{0, 5}, range2: Range{8, 19}},
		{name: "departure_seat", range1: Range{0, 13}, range2: Range{16, 19}},
	}
	tickets := []Ticket{
		{fields: []Field{{value: 3}, {value: 9}, {value: 18}}},
		{fields: []Field{{value: 15}, {value: 1}, {value: 5}}},
		{fields: []Field{{value: 5}, {value: 14}, {value: 9}}},
	}
	survey := Survey{
		fields:  surveyedFields,
		tickets: tickets,
	}
	myTicket := Ticket{fields: []Field{{value: 11}, {value: 12}, {value: 13}}}
	survey.setTicketInvalidFields()
	names := survey.findFieldName()
	myTicket.setFieldNames(&names)
	got := myTicket.getDepartureValues()
	expected := 156
	assert.Equal(t, got, expected)
}

func getMatrix() Matrix {
	return Matrix{
		columns: []string{"a", "b", "c", "d"},
		values: [][]int{
			{1, 0, 0, 0},
			{1, 1, 0, 1},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
		},
		rowIndices: map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}},
		colIndices: map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}},
	}
}

func TestMatrix_init(t *testing.T) {
	matrix := Matrix{}
	matrix.init([]string{"a", "b", "c", "d"})
	matrix.values[0][0] = 1
	matrix.values[1][1] = 1
	matrix.values[2][2] = 1
	matrix.values[3][3] = 1
	expected := [][]int{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	assert.Equal(t, expected, matrix.values)
}

func TestMatrix_getNameByRow(t *testing.T) {
	matrix := getMatrix()
	expected := map[string]int{
		"a": 0,
		"b": 2,
		"c": 3,
		"d": 1,
	}
	got := make(map[string]int, 0)
	for {
		if gotCol, gotIndex, gotFound := matrix.getNameByRow(); gotFound {
			got[gotCol] = gotIndex
		} else {
			break
		}
	}
	assert.Equal(t, expected, got)
}

func TestMatrix_getNameByColumn(t *testing.T) {
	matrix := getMatrix()
	expected := map[string]int{
		"a": 0,
		"b": 2,
		"c": 3,
		"d": 1,
	}
	got := make(map[string]int, 0)
	for {
		if gotCol, gotIndex, gotFound := matrix.getNameByColumn(); gotFound {
			got[gotCol] = gotIndex
		} else {
			break
		}
	}
	assert.Equal(t, expected, got)
}
