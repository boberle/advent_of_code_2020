package day08

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Op struct {
	op      string
	value   int
	visited bool
}

func Run(infile string) {

	fh, err := os.Open(infile)
	if err != nil {
		log.Fatalln(err)
	}
	defer fh.Close()

	program := parseFile(fh)

	accValue, _ := executeProgram(&program)
	fmt.Printf("Part 1: the accumulator value is %d\n", accValue)

	accValue = fixProgram(&program)
	fmt.Printf("Part 2: the accumulator value is %d\n", accValue)

}

func parseFile(reader io.Reader) []Op {
	scanner := bufio.NewScanner(reader)
	ops := []Op{}

	for scanner.Scan() {
		op := Op{}
		fmt.Sscanf(scanner.Text(), "%s %d", &op.op, &op.value)
		ops = append(ops, op)
	}

	return ops
}

func executeProgram(ops *[]Op) (int, bool) {
	var acc, pos int

	length := len(*ops)
	for i := 0; i < length; i++ {
		(*ops)[i].visited = false
	}

	for {
		op := &(*ops)[pos]
		op.visited = true
		switch op.op {
		case "nop":
			pos++
			break
		case "acc":
			acc += op.value
			pos++
			break
		case "jmp":
			pos += op.value
			break
		default:
			log.Fatalf("Unknow op: %s\n", op.op)
		}
		if pos == len(*ops) {
			return acc, true
		}
		if (*ops)[pos].visited {
			return acc, false
		}
	}
}

func fixProgram(ops *[]Op) int {
	length := len(*ops)
	for i := 0; i < length; i++ {
		op := &(*ops)[i]
		if op.op == "acc" {
			continue
		}
		savedOp := op.op
		if op.op == "jmp" {
			op.op = "nop"
		} else if op.op == "nop" {
			op.op = "jmp"
		}
		acc, finished := executeProgram(ops)
		op.op = savedOp
		if finished {
			return acc
		}
	}
	log.Fatalln("couldn't fix the program")
	return 0
}
