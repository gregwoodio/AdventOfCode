package day09

import (
	"fmt"
	"strconv"
)

type encoding struct {
	preambleLen int
	nums        []int64
	dict        map[int64]bool
}

func newEncoding(preambleLen int, nums []int64) *encoding {
	dict := make(map[int64]bool)
	nums = nums[:preambleLen]

	for _, num := range nums {
		dict[num] = true
	}

	return &encoding{
		preambleLen: preambleLen,
		nums:        nums,
		dict:        dict,
	}
}

func (e *encoding) isValid(num int64) bool {
	for i := 0; i < e.preambleLen; i++ {
		_, ok := e.dict[num-e.nums[i]]
		if ok {
			return true
		}
	}

	return false
}

func (e *encoding) add(num int64) {
	toDelete := e.nums[0]
	e.nums = e.nums[1:]
	e.nums = append(e.nums, num)

	delete(e.dict, toDelete)
	e.dict[num] = true
}

func solvePartOne(preambleLen int, input []string) int64 {
	nums := toInt64Slice(input)
	enc := newEncoding(preambleLen, nums)
	for i := preambleLen; i < len(nums); i++ {
		if enc.isValid(nums[i]) {
			enc.add(nums[i])
		} else {
			return nums[i]
		}
	}

	return -1
}

func solvePartTwo(invalidNum int64, input []string) int64 {
	nums := toInt64Slice(input)

	var i int64
	for i = 0; i < int64(len(nums)); i++ {
		current := nums[i]
		contigRange := []int64{nums[i]}
		for j := i + 1; j < int64(len(nums)); j++ {
			current += nums[j]
			contigRange = append(contigRange, nums[j])
			if current == invalidNum {
				low := contigRange[0]
				high := contigRange[0]
				for _, val := range contigRange {
					if val < low {
						low = val
					}

					if val > high {
						high = val
					}
				}

				return low + high
			}

			if current > invalidNum {
				break
			}
		}
	}

	return -1
}

func toInt64Slice(input []string) []int64 {
	nums := []int64{}
	for _, line := range input {
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Errorf("could not parse int64: '%s'", line)
		}

		nums = append(nums, num)
	}
	return nums
}
