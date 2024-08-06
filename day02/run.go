package day02

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type policy struct {
	i1, i2 int
	char   rune
}

type password string
type checkPassword func(policy, password) bool

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	policies, passwords := parseFile(fh)

	counterOfficialToboggan := checkPasswords(isPasswordConformPart1, policies, passwords)
	fmt.Println("Number of passwords that satisfy policies (part 1):", counterOfficialToboggan)

	counterSledRental := checkPasswords(isPasswordConformPart2, policies, passwords)
	fmt.Println("Number of passwords that satisfy policies (part 2):", counterSledRental)

}

func parseFile(reader io.Reader) ([]policy, []password) {
	policies := []policy{}
	passwords := []password{}

	for {
		var (
			i1, i2   int
			char     rune
			password password
		)

		n, err := fmt.Fscanf(reader, "%d-%d %c: %s\n", &i1, &i2, &char, &password)
		if err != nil {
			break
		}
		if n != 4 {
			log.Fatalf("Unable to read all the data, expected 4, got %d\n", n)
		}
		policies = append(policies, policy{i1, i2, char})
		passwords = append(passwords, password)
	}

	log.Println("Number of policies read:", len(policies))
	log.Println("Number of passwords read:", len(passwords))

	if len(policies) != len(passwords) {
		log.Fatalf("Numbers of passwords and policies don't match")
	}

	if len(policies) == 0 {
		log.Fatalf("No policy or password read")
	}

	return policies, passwords
}

func isPasswordConformPart1(policy policy, password password) bool {
	count := strings.Count(string(password), string(policy.char))
	return policy.i1 <= count && count <= policy.i2
}

func isPasswordConformPart2(policy policy, password password) bool {
	runes := []rune(password)
	return (runes[policy.i1-1] == policy.char || runes[policy.i2-1] == policy.char) && runes[policy.i1-1] != runes[policy.i2-1]
}

func checkPasswords(fn checkPassword, policies []policy, passwords []password) int {
	counter := 0
	for i := 0; i < len(passwords); i++ {
		if fn(policies[i], passwords[i]) {
			counter++
		}
	}
	return counter
}
