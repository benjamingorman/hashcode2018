package main

import (
	"bufio"
	_ "errors"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	_ "time"

	"gopkg.in/cheggaaa/pb.v1"
)

type Problem struct {
	rows         int
	cols         int
	fleetSize    int
	numRides     int
	onTimeBonus  int
	numTimesteps int
	rides        []*Ride
}

type Ride struct {
	originalIndex int
	startX        int
	startY        int
	endX          int
	endY          int
	start         int
	finish        int
}

type Solution struct {
	routes [][]int
}

func (ride *Ride) Distance() int {
	return (ride.endY - ride.startY) + (ride.endX - ride.startX)
}

func (ride *Ride) LatestPossibleStartTime() int {
	return ride.finish - 1 - ride.Distance()
}

func (ride *Ride) EarliestPossibleFinishTime() int {
	return ride.start + ride.Distance()
}

func absInt(val int) int {
	if val < 0 {
		return -val
	} else {
		return val
	}
}

func (r1 *Ride) TravelTime(r2 *Ride) int {
	return absInt(r2.startX-r1.endX) + absInt(r2.startY-r1.endY)
}

type ridesList []*Ride

// Implement sort interface for ride list
func (rides ridesList) Len() int {
	return len(rides)
}
func (rides ridesList) Swap(i, j int) {
	rides[i], rides[j] = rides[j], rides[i]
}
func (rides ridesList) Less(i, j int) bool {
	r1 := rides[i]
	r2 := rides[j]
	if r1.finish < r2.finish {
		return true
	} else {
		return r1.start < r2.start
	}
}

func OrderRidesByEndTime(rides []*Ride) []*Ride {
	// Make a copy of the input to avoid mutating it
	sortedRides := make([]*Ride, len(rides))
	for i, r := range rides {
		sortedRides[i] = r
	}
	sort.Sort(ridesList(sortedRides))
	//fmt.Print(sortedRides)
	return sortedRides
}

func AreRidesCompatible(r1 *Ride, r2 *Ride) bool {
	if r1.originalIndex == r2.originalIndex {
		return false
	} else {
		return r1.EarliestPossibleFinishTime()+r1.TravelTime(r2) <= r2.LatestPossibleStartTime()
	}
}

func AreRidesCompatibleConcrete(r1 *Ride, r2 *Ride, r2_start int) bool {
	if r1.originalIndex == r2.originalIndex {
		return false
	} else {
		return r1.finish <= r2_start-r1.TravelTime(r2)
	}
}

// Assumes they are compatible
func RecommendConcreteStartTimes(r1 *Ride, r2 *Ride) (int, int) {
	r2_t := r2.LatestPossibleStartTime()
	r1_t := r2_t - r1.TravelTime(r2) - r1.Distance()
	return r1_t, r2_t
}

// Greedily computes a route using the given rides
// Returns a list of ints which are the indexes of each ride in the route
// Assumes rides is sorted
// original Rides passed as param to allow quick lookup by index
func GreedyCarRoute(originalRides []*Ride, sortedRides []*Ride, usedSet map[int]bool) []int {
	var route []int

	//fmt.Println("length of usedSet", len(usedSet))

	// Initialize the last ride, should probably be the first one from the end
	// that hasn't already been used
	var lastRideSeenIndex int = -1
	var lastRideStartTime int
	for i := len(sortedRides) - 1; i >= 0; i-- {
		if !usedSet[i] {
			ride := sortedRides[i]
			lastRideSeenIndex = i
			lastRideStartTime = ride.LatestPossibleStartTime()
			//fmt.Println("First ride", lastRideSeenIndex)
			// fmt.Printf("%+v\n", *ride)
			break
		}
	}
	if lastRideSeenIndex == -1 {
		panic("lastRideSeenIndex should always be definable")
	}

	route = append(route, lastRideSeenIndex)

	for i := len(sortedRides) - 1; i >= 0; i-- {
		// Skip if already used
		if usedSet[i] {
			//fmt.Println("Skipping used ride", i)
			continue
		} else if i >= lastRideSeenIndex {
			continue
		}

		lastRide := originalRides[lastRideSeenIndex]
		ride := sortedRides[i]

		if AreRidesCompatibleConcrete(ride, lastRide, lastRideStartTime) {
			//fmt.Println("Compatible:", i, lastRideSeenIndex)
			lastRideSeenIndex = i
			route = append(route, lastRideSeenIndex)

			recommended_t1, _ := RecommendConcreteStartTimes(ride, lastRide)
			lastRideStartTime = recommended_t1
		}
	}

	// Add every ride in the route to the used set
	for _, rideIndex := range route {
		usedSet[rideIndex] = true
	}

	// Convert between sorted routes indices and original
	for i, sortedRideIndex := range route {
		route[i] = sortedRides[sortedRideIndex].originalIndex
	}

	// Reverse the route (since it's back to front)
	return reverse(route)
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func Solve(problem *Problem) (*Solution, error) {
	var solution Solution
	solution.routes = make([][]int, problem.fleetSize)

	sortedRides := OrderRidesByEndTime(problem.rides)

	// The set of rides which we've allocated (their indices)
	bar := pb.StartNew(problem.fleetSize)
	usedSet := make(map[int]bool)
	for i := 0; i < problem.fleetSize; i++ {
		route := GreedyCarRoute(problem.rides, sortedRides, usedSet)
		solution.routes[i] = route
		bar.Increment()
		//fmt.Println(len(usedSet))

		if len(usedSet) == problem.numRides {
			break
		}
	}
	bar.FinishPrint("DONE")

	//fmt.Print(solution)
	return &solution, nil
}

func ParseInputFile(filePath string) (*Problem, error) {
	var problem Problem

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
			parts := strings.Split(line, " ")
			problem.rows, _ = strconv.Atoi(parts[0])
			problem.cols, _ = strconv.Atoi(parts[1])
			problem.fleetSize, _ = strconv.Atoi(parts[2])
			problem.numRides, _ = strconv.Atoi(parts[3])
			problem.onTimeBonus, _ = strconv.Atoi(parts[4])
			problem.numTimesteps, _ = strconv.Atoi(parts[5])
			problem.rides = make([]*Ride, problem.numRides)
		} else {
			// Ride description line
			parts := strings.Split(line, " ")
			ride := Ride{}
			ride.startY, _ = strconv.Atoi(parts[0])
			ride.startX, _ = strconv.Atoi(parts[1])
			ride.endY, _ = strconv.Atoi(parts[2])
			ride.endX, _ = strconv.Atoi(parts[3])
			ride.start, _ = strconv.Atoi(parts[4])
			ride.finish, _ = strconv.Atoi(parts[5])
			ride.originalIndex = lineNumber - 1
			problem.rides[lineNumber-1] = &ride
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

	for _, route := range solution.routes {
		f.WriteString(fmt.Sprintf("%d ", len(route)))
		f.WriteString(strings.Trim(strings.Join(strings.Split(fmt.Sprint(route), " "), " "), "[]"))
		f.WriteString("\n")
	}
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

	//fmt.Printf("Problem%+v\n", *problem)

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
