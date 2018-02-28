package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgMaxFunc(t *testing.T) {
	numbers := []int{4, 8, 15, 16, 23, 42}
	maxItem, maxValue, maxIndex := ArgMaxFunc(numbers, func(x int) int { return x % 7 })
	assert.Equal(t, maxItem, 4)
	assert.Equal(t, maxValue, 4)
	assert.Equal(t, maxIndex, 0)
}

func TestArgMinFunc(t *testing.T) {
	numbers := []int{4, 8, 15, 16, 23, 42}
	minItem, minValue, minIndex := ArgMinFunc(numbers, func(x int) int { return x % 7 })
	assert.Equal(t, minItem, 42)
	assert.Equal(t, minValue, 0)
	assert.Equal(t, minIndex, 5)
}

func TestArgMax(t *testing.T) {
	numbers := []int{4, 8, 15, 16, 23, 42}
	maxValue, maxIndex := ArgMax(numbers)
	assert.Equal(t, maxValue, 42)
	assert.Equal(t, maxIndex, 5)
}

func TestArgMin(t *testing.T) {
	numbers := []int{4, 8, 15, 16, 23, 42}
	minValue, minIndex := ArgMin(numbers)
	assert.Equal(t, minValue, 4)
	assert.Equal(t, minIndex, 0)
}
