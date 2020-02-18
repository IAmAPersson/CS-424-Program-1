/*
	Phil Lane
	02/17/2020
	CS-424-02
	Mary Allen

	Programming Assignment 1 (Golang)
*/

package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

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

type CalculatedBatterInfo struct {
	firstName string
	lastName string
	average float64
	slugging float64
	onBase float64
}

func main() {
	fmt.Println("Welcome to the player statistics calculator test program! I am going to\n" +
		"read players from an input data file. You will tell me the name of your\n" +
		"input file. I will store all of the players in a list, compute each player's\n" +
		"averages, and then write the resulting team report to your output file!\n")
		
	fmt.Print("Enter the name of your input file: ")
	
	reader := bufio.NewReader(os.Stdin)
	path, _ := reader.ReadString('\n')
	path = path[0 : len(path) - 1]
	
	data, err := ReadInFile(path)
	
	if err != nil {
		return
	}
	
	batters, badlines := ParseInfo(data)
	batters = PlayerSort(batters)
	
	calcData := Calculate(batters)
	
	fmt.Println()
	fmt.Printf("BASEBALL TEAM REPORT --- %d PLAYERS FOUND IN FILE\n", len(batters))
	fmt.Printf("OVERALL BATTING AVERAGE is %0.3f\n", Average(calcData))
	
	for i := 0; i < len(badlines); i++ {
		fmt.Println(badlines[i])
	}
}

func ReadInFile(path string) (string, error) {
	file, err := os.Open(path)
	
	if err != nil {
		fmt.Println("Error opening file.")
		fmt.Println(err)
		return "", err
	}
	
	filedata, err := ioutil.ReadAll(file)
	
	if err != nil {
		fmt.Println("Error reading from file.")
		fmt.Println(err)
		return "", err
	}
	
	file.Close()
	
	return string(filedata), nil
}

func ParseInfo(data string) ([]BatterInfo, []string) {
	batters := []BatterInfo { }
	invalidstr := []string { }
	
	data = strings.Replace(data, "\r", "", -1)
	lines := strings.Split(data, "\n")
	
PrimaryLoop:
	for i := 0; i < len(lines); i++ {
		//fmt.Println(lines[i])
		
		var batter BatterInfo
		
		tokens := []string { }
		spaceSeparatedValues := strings.Split(lines[i], " ")
		for j := 0; j < len(spaceSeparatedValues); j++ {
			if spaceSeparatedValues[j] != "" {
				tokens = append(tokens, spaceSeparatedValues[j])
			}
		}
		
		if len(tokens) != 10 {
			invalidstr = append(invalidstr, "Invalid line entered (line " + strconv.Itoa(i) + ")-- incorrect number of parameters.")
			continue
		}
		
		var err error
		
		batter.firstName = tokens[0]
		batter.lastName = tokens[1]
		
		//this is to save a lot of lines and duplicated code
		batterNumericParts := [...]*uint64 { &batter.plateAppearances, &batter.atBats, &batter.singles, &batter.doubles, &batter.triples, &batter.homeRuns, &batter.walks, &batter.hitByPitch }
		for j := 0; j < 8; j++ {
			*batterNumericParts[j], err = strconv.ParseUint(tokens[j + 2], 10, 32)
			if err != nil {
				invalidstr = append(invalidstr, "Invalid line entered (line " + strconv.Itoa(i) + ")-- illegal type of parameter.")
				continue PrimaryLoop
			}
		}
		
		batters = append(batters, batter)
	}
	
	return batters, invalidstr
}

func PlayerSort(batters []BatterInfo) []BatterInfo {
	sort.Slice(batters, func(i, j int) bool {
		if batters[i].lastName != batters[j].lastName {
			return batters[i].lastName < batters[j].lastName
		} else {
			return batters[i].firstName < batters[j].firstName
		}
	})
	
	return batters
}

func Calculate(batters []BatterInfo) []CalculatedBatterInfo {
	 newbatters := make([]CalculatedBatterInfo, len(batters))
	
	for i := 0; i < len(batters); i++ {
		newbatters[i].firstName = batters[i].firstName
		newbatters[i].lastName = batters[i].lastName
		newbatters[i].average = float64(batters[i].singles + batters[i].doubles + batters[i].triples + batters[i].homeRuns) / float64(batters[i].atBats)
		newbatters[i].slugging = float64(batters[i].singles + 2 * batters[i].doubles + 3 * batters[i].triples + 4 * batters[i].homeRuns) / float64(batters[i].atBats)
		newbatters[i].onBase = float64(batters[i].singles + batters[i].doubles + batters[i].triples + batters[i].homeRuns + batters[i].walks + batters[i].hitByPitch) / float64(batters[i].plateAppearances)
	}
	
	return newbatters
}

func Average(batters []CalculatedBatterInfo) float64 {
	runningTotal := float64(0)
	
	for i := 0; i < len(batters); i++ {
		runningTotal += batters[i].average
	}
	
	return runningTotal / float64(len(batters))
}