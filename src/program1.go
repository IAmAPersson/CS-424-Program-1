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
	//"strings"
)

type BatterInfo struct {
	firstName string
	lastName string
	singles uint
	doubles uint
	triples uint
	homeruns uint
	atbats uint
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
	
	fmt.Println(string(data))
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