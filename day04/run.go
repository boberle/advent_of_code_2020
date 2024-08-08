package day04

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type passportData struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *passportData) isValid() bool {
	return len(p.byr) > 0 && len(p.iyr) > 0 && len(p.eyr) > 0 && len(p.hgt) > 0 && len(p.hcl) > 0 && len(p.ecl) > 0 && len(p.pid) > 0
}

func (p *passportData) isStrictlyValid() bool {
	return p.isByrValid() && p.isIyrValid() && p.isEyrValid() && p.isHgtValid() && p.isHclValid() && p.isEclValid() && p.isPidValid()
}

func (p *passportData) isByrValid() bool {
	n, err := strconv.Atoi(p.byr)
	return err == nil && 1920 <= n && n <= 2002
}

func (p *passportData) isIyrValid() bool {
	n, err := strconv.Atoi(p.iyr)
	return err == nil && 2010 <= n && n <= 2020
}

func (p *passportData) isEyrValid() bool {
	n, err := strconv.Atoi(p.eyr)
	return err == nil && 2020 <= n && n <= 2030
}

func (p *passportData) isHgtValid() bool {
	var n int
	_, err := fmt.Sscanf(p.hgt, "%din", &n)
	if err != nil {
		_, err = fmt.Sscanf(p.hgt, "%dcm", &n)
		if err != nil {
			return false
		}
		return 150 <= n && n <= 193
	}
	return 59 <= n && n <= 76
}

func (p *passportData) isHclValid() bool {
	if len(p.hcl) != 7 {
		return false
	}
	for i, c := range p.hcl {
		if i == 0 {
			if c != '#' {
				return false
			}
		} else if !strings.ContainsRune("0123456789abcdef", c) {
			return false
		}
	}
	return true
}

func (p *passportData) isEclValid() bool {
	t := strings.ToLower(p.ecl)
	return t == "amb" || t == "blu" || t == "brn" || t == "gry" || t == "grn" || t == "hzl" || t == "oth"
}

func (p *passportData) isPidValid() bool {
	if len(p.pid) != 9 {
		return false
	}
	for _, c := range p.pid {
		if !strings.ContainsRune("0123456789", c) {
			return false
		}
	}
	return true
}

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	passports := parseFile(fh)

	validPassportCount := countValidPassports(passports, false)
	fmt.Printf("Part 1: number of valid passports: %d\n", validPassportCount)

	strictlyValidPassportCount := countValidPassports(passports, true)
	fmt.Printf("Part 2: number of strictly valid passports: %d\n", strictlyValidPassportCount)
}

func parseFile(reader io.Reader) []passportData {
	passports := make([]passportData, 1)
	var count int

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			passports = append(passports, passportData{})
			count++
			continue
		}
		for _, info := range strings.Split(line, " ") {
			pair := strings.Split(info, ":")
			switch pair[0] {
			case "byr":
				passports[count].byr = pair[1]
			case "iyr":
				passports[count].iyr = pair[1]
			case "eyr":
				passports[count].eyr = pair[1]
			case "hgt":
				passports[count].hgt = pair[1]
			case "hcl":
				passports[count].hcl = pair[1]
			case "ecl":
				passports[count].ecl = pair[1]
			case "pid":
				passports[count].pid = pair[1]
			case "cid":
				passports[count].cid = pair[1]
			default:
				log.Fatalf("unkown key data: %s", pair[0])
			}
		}

	}
	return passports
}

func countValidPassports(passports []passportData, strict bool) int {
	var counter int
	for _, passport := range passports {
		if (strict && passport.isStrictlyValid()) || (!strict && passport.isValid()) {
			counter++
		}
	}
	return counter
}
