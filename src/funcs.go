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

//Function: Get the path of file to read from
func GetPath() string {
	//Prompt the user
	fmt.Println("Welcome to the player statistics calculator test program! I am going to\n" +
		"read players from an input data file. You will tell me the name of your\n" +
		"input file. I will store all of the players in a list, compute each player's\n" +
		"averages, and then write the resulting team report to your output file!\n")
		
	fmt.Print("Enter the name of your input file: ")
	
	//Create a reader and read in the path of the file
	reader := bufio.NewReader(os.Stdin)
	path, _ := reader.ReadString('\n')
	
	//Trim off the \n from the end
	path = path[0 : len(path) - 1]
	
	//Return the path
	return path
}

//Function: Read in a file
func ReadInFile(path string) (string, error) {
	//Create a file object to read in from
	file, err := os.Open(path)
	
	//If error opening file, print error and pass the error up the chain
	if err != nil {
		fmt.Println("Error opening file.")
		fmt.Println(err)
		return "", err
	}
	
	//Read in the file data
	filedata, err := ioutil.ReadAll(file)
	
	//If error reading from file, print error and pass the error up the chain (making sure to close the file)
	if err != nil {
		fmt.Println("Error reading from file.")
		fmt.Println(err)
		file.Close()
		return "", err
	}
	
	//Close the file
	file.Close()
	
	//Cast the []byte to a string and return
	return string(filedata), nil
}

//Function: Parse the input file into a slice of BatterInfo (defined in types.go) and a []string of all the error messages
func ParseInfo(data string) ([]BatterInfo, []string) {
	//Delcare two slices, a []BatterInfo to hold the parsed data for each batter, and a []string for all the error messages
	batters := []BatterInfo { }
	invalidstr := []string { }
	
	//Comb out the \r's and split on \n, getting each line separately in an array
	data = strings.Replace(data, "\r", "", -1)
	lines := strings.Split(data, "\n")
	
	//Loop: loop through each line of the input file and parse
PrimaryLoop:
	for i := 0; i < len(lines); i++ {
		//fmt.Println(lines[i])
		
		//Declare a temporary variable to mutate
		var batter BatterInfo
		
		//Declare two variables: one a slice of strings for each token in the line, and one that's simply the input line split on spaces
		tokens := []string { }
		spaceSeparatedValues := strings.Split(lines[i], " ")
		
		//Iterate through spaceSeparatedValues, and if the element is not empty, copy over to tokens
		//This ensures that the extra spaces are ignored
		for j := 0; j < len(spaceSeparatedValues); j++ {
			if spaceSeparatedValues[j] != "" {
				tokens = append(tokens, spaceSeparatedValues[j])
			}
		}
		
		//If there are not 10 tokens in the line, then append the error to the []string to be returned, and continue from the top with the next line
		if len(tokens) != 10 {
			invalidstr = append(invalidstr, "Invalid line entered (line " + strconv.Itoa(i) + ")-- incorrect number of parameters.")
			continue
		}
		
		//Create an object of type error
		var err error
		
		//Copy over the first and last names
		batter.firstName = tokens[0]
		batter.lastName = tokens[1]
		
		//This is to save a lot of lines and duplicated code
		//To avoid having to copy the assignment and error checking eight times over for each field, we put a pointer to each field in an array
		//Then, we iterate through the array, dereference the pointer, and store into that field
		//That technique allows us to avoid duplicating the assignment and if statement multiple times
		
		//The *uint64 array of all of the fields
		batterNumericParts := [...]*uint64 { &batter.plateAppearances, &batter.atBats, &batter.singles, &batter.doubles, &batter.triples, &batter.homeRuns, &batter.walks, &batter.hitByPitch }
		//Iterate through the array
		for j := 0; j < 8; j++ {
			//Attempt to parse the token as a base-10 32-bit Uint
			*batterNumericParts[j], err = strconv.ParseUint(tokens[j + 2], 10, 32)
			//If the parsing failed, append the error line to the []string, and continue from the outer loop
			//This is what the label was for
			if err != nil {
				invalidstr = append(invalidstr, "Invalid line entered (line " + strconv.Itoa(i) + ")-- illegal type of parameter.")
				continue PrimaryLoop
			}
		}
		
		//Append the temporary BatterInfo object to the BatterInfo slice
		batters = append(batters, batter)
	}
	
	//Return all the batters and the error lines
	return batters, invalidstr
}

//Function: Sort the players
func PlayerSort(batters []BatterInfo) []BatterInfo {
	//Sort the players
	sort.Slice(batters, func(i, j int) bool {
		//Sort by last name first, and if those are the same, then sort by first name
		if batters[i].lastName != batters[j].lastName {
			return batters[i].lastName < batters[j].lastName
		} else {
			return batters[i].firstName < batters[j].firstName
		}
	})
	
	//Return the sorted array
	return batters
}

//Function: Calculate all the data about the batters to a CalculatedBatterInfo slice
func Calculate(batters []BatterInfo) []CalculatedBatterInfo {
	//Make a slice of CalculatedBatterInfo with as many elements as there are in the batters slice
	newbatters := make([]CalculatedBatterInfo, len(batters))
	
	//Iterate through the batters slice
	for i := 0; i < len(batters); i++ {
		//Copy over the first and last name
		newbatters[i].firstName = batters[i].firstName
		newbatters[i].lastName = batters[i].lastName
		//Calculate the batting average
		newbatters[i].average = float64(batters[i].singles + batters[i].doubles + batters[i].triples + batters[i].homeRuns) / float64(batters[i].atBats)
		//Calculate the slugging average
		newbatters[i].slugging = float64(batters[i].singles + 2 * batters[i].doubles + 3 * batters[i].triples + 4 * batters[i].homeRuns) / float64(batters[i].atBats)
		//Calculate the on base percent
		newbatters[i].onBase = float64(batters[i].singles + batters[i].doubles + batters[i].triples + batters[i].homeRuns + batters[i].walks + batters[i].hitByPitch) / float64(batters[i].plateAppearances)
	}
	
	//Return the CalculatedBatterInfo slice
	return newbatters
}

//Function: Get the team batting average
func Average(batters []CalculatedBatterInfo) float64 {
	//Create a float64 for a running total, initialized to 0.0
	runningTotal := float64(0)
	
	//Iterate through the batters slice
	for i := 0; i < len(batters); i++ {
		//Add to the running total
		runningTotal += batters[i].average
	}
	
	//Return the running total divided by the length of the inputted slice (i.e. the average)
	return runningTotal / float64(len(batters))
}

//Function: Print the formatted data to the screen
func FormatData(batters []CalculatedBatterInfo, errorlines []string) {
	//Print the players found and the team average
	fmt.Printf("\nBASEBALL TEAM REPORT --- %d PLAYERS FOUND IN FILE\n", len(batters))
	fmt.Printf("OVERALL BATTING AVERAGE is %0.3f\n\n", Average(batters))
	
	//Print the top of the chart
	fmt.Println("    PLAYER NAME      :    AVERAGE  SLUGGING   ONBASE%")
	fmt.Println("-----------------------------------------------------")
	
	//Iterate through the batters slice and print the information, formatted to the screen
	for i := 0; i < len(batters); i++ {
		fmt.Printf("%20v :      %0.3f     %0.3f     %0.3f\n", batters[i].lastName + ", " + batters[i].firstName, batters[i].average, batters[i].slugging, batters[i].onBase)
	}
	
	//Print the number of error lines
	fmt.Printf("\n----- %d ERROR LINES FOUND IN INPUT DATA -----\n\n", len(errorlines))
	
	//Print the error lines
	for i := 0; i < len(errorlines); i++ {
		fmt.Println(errorlines[i])
	}
}
