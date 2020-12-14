package day14

import (
	"fmt"
	"strconv"
	"strings"
)

func solvePartOne(input []string) int {
	var orMask, andMask int64
	mem := make(map[int]int64) // memory location to value

	for _, line := range input {
		parts := strings.Split(line, " = ")

		if parts[0] == "mask" {
			var orMaskStr string
			var andMaskStr string

			for _, ch := range parts[1] {
				if ch == 'X' {
					orMaskStr = orMaskStr + "0"
					andMaskStr = andMaskStr + "1"
				} else if ch == '1' {
					orMaskStr = orMaskStr + "1"
					andMaskStr = andMaskStr + "1"
				} else if ch == '0' {
					orMaskStr = orMaskStr + "0"
					andMaskStr = andMaskStr + "0"
				}

				var err error
				orMask, err = strconv.ParseInt(orMaskStr, 2, 64)
				if err != nil {
					fmt.Errorf("could not parse string '%s'", orMaskStr)
				}
				andMask, err = strconv.ParseInt(andMaskStr, 2, 64)
				if err != nil {
					fmt.Errorf("could not parse string '%s'", andMaskStr)
				}
			}
		} else if strings.HasPrefix(parts[0], "mem[") {
			addressStr := parts[0][4 : len(parts[0])-1]
			addr, err := strconv.Atoi(addressStr)
			if err != nil {
				fmt.Errorf("could not parse string '%s'", addressStr)
			}

			val, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				fmt.Errorf("could not parse string '%s'", parts[1])
			}

			val = val | orMask
			val = val & andMask

			mem[addr] = val
		}
	}

	var sum int64
	for _, v := range mem {
		sum += v
	}
	return int(sum)
}

func solvePartTwo(input []string) int64 {
	var mask string
	mem := make(map[int64]int64) // memory location to value

	for _, line := range input {
		parts := strings.Split(line, " = ")

		if parts[0] == "mask" {
			mask = parts[1]
		} else if strings.HasPrefix(parts[0], "mem[") {
			addressStr := parts[0][4 : len(parts[0])-1]
			addr, err := strconv.Atoi(addressStr)
			if err != nil {
				fmt.Errorf("could not parse string '%s'", addressStr)
			}

			val, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				fmt.Errorf("could not parse string '%s'", parts[1])
			}

			addressBinary := fmt.Sprintf("%b", addr)
			masks := expandMask(applyMask(addressBinary, mask))

			for _, m := range masks {
				mVal, err := strconv.ParseInt(m, 2, 64)
				if err != nil {
					fmt.Errorf("could not parse string '%s'", m)
				}
				mem[mVal] = val
			}
		}
	}

	var sum int64
	for _, v := range mem {
		sum += v
	}
	return sum
}

func expandMask(input string) []string {
	xIndex := strings.Index(input, "X")

	if xIndex == -1 {
		return []string{input}
	}

	bytes := []byte(input)
	bytes[xIndex] = '1'
	results := expandMask(string(bytes))
	bytes[xIndex] = '0'
	results = append(results, expandMask(string(bytes))...)

	return results
}

func applyMask(addr, mask string) string {
	str := ""
	for len(addr) < len(mask) {
		addr = "0" + addr
	}

	for i := 0; i < len(mask); i++ {
		if mask[i] == 'X' {
			str = str + "X"
		} else if mask[i] == '1' {
			str = str + "1"
		} else if mask[i] == '0' {
			str = str + string(addr[i])
		}
	}

	return str
}
