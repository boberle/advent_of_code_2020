package day16

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

type SurveyedField struct {
	name   string
	range1 Range
	range2 Range
}

type Field struct {
	name    string
	value   int
	isValid bool
}

type Ticket struct {
	fields  []Field
	isValid bool
}

type Survey struct {
	fields  []SurveyedField
	tickets []Ticket
}

func Run(infile string) {

	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	survey, myTicket := parseFile(fh)
	survey.setTicketInvalidFields()
	invalidFieldSum := survey.sumInvalidValues()
	fmt.Printf("Invalid field sum (part 1): %d\n", invalidFieldSum)

	names := survey.findFieldName()
	myTicket.setFieldNames(&names)
	departureValues := myTicket.getDepartureValues()
	fmt.Printf("Departure value sum (part 2): %d\n", departureValues)

}

func parseFile(reader io.Reader) (Survey, Ticket) {
	scanner := bufio.NewScanner(reader)

	survey := Survey{fields: []SurveyedField{}, tickets: []Ticket{}}
	myTicket := Ticket{fields: []Field{}}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		pat, err := regexp.Compile("([\\w ]+): (\\d+)-(\\d+) or (\\d+)-(\\d+)")
		if err != nil {
			panic(err)
		}
		if m := pat.FindStringSubmatch(line); len(m) > 0 {
			s1, _ := strconv.Atoi(m[2])
			e1, _ := strconv.Atoi(m[3])
			s2, _ := strconv.Atoi(m[4])
			e2, _ := strconv.Atoi(m[5])
			survey.fields = append(survey.fields, SurveyedField{
				name:   m[1],
				range1: Range{s1, e1},
				range2: Range{s2, e2},
			})
		} else {
			log.Fatalf("Unable to parse line '%s'", line)
		}
	}

	scanner.Scan()
	scanner.Scan()
	fields := strings.Split(scanner.Text(), ",")
	for _, fieldString := range fields {
		fieldValue, err := strconv.Atoi(fieldString)
		if err != nil {
			log.Fatalf("unable to parse your ticket line: '%s'", scanner.Text())
		}
		myTicket.fields = append(myTicket.fields, Field{value: fieldValue})
	}

	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		ticket := Ticket{fields: []Field{}}
		fields := strings.Split(scanner.Text(), ",")
		for _, fieldString := range fields {
			fieldValue, err := strconv.Atoi(fieldString)
			if err != nil {
				log.Fatalf("unable to parse your ticket line: '%s'", scanner.Text())
			}
			ticket.fields = append(ticket.fields, Field{value: fieldValue})
		}
		survey.tickets = append(survey.tickets, ticket)
	}

	return survey, myTicket
}

func (survey *Survey) setTicketInvalidFields() {
	for i := range survey.tickets {
		survey.tickets[i].setInvalidFields(survey.fields)
	}
}

func (ticket *Ticket) setInvalidFields(surveyedFields []SurveyedField) {
	ticket.isValid = true
	for i, field := range ticket.fields {
		isSet := false
		for _, surveyedField := range surveyedFields {
			if (surveyedField.range1.start <= field.value && field.value <= surveyedField.range1.end) || (surveyedField.range2.start <= field.value && field.value <= surveyedField.range2.end) {
				ticket.fields[i].isValid = true
				isSet = true
				break
			}
		}
		if !isSet {
			ticket.isValid = false
			ticket.fields[i].isValid = false
		}
	}
}

func (survey *Survey) sumInvalidValues() int {
	rv := 0
	for _, ticket := range survey.tickets {
		rv += ticket.sumInvalidValues()
	}
	return rv
}

func (ticket *Ticket) sumInvalidValues() int {
	rv := 0
	for _, field := range ticket.fields {
		if !field.isValid {
			rv += field.value
		}
	}
	return rv
}

func (survey *Survey) findFieldName() map[string]int {
	fieldCount := len(survey.tickets[0].fields)

	matrix := Matrix{}
	fieldNames := make([]string, fieldCount)
	for i, field := range survey.fields {
		fieldNames[i] = field.name
	}
	matrix.init(fieldNames)

	for i := 0; i < fieldCount; i++ {
		found := false
		for _, surveyedField := range survey.fields {
			isFieldOK := true
			for _, ticket := range survey.tickets {
				if ticket.isValid {
					if !((surveyedField.range1.start <= ticket.fields[i].value && ticket.fields[i].value <= surveyedField.range1.end) || (surveyedField.range2.start <= ticket.fields[i].value && ticket.fields[i].value <= surveyedField.range2.end)) {
						isFieldOK = false
						break
					}
				}
			}
			if isFieldOK {
				matrix.add(i, surveyedField.name)
				found = true
			}
		}
		if !found {
			panic(fmt.Sprintf("no name found for field %d", i))
		}
	}

	names := make(map[string]int, fieldCount)
	for {
		if name, index, found := matrix.getNameByRow(); found {
			if i, found := names[name]; found {
				if i != index {
					log.Fatalf("name %s already paired to index %d, can't replace with index %d\n", name, i, index)
				}
			}
			names[name] = index
		} else {
			break
		}
	}

	return names
}

func (ticket *Ticket) setFieldNames(names *map[string]int) {
	for name, index := range *names {
		ticket.fields[index].name = name
	}
}

func (ticket *Ticket) getDepartureValues() int {
	rv := 1
	for _, field := range ticket.fields {
		if strings.HasPrefix(field.name, "departure") {
			rv *= field.value
		}
	}
	return rv
}

type Matrix struct {
	columns    []string
	values     [][]int
	rowIndices map[int]struct{}
	colIndices map[int]struct{}
}

func (m *Matrix) init(fields []string) {
	m.columns = fields
	m.values = make([][]int, len(fields))
	m.rowIndices = make(map[int]struct{}, len(fields))
	m.colIndices = make(map[int]struct{}, len(fields))
	for i := range fields {
		m.values[i] = make([]int, len(fields))
		m.rowIndices[i] = struct{}{}
		m.colIndices[i] = struct{}{}
	}
}

func (m *Matrix) add(row int, name string) {
	if nameIndex, found := m.getNameIndex(name); found {
		m.values[row][nameIndex] = 1
	} else {
		panic(fmt.Sprintf("name '%s' not found", name))
	}
}

func (m *Matrix) getNameIndex(name string) (int, bool) {
	for i, col := range m.columns {
		if col == name {
			return i, true
		}
	}
	return 0, false
}

// this is redundant with getNameByRow. Just choose one
func (m *Matrix) getNameByColumn() (string, int, bool) {
	for colIndex := range m.colIndices {
		last := -1
		for rowIndex := range m.rowIndices {
			if m.values[rowIndex][colIndex] > 0 {
				if last != -1 {
					last = -1
					break
				}
				last = rowIndex
			}
		}
		if last != -1 {
			delete(m.colIndices, colIndex)
			delete(m.rowIndices, last)
			return m.columns[colIndex], last, true
		}
	}
	return "", 0, false
}

func (m *Matrix) getNameByRow() (string, int, bool) {
	for rowIndex := range m.rowIndices {
		last := -1
		for colIndex := range m.colIndices {
			if m.values[rowIndex][colIndex] > 0 {
				if last != -1 {
					last = -1
					break
				}
				last = colIndex
			}
		}
		if last != -1 {
			delete(m.rowIndices, rowIndex)
			delete(m.colIndices, last)
			return m.columns[last], rowIndex, true
		}
	}
	return "", 0, false
}
