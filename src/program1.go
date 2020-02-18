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
	"unsafe"
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
	
	batters := ParseInfo(data)
	
	batters = batters
	
	//fmt.Println(data)
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

func ParseInfo(data string) []BatterInfo {
	batters := []BatterInfo { }
	
	data = strings.Replace(data, "\r", "", -1)
	lines := strings.Split(data, "\n")
	
PrimaryLoop:
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
		
		var batter BatterInfo
		
		tokens := []string { }
		spaceSeparatedValues := strings.Split(lines[i], " ")
		for j := 0; j < len(spaceSeparatedValues); j++ {
			if spaceSeparatedValues[j] != "" {
				tokens = append(tokens, spaceSeparatedValues[j])
			}
		}
		
		if len(tokens) != 10 {
			fmt.Println("Invalid line entered-- incorrect number of parameters.")
			continue
		}
		
		var err error
		
		batter.firstName = tokens[0]
		batter.lastName = tokens[1]
		for j := 0; j < 8; j++ {
			size := unsafe.Sizeof(uint64(0))
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(&batter.plateAppearances)) + size * uintptr(j))), err = strconv.ParseUint(tokens[j + 2], 10, 32) //all i can say about this is that i'm sorry
			if err != nil {
				fmt.Println("Invalid line entered-- illegal type of parameter.")
				continue PrimaryLoop
			}
		}
		
		//batter.plateAppearances, err = strconv.ParseUint(tokens[2], 10, 32)
		//if err != nil {
		//	fmt.Println("Invalid line entered-- illegal type of parameter.")
		//	continue
		//}
		//batter.atBats, err = strconv.ParseUint(tokens[3], 10, 32)
		//if err != nil {
		//	fmt.Println("Invalid line entered-- illegal type of parameter.")
		//	continue
		//}
		//batter.singles, err = strconv.ParseUint(tokens[4], 10, 32)
		//if err != nil {
		//	fmt.Println("Invalid line entered-- illegal type of parameter.")
		//	continue
		//}
		//batter.doubles, err = strconv.ParseUint(tokens[5], 10, 32)
		//if err != nil {
		//	fmt.Println("Invalid line entered-- illegal type of parameter.")
		//	continue
		//}
		//batter.triples, err = strconv.ParseUint(tokens[6], 10, 32)
		//if err != nil {
		//	fmt.Println("Invalid line entered-- illegal type of parameter.")
		//	continue
		//}
		//batter.homeRuns, err = strconv.ParseUint(tokens[7], 10, 32)
		//if err != nil {
		//	fmt.Println("Invalid line entered-- illegal type of parameter.")
		//	continue
		//}
		//batter.walks, err = strconv.ParseUint(tokens[8], 10, 32)
		//if err != nil {
		//	fmt.Println("Invalid line entered-- illegal type of parameter.")
		//	continue
		//}
		//batter.hitByPitch, err = strconv.ParseUint(tokens[9], 10, 32)
		//if err != nil {
		//	fmt.Println("Invalid line entered-- illegal type of parameter.")
		//	continue
		//}
		
		batters = append(batters, batter)
	}
	
	return batters
}