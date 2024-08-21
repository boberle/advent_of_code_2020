package day14

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

const BIT_LEN = 36

type Processor struct {
	currentMask    *Mask
	memory         map[int]int
	decoderVersion int
}

type Instruction interface {
	process(processor *Processor)
}

type MaskBit int

const (
	MaskBitOff    MaskBit = 0
	MaskBitOn     MaskBit = 1
	MaskBitIgnore MaskBit = 2
)

type Mask struct {
	value [BIT_LEN]MaskBit
}

func (mask Mask) process(processor *Processor) {
	processor.currentMask = &mask
}

type MemorySet struct {
	location int
	value    int
}

func (memorySet MemorySet) process(processor *Processor) {
	if processor.decoderVersion == 1 {
		var value int
		if processor.currentMask == nil {
			value = memorySet.value
		} else {
			value = applyMask(processor.currentMask, memorySet.value)
		}
		processor.memory[memorySet.location] = value
	} else if processor.decoderVersion == 2 {
		addresses := applyMaskWithFloatingBits(processor.currentMask, memorySet.location)
		for _, address := range addresses {
			processor.memory[address] = memorySet.value
		}
	} else {
		log.Fatalf("Invalid decoder version %d\n", processor.decoderVersion)
	}
}

func applyMask(mask *Mask, value int) int {
	rv := 0
	for i := 0; i < BIT_LEN; i++ {
		pow := int(math.Pow(2, float64(BIT_LEN-i-1)))
		switch mask.value[i] {
		case MaskBitOn:
			rv += pow
		case MaskBitIgnore:
			if value >= pow {
				rv += pow
			}
		}
		if value >= pow {
			value -= pow
		}
	}
	return rv
}

func applyMaskWithFloatingBits(mask *Mask, value int) []int {
	rv := make([]int, 1)
	for i := 0; i < BIT_LEN; i++ {
		pow := int(math.Pow(2, float64(BIT_LEN-i-1)))
		switch mask.value[i] {
		case MaskBitOn:
			for j := 0; j < len(rv); j++ {
				rv[j] += pow
			}
		case MaskBitIgnore:
			toAdd := make([]int, 0)
			for _, v := range rv {
				toAdd = append(toAdd, v+pow)
			}
			rv = append(rv, toAdd...)
		case MaskBitOff:
			if value >= pow {
				for j := 0; j < len(rv); j++ {
					rv[j] += pow
				}
			}
		}
		if value >= pow {
			value -= pow
		}
	}
	return rv
}

func Run(infile string) {

	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	instructions := parseFile(fh)

	processor := Processor{memory: map[int]int{}, decoderVersion: 1}
	initializeProgram(instructions, &processor)
	initializationValue := computeAllValuesInMemory(&processor)
	fmt.Printf("The initialization value for part 1 is: %d\n", initializationValue)

	if infile != "day14/input_test" {
		processor = Processor{memory: map[int]int{}, decoderVersion: 2}
		initializeProgram(instructions, &processor)
		initializationValue = computeAllValuesInMemory(&processor)
		fmt.Printf("The initialization value for part 2 is: %d\n", initializationValue)
	}

}

func parseFile(reader io.Reader) []Instruction {
	instructions := make([]Instruction, 0)

	maskPattern, err := regexp.Compile(fmt.Sprintf("mask = ([X01]{%d})", BIT_LEN))
	if err != nil {
		panic(err)
	}

	memoryPattern, err := regexp.Compile("mem\\[(\\d+)] = (\\d+)")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if mask := maskPattern.FindStringSubmatch(line); mask != nil {
			value := [BIT_LEN]MaskBit{}
			for i, c := range mask[1] {
				switch c {
				case 'X':
					value[i] = MaskBitIgnore
				case '0':
					value[i] = MaskBitOff
				case '1':
					value[i] = MaskBitOn
				}
			}
			instructions = append(instructions, Mask{value})
		} else if memory := memoryPattern.FindStringSubmatch(line); memory != nil {
			location, _ := strconv.Atoi(memory[1])
			value, _ := strconv.Atoi(memory[2])
			instructions = append(instructions, MemorySet{
				location: location,
				value:    value,
			})
		} else {
			log.Fatalf("Enable to parse the line: %s\n", line)
		}
	}

	return instructions
}

func initializeProgram(instructions []Instruction, processor *Processor) {
	for _, instruction := range instructions {
		instruction.process(processor)
	}
}

func computeAllValuesInMemory(processor *Processor) int {
	rv := 0
	for _, value := range processor.memory {
		rv += value
	}
	return rv
}
