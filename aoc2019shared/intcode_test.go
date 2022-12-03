package aoc2019shared

import (
	"fmt"
	"sync"
	"testing"
)

type testData struct {
	input          string
	expectedZero   int64
	expectedOutput int64
	userInput      int64
}

func TestProcessIntCode(t *testing.T) {
	testDatas := []testData{
		testData{
			input:          "3,1,4,1,99",
			expectedOutput: 42,
			expectedZero:   3,
			userInput:      42,
		},
		// Day 5 tests (less than and equal operators)
		testData{
			input:          "3,9,8,9,10,9,4,9,99,-1,8",
			expectedOutput: 1,
			expectedZero:   3,
			userInput:      8,
		},
		testData{
			input:          "3,9,8,9,10,9,4,9,99,-1,8",
			expectedOutput: 0,
			expectedZero:   3,
			userInput:      7,
		},
		testData{
			input:          "3,9,7,9,10,9,4,9,99,-1,8",
			expectedOutput: 0,
			expectedZero:   3,
			userInput:      8,
		},
		testData{
			input:          "3,9,7,9,10,9,4,9,99,-1,8",
			expectedOutput: 1,
			expectedZero:   3,
			userInput:      7,
		},
		testData{
			input:          "3,3,1108,-1,8,3,4,3,99",
			expectedOutput: 1,
			expectedZero:   3,
			userInput:      8,
		},
		testData{
			input:          "3,3,1108,-1,8,3,4,3,99",
			expectedOutput: 0,
			expectedZero:   3,
			userInput:      7,
		},
		testData{
			input:          "3,3,1107,-1,8,3,4,3,99",
			expectedOutput: 0,
			expectedZero:   3,
			userInput:      8,
		},
		testData{
			input:          "3,3,1107,-1,8,3,4,3,99",
			expectedOutput: 1,
			expectedZero:   3,
			userInput:      7,
		},
		// Day 5 tests (jump operators)
		testData{
			input:          "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			expectedOutput: 0,
			expectedZero:   3,
			userInput:      0,
		},
		testData{
			input:          "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			expectedOutput: 1,
			expectedZero:   3,
			userInput:      999,
		},
		testData{
			input:          "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			expectedOutput: 0,
			expectedZero:   3,
			userInput:      0,
		},
		testData{
			input:          "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			expectedOutput: 1,
			expectedZero:   3,
			userInput:      999,
		},
	}

	for _, td := range testDatas {
		ici := NewIntCodeInterpreter("test", td.input)

		var wg sync.WaitGroup

		wg.Add(1)
		go ici.Process(&wg)

		ici.Input <- td.userInput
		out := <-ici.Output

		wg.Wait()

		if out != td.expectedOutput {
			t.Errorf("Expected %d but was %d\n", td.expectedOutput, out)
		}
	}
}

func TestProcessIntCodeDay05_01(t *testing.T) {
	testCases := []testData{
		testData{
			input: "1002,4,3,4,33",
		},
		testData{
			input: "1101,100,-1,4,0",
		},
	}

	for _, td := range testCases {
		ici := NewIntCodeInterpreter("test", td.input)

		var wg sync.WaitGroup

		wg.Add(1)
		go ici.Process(&wg)
		wg.Wait()

		if ici.Inst[4] != 99 {
			t.Errorf("Expected instruction 4 to be %d but was %d", 99, ici.Inst[4])
		}
	}
}

func TestProcessIntCodeDay09_01(t *testing.T) {
	td := testData{
		input: "1102,34915192,34915192,7,4,7,99,0",
	}

	ici := NewIntCodeInterpreter("test", td.input)

	var wg sync.WaitGroup

	wg.Add(1)
	go ici.Process(&wg)

	ici.Input <- td.userInput
	out := <-ici.Output

	wg.Wait()

	fmt.Println(out)
	outStr := fmt.Sprintf("%d", out)
	if len(outStr) != 16 {
		t.Errorf("Expected output was not a 16 digit number")
	}
}

func TestProcessIntCodeDay09_02(t *testing.T) {
	td := testData{
		input:          "104,1125899906842624,99",
		expectedOutput: 1125899906842624,
	}

	ici := NewIntCodeInterpreter("test", td.input)

	var wg sync.WaitGroup

	wg.Add(1)
	go ici.Process(&wg)

	ici.Input <- td.userInput
	out := <-ici.Output

	wg.Wait()

	if out != td.expectedOutput {
		t.Errorf("Expected %d but was %d\n", td.expectedOutput, out)
	}
}

func TestProcessIntCodeDay09_03(t *testing.T) {
	td := testData{
		input:          "9,20,204,-5,99,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20",
		expectedOutput: 15,
	}

	ici := NewIntCodeInterpreter("test", td.input)

	var wg sync.WaitGroup

	wg.Add(1)
	go ici.Process(&wg)

	ici.Input <- td.userInput
	out := <-ici.Output

	wg.Wait()

	if out != td.expectedOutput {
		t.Errorf("Expected %d but was %d\n", td.expectedOutput, out)
	}
}

// Additional testcases supplied by /u/Kache
func TestProcessIntCodeDay09_04(t *testing.T) {
	testCases := []testData{
		testData{
			input:          "109,-1,4,1,99",
			expectedOutput: -1,
		},
		testData{
			input:          "109,-1,104,1,99",
			expectedOutput: 1,
		},
		testData{
			input:          "109,-1,204,1,99",
			expectedOutput: 109,
		},
		testData{
			input:          "109,1,9,2,204,-6,99",
			expectedOutput: 204,
		},
		testData{
			input:          "109,1,109,9,204,-6,99",
			expectedOutput: 204,
		},
		testData{
			input:          "109,1,209,-1,204,-106,99",
			expectedOutput: 204,
		},
		testData{
			input:          "109,1,3,3,204,2,99",
			userInput:      42,
			expectedOutput: 42,
		},
		testData{
			input:          "109,1,203,2,204,2,99",
			userInput:      42,
			expectedOutput: 42,
		},
	}

	for _, td := range testCases {
		ici := NewIntCodeInterpreter("test", td.input)

		var wg sync.WaitGroup

		wg.Add(1)
		go ici.Process(&wg)

		ici.Input <- td.userInput
		out := <-ici.Output

		wg.Wait()
		fmt.Println(out)

		if out != td.expectedOutput {
			t.Errorf("Expected %d but was %d\n", td.expectedOutput, out)
		}
	}
}
