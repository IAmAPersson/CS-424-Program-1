/*
	Phil Lane
	02/17/2020
	CS-424-02
	Mary Allen

	Programming Assignment 1 (Golang)
	
	This program takes in a file containing baseball player information, and outputs to the screen calculated data about each player, such as batting average and on base percentage.

	I have opted to program in a traditional structured paradigm. Hence, the paradigm used in this program is very similar to what you would see in a C program.
	Classes are unused-- I stick to global functions and structs.
*/

package main

//Function: main, entrypoint
func main() {
	//Get the path of the file to read from
	path := GetPath()

	//Read all the data from the file
	data, err := ReadInFile(path)
	
	//If there has been an error in reading from the file, exit program
	if err != nil {
		return
	}
	
	//Parse file and sort info
	batters, badlines := ParseInfo(data)
	batters = PlayerSort(batters)
	
	//Calculate data about batters
	calcData := Calculate(batters)
	
	//Output calculated data to the screen
	FormatData(calcData, badlines)
}
