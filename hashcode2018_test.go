package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInputFile(t *testing.T) {
	problem, _ := ParseInputFile("data/test.in")
	assert.Equal(t, problem.numLines, 5)
	assert.Equal(t, problem.numbers[3], 16)
	assert.Equal(t, problem.data[0], "lol")
}

func TestSolve(t *testing.T) {
	problem, _ := ParseInputFile("data/test.in")
	solution, _ := Solve(problem)
	assert.Equal(t, solution.answer, 42)
}
