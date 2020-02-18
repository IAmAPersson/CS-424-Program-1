/*
	Phil Lane
	02/17/2020
	CS-424-02
	Mary Allen

	Programming Assignment 1 (Golang)
*/

package main

func main() {
	path := GetPath()

	data, err := ReadInFile(path)
	
	if err != nil {
		return
	}
	
	batters, badlines := ParseInfo(data)
	batters = PlayerSort(batters)
	
	calcData := Calculate(batters)
	
	FormatData(calcData, badlines)

}