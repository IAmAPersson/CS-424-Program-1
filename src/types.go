/*
	Phil Lane
	02/17/2020
	CS-424-02
	Mary Allen

	Programming Assignment 1 (Golang)
*/

package main

//Type: Inputted batter information
type BatterInfo struct {
	firstName string
	lastName string
	plateAppearances uint64
	atBats uint64
	singles uint64
	doubles uint64
	triples uint64
	homeRuns uint64
	walks uint64
	hitByPitch uint64
}

//Type: Calculated batter information
type CalculatedBatterInfo struct {
	firstName string
	lastName string
	average float64
	slugging float64
	onBase float64
}
