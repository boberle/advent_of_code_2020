package day25

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	cardPublicKey, doorPublicKey := parseFile(fh)
	cardLoopSize, _ := transform(7, cardPublicKey, 0)

	_, encryptionKey := transform(doorPublicKey, 0, cardLoopSize)
	fmt.Printf("encryption key: %d\n", encryptionKey)
}

func parseFile(reader io.Reader) (int, int) {
	scanner := bufio.NewScanner(reader)

	if !scanner.Scan() {
		panic("can't read card public key")
	}
	cardPublicKey, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	if !scanner.Scan() {
		panic("can't read door public key")
	}
	doorPublicKey, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	return cardPublicKey, doorPublicKey
}

func transform(subjectNumber, publicKey, loopSize int) (int, int) {
	size := 0
	value := 1
	for {
		size++
		value = (value * subjectNumber) % 20201227
		if value == publicKey || size == loopSize {
			return size, value
		}
	}
}
