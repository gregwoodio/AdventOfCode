package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var day, year int
	var help bool

	flag.IntVar(&day, "d", -1, "Specify day")
	flag.IntVar(&year, "y", -2, "Specify year")
	flag.BoolVar(&help, "h", false, "Show the help")

	flag.Parse()

	if help {
		printHelp()
		return
	}

	if day <= 0 || day > 25 {
		return
	}

	dayStr := fmt.Sprintf("day%02d", day)
	yearStr := fmt.Sprintf("%d", year)

	_, err := os.Open(yearStr)
	if err != nil {
		// try to create directory
		err = os.Mkdir(yearStr, os.ModeDir)
		if err != nil {
			return
		}
	}

	path := fmt.Sprintf("%s/%s", yearStr, dayStr)
	err = os.Mkdir(path, os.ModeDir)
	if err != nil {
		return
	}

	err = makePuzzleFile(dayStr, path)
	if err != nil {
		return
	}

	err = makeTestFile(day, year, path)
	if err != nil {
		return
	}

	err = makeInputFile(day, path)
	if err != nil {
		return
	}
}

func makePuzzleFile(day, path string) error {
	puzzleFile := fmt.Sprintf("%s/%s.go", path, day)
	file, err := os.Create(puzzleFile)
	if err != nil {
		return fmt.Errorf("unable to create file %s", puzzleFile)
	}
	defer file.Close()

	fmt.Fprintf(file, "package %s\n\n", day)
	fmt.Fprintf(file, "func solvePartOne(input []string) int {\n\treturn -1\n}\n\n")
	fmt.Fprintf(file, "func solvePartTwo(input []string) int {\n\treturn -1\n}\n")

	return nil
}

func makeTestFile(day, year int, path string) error {
	testFile := fmt.Sprintf("%s/day%02d_test.go", path, day)
	file, err := os.Create(testFile)
	if err != nil {
		return fmt.Errorf("unable to create file %s", testFile)
	}
	defer file.Close()

	fmt.Fprintf(file, "package day%02d\n\n", day)
	fmt.Fprintf(file, "import (\n\t\"fmt\"\n\t\"testing\"\n\n\t\"github.com/gregwoodio/aocutil\"\n)\n\n")

	fmt.Fprintf(file, "func Test%dDay%02d_SolvePartOne(t *testing.T) {\n", year, day)
	fmt.Fprintf(file, "\tinputs := []string{\n\t\t\"foo\",\n\t}\n")
	fmt.Fprintf(file, "\n\texpected := -1\n\tactual := solvePartOne(inputs)\n")
	fmt.Fprintf(file, "\tif actual != expected {\n")
	fmt.Fprint(file, "\t\tt.Errorf(\"Expected %"+"d but was %"+"d\", expected, actual)\n\t}\n}")

	fmt.Fprintf(file, "\n\nfunc Test%dDay%02d_SolvePartOneActual(t *testing.T) {\n", year, day)
	fmt.Fprintf(file, "\tif testing.Short() {\n")
	fmt.Fprintf(file, "\t\tt.Skip(\"Skipping actual solution\")\n\t}\n\n")
	fmt.Fprintf(file, "\tsolution := solvePartOne(aocutil.ReadStringsFromFile(\"./day%02d_input.txt\"))\n", day)
	fmt.Fprintf(file, "\tfmt.Println(solution)\n}")

	fmt.Fprintf(file, "\n\nfunc Test%dDay%02d_SolvePartTwo(t *testing.T) {\n\n", year, day)
	fmt.Fprintf(file, "\tinputs := []string{\n\t\t\"foo\",\n\t}\n\n")
	fmt.Fprintf(file, "\texpected := -1\n\tactual := solvePartTwo(inputs)\n")
	fmt.Fprintf(file, "\tif actual != expected {\n")
	fmt.Fprint(file, "\t\tt.Errorf(\"Expected %"+"d but was %"+"d\", expected, actual)\n\t}\n}")

	fmt.Fprintf(file, "\n\nfunc Test%dDay%02d_SolvePartTwoActual(t *testing.T) {\n\n", year, day)
	fmt.Fprintf(file, "\tif testing.Short() {\n")
	fmt.Fprintf(file, "\t\tt.Skip(\"Skipping actual solution\")\n\t}\n\n")
	fmt.Fprintf(file, "\tsolution := solvePartOne(aocutil.ReadStringsFromFile(\"./day%02d_input.txt\"))\n", day)
	fmt.Fprintf(file, "\tfmt.Println(solution)\n}")

	return nil
}

func makeInputFile(day int, path string) error {
	inputFile := fmt.Sprintf("%s/day%02d_input.txt", path, day)
	file, err := os.Create(inputFile)
	if err != nil {
		return fmt.Errorf("unable to create file %s", inputFile)
	}
	defer file.Close()

	return nil
}

func printHelp() {
	fmt.Println("NAME")
	fmt.Println("\tboilerplate - Create boilerplate code for Advent of Code puzzles")

	fmt.Println("\nDESCRIPTION")
	fmt.Println("\tMandatory arguments")
	fmt.Println("\n\t-y\tYear")
	fmt.Println("\t-d\tDay")
}
