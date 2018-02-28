package main

import (
	"bufio"
	_ "errors"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"gopkg.in/cheggaaa/pb.v1"
)

type Problem struct {
	numLines int
	numbers  []int
	data     []string
}

type Solution struct {
	answer int
}

func Solve(problem *Problem) (*Solution, error) {
	var solution Solution

	count := 1500
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.FinishPrint("DONE")
	solution.answer = 42

	return &solution, nil
}

func ParseInputFile(filePath string) (*Problem, error) {
	var problem Problem
	problem.numbers = make([]int, 6)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Can't open file for reading")
		log.Fatal(err)
		return &problem, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Printf("%d: %s\n", lineNumber, line)

		if lineNumber == 0 {
			// Header line
			problem.numLines, _ = strconv.Atoi(line)
			problem.data = make([]string, problem.numLines)
		} else if lineNumber == 1 {
			// Numbers line
			parts := strings.Split(line, " ")
			for i, part := range parts {
				problem.numbers[i], _ = strconv.Atoi(part)
			}
		} else {
			// Data lines
			problem.data[lineNumber-2] = line
		}

		lineNumber++
	}

	return &problem, nil
}

func SaveSolutionFile(solution *Solution, filePath string) {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("%d\n", solution.answer))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Print("Usage: hashcode2018 <path-to-input>")
		return
	}

	fmt.Println("Parsing input file...")
	inputPath := os.Args[1]
	problem, err := ParseInputFile(inputPath)
	if err != nil {
		log.Fatal("Couldn't parse input file")
		log.Fatal(err)
		return
	}

	fmt.Printf("Problem%+v\n", *problem)

	fmt.Println("Solving...")
	solution, err := Solve(problem)
	if err != nil {
		log.Fatal("Error whilst solving problem")
		log.Fatal(err)
		return
	} else {
		_, fileName := path.Split(inputPath)
		solutionPath := "solutions/" + fileName
		fmt.Println("Saving solution to", solutionPath)
		SaveSolutionFile(solution, solutionPath)
	}
}
