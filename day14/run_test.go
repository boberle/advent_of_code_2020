package day14

import (
	"github.com/stretchr/testify/assert"
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name  string
		input string
		want  []Instruction
	}{
		{"mask", "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", []Instruction{
			Mask{value: [36]MaskBit{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 0, 2}},
		}},
		{"memory", "mem[123] = 456", []Instruction{
			MemorySet{location: 123, value: 456},
		}},
		{"mask and memory", "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X\nmem[123] = 456", []Instruction{
			Mask{value: [36]MaskBit{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 0, 2}},
			MemorySet{location: 123, value: 456},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			if got := parseFile(reader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_applyMask(t *testing.T) {
	mask := Mask{[36]MaskBit{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 0, 2}}
	tests := []struct {
		name  string
		mask  Mask
		value int
		want  int
	}{
		{"11 -> 73", mask, 11, 73},
		{"101 -> 101", mask, 101, 101},
		{"0 -> 64", mask, 0, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyMask(&mask, tt.value); got != tt.want {
				t.Errorf("applyMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initializeProgram(t *testing.T) {
	instructions := []Instruction{
		Mask{[36]MaskBit{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 0, 2}},
		MemorySet{8, 11},
		MemorySet{7, 101},
		MemorySet{8, 0},
	}
	processor := Processor{memory: map[int]int{}, decoderVersion: 1}
	initializeProgram(instructions, &processor)
	want := map[int]int{
		7: 101,
		8: 64,
	}
	assert.Equal(t, want, processor.memory)
}

func Test_applyMaskWithFloatingBits(t *testing.T) {
	mask1 := Mask{[36]MaskBit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 0, 0, 1, 2}}
	mask2 := Mask{[36]MaskBit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 2, 2}}
	tests := []struct {
		name  string
		mask  Mask
		value int
		want  []int
	}{
		{"case 1", mask1, 42, []int{26, 58, 27, 59}},
		{"case 2", mask2, 26, []int{16, 24, 18, 26, 17, 25, 19, 27}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, applyMaskWithFloatingBits(&tt.mask, tt.value), tt.name)
		})
	}
}
