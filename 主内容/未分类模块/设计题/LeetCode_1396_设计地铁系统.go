package main

import "fmt"

// ------------------------- Arrival -------------------------
type Arrival struct {
	StationName string
	TimeStamp   int
}

func NewArrival(stationName string, timeStamp int) *Arrival {
	return &Arrival{stationName, timeStamp}
}

// ------------------------- Trip -------------------------
type Trip struct {
	Source      *Arrival
	Destination *Arrival
}

func NewTrip() *Trip {
	return &Trip{}
}

func (t *Trip) GetBackAndForthCost() int {
	return t.Destination.TimeStamp - t.Source.TimeStamp
}

// ------------------------- UndergroundSystem -------------------------
type UndergroundSystem struct {
	userIdToTrip      map[int]*Trip
	backAndForthCosts map[string][]int // the costs of startStation to endStation
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		make(map[int]*Trip, 500),
		make(map[string][]int, 500),
	}
}

func (udg *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	udg.recordUserTripStart(id, stationName, t)
}

func (udg *UndergroundSystem) CheckOut(id int, stationName string, timeStamp int) {
	udg.recordUserTripEnd(id, stationName, timeStamp)
	udg.appendNewBackAndForthCost(id)
}

func (udg *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	tripIdentity := formTripIdentity(startStation, endStation)
	return getAverage(udg.backAndForthCosts[tripIdentity])
}
func (udg *UndergroundSystem) recordUserTripStart(id int, startStation string, timeStamp int) {
	trip := NewTrip()
	trip.Source = NewArrival(startStation, timeStamp)
	udg.userIdToTrip[id] = trip
}
func (udg *UndergroundSystem) recordUserTripEnd(id int, endStation string, timeStamp int) {
	trip := udg.userIdToTrip[id]
	trip.Destination = NewArrival(endStation, timeStamp)
}
func (udg *UndergroundSystem) appendNewBackAndForthCost(id int) {
	trip := udg.userIdToTrip[id]
	tripIdentity := formTripIdentity(trip.Source.StationName, trip.Destination.StationName)
	tripCost := trip.GetBackAndForthCost()
	udg.backAndForthCosts[tripIdentity] = append(udg.backAndForthCosts[tripIdentity], tripCost)
}

func getAverage(array []int) float64 {
	sum := 0.0
	for _, element := range array {
		sum += float64(element)
	}
	return sum / float64(len(array))
}

// this function is to assist udg.backAndForthCosts
func formTripIdentity(startStation string, endStation string) string {
	return fmt.Sprintf("%s->%s", startStation, endStation)
}
