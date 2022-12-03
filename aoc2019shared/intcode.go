package aoc2019shared

import (
	"log"
	"math"
	"strconv"
	"strings"
	"sync"
)

type operation int

const (
	add operation = iota + 1
	multiply
	input
	output
	jumpTrue
	jumpFalse
	lessThan
	equals
	relativeBaseOffset
	eof
)

func isValid(inst int64) bool {
	return inst >= int64(add) && inst < int64(eof)
}

// IntCodeInterpreter is a interpreter for the int code language defined in
// Advent Of Code 2019 day 2 and 5 (and more!)
type IntCodeInterpreter struct {
	name         string
	ip           int64
	Inst         []int64
	RelativeBase int64
	Input        chan int64
	Output       chan int64
}

// NewIntCodeInterpreter creates an int code interpreter with the
// given instructions.
func NewIntCodeInterpreter(name, input string) *IntCodeInterpreter {
	interpreter := IntCodeInterpreter{
		name:         name,
		Inst:         parseInstructions(input),
		ip:           0,
		RelativeBase: 0,
		Input:        make(chan int64, 2),
		Output:       make(chan int64, 2),
	}

	return &interpreter
}

// Process runs the program in the IntCodeInterpreter's instructions. It returns
// the value in the 0 instruction at the end.
func (ici *IntCodeInterpreter) Process(wg *sync.WaitGroup) int64 {
	for {
		oper := ici.Inst[ici.ip] % 100

		if !isValid(oper) {
			if wg != nil {
				wg.Done()
			}
			return ici.Inst[0]
		}

		switch operation(oper) {
		case add:
			ici.setParam(3, ici.getParam(1)+ici.getParam(2))
			ici.ip += 4
			break

		case multiply:
			ici.setParam(3, ici.getParam(1)*ici.getParam(2))
			ici.ip += 4
			break

		case input:
			ici.setParam(1, <-ici.Input)
			ici.ip += 2
			break

		case output:
			ici.Output <- ici.getParam(1)
			ici.ip += 2
			break

		case jumpTrue:
			if ici.getParam(1) != 0 {
				ici.ip = ici.getParam(2)
			} else {
				ici.ip += 3
			}
			break

		case jumpFalse:
			if ici.getParam(1) == 0 {
				ici.ip = ici.getParam(2)
			} else {
				ici.ip += 3
			}
			break

		case lessThan:
			if ici.getParam(1) < ici.getParam(2) {
				ici.setParam(3, 1)
			} else {
				ici.setParam(3, 0)
			}
			ici.ip += 4
			break

		case equals:
			if ici.getParam(1) == ici.getParam(2) {
				ici.setParam(3, 1)
			} else {
				ici.setParam(3, 0)
			}
			ici.ip += 4
			break

		case relativeBaseOffset:
			ici.RelativeBase += ici.getParam(1)
			ici.ip += 2
			break
		}
	}
}

func parseInstructions(input string) []int64 {
	output := make([]int64, math.MaxUint16)

	split := strings.Split(input, ",")

	for i, strVal := range split {
		intVal, err := strconv.ParseInt(strVal, 10, 0)

		if err != nil {
			log.Fatal(err)
		}

		output[i] = intVal
	}

	return output
}

func (ici IntCodeInterpreter) getParam(num int) int64 {
	if num < 1 || num > 3 {
		log.Fatalf("Invalid parameter number: %d\n", num)
	}

	val := ici.Inst[ici.ip+int64(num)]
	inst := ici.Inst[ici.ip]
	mult := 10
	for i := 0; i < num; i++ {
		mult *= 10
	}

	mode := (inst / int64(mult)) % 10

	if mode == 0 {
		return ici.Inst[val]
	} else if mode == 1 {
		return val
	}

	// mode 2
	return ici.Inst[val+ici.RelativeBase]

}

func (ici IntCodeInterpreter) setParam(index, num int64) {
	val := ici.Inst[ici.ip+int64(index)]
	inst := ici.Inst[ici.ip]
	mult := 10
	for i := 0; i < int(index); i++ {
		mult *= 10
	}

	mode := (inst / int64(mult)) % 10

	if mode == 0 {
		ici.Inst[val] = num
		return
	}

	// cannot set mode 1 values

	// mode 2
	ici.Inst[val+ici.RelativeBase] = num
}
