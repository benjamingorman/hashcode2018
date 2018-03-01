package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderRidesByEndTime(t *testing.T) {
	rides := make([]*Ride, 3)
	r0 := Ride{}
	r0.start = 0
	r0.finish = 5

	r1 := Ride{}
	r1.start = 11
	r1.finish = 15

	r2 := Ride{}
	r2.start = 6
	r2.finish = 10

	rides[0] = &r0
	rides[1] = &r1
	rides[2] = &r2

	fmt.Print(&r0, &r1, &r2)
	sortedRides := OrderRidesByEndTime(ridesList(rides))
	assert.Equal(t, sortedRides[0], &r0)
	assert.Equal(t, sortedRides[1], &r2)
	assert.Equal(t, sortedRides[2], &r1)
}

func TestParseInputFile(t *testing.T) {
	problem, _ := ParseInputFile("data/a_example.in")
	assert.Equal(t, problem.rows, 3)
	assert.Equal(t, problem.cols, 4)
	assert.Equal(t, problem.fleetSize, 2)
	assert.Equal(t, problem.numRides, 3)
	assert.Equal(t, problem.onTimeBonus, 2)
	assert.Equal(t, problem.numTimesteps, 10)

	assert.Equal(t, len(problem.rides), 3)
	ride := problem.rides[0]
	assert.Equal(t, ride.startY, 0)
	assert.Equal(t, ride.startX, 0)
	assert.Equal(t, ride.endY, 1)
	assert.Equal(t, ride.endX, 3)
	assert.Equal(t, ride.start, 2)
	assert.Equal(t, ride.finish, 9)
}

func TestSolve(t *testing.T) {
	//problem, _ := ParseInputFile("data/a_example.in")
	//solution, _ := Solve(problem)
	//Solve(problem)
	//assert.Equal(t, solution.answer, 42)
}

func TestRideDistance(t *testing.T) {
	testRide := Ride{startX: 0, startY: 0, endX: 2, endY: 2}
	assert.Equal(t, testRide.Distance(), 4)
}

func TestAreRidesCompatible(t *testing.T) {
	r1 := Ride{startX: 0, startY: 0, endX: 1, endY: 1, start: 0, finish: 5}
	r2 := Ride{startX: 1, startY: 2, endX: 1, endY: 4, start: 5, finish: 10}
	assert.Equal(t, AreRidesCompatible(&r1, &r2), true)

	r3 := Ride{startX: 1, startY: 2, endX: 1, endY: 5, start: 5, finish: 5}
	assert.Equal(t, AreRidesCompatible(&r1, &r3), false)
}

func TestAreRidesCompatibleConcrete(t *testing.T) {
	r1 := Ride{startX: 0, startY: 0, endX: 1, endY: 1, start: 0, finish: 5}
	r2 := Ride{startX: 1, startY: 2, endX: 1, endY: 4, start: 5, finish: 10}
	assert.Equal(t, AreRidesCompatibleConcrete(&r1, &r2, 8), true)
	assert.Equal(t, AreRidesCompatibleConcrete(&r1, &r2, 5), false)
}

func TestLatestPossibleStartTime(t *testing.T) {
	r1 := Ride{startX: 0, startY: 0, endX: 0, endY: 5, start: 0, finish: 7}
	assert.Equal(t, r1.LatestPossibleStartTime(), 1)
}

func TestEarliestPossibleFinishTime(t *testing.T) {
	r1 := Ride{startX: 0, startY: 0, endX: 0, endY: 5, start: 0, finish: 7}
	assert.Equal(t, r1.EarliestPossibleFinishTime(), 5)
}

func TestGreedyCarRoute(t *testing.T) {
}

func TestTravelTime(t *testing.T) {
	r1 := Ride{startX: 0, startY: 0, endX: 0, endY: 5, start: 0, finish: 7}
	r2 := Ride{startX: 0, startY: 0, endX: 0, endY: 5, start: 0, finish: 7}
	assert.Equal(t, 0, (&r1).TravelTime(&r2))
}
