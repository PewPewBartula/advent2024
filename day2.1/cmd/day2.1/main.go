package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"flag"
)

func readFile(fname string) (data string, err error) {
    dataFileFormat, err := ioutil.ReadFile(fname)
    if err != nil { return "", err }
    data = string(dataFileFormat)
    return data, nil
}

func getArrFromFile(data string) (arr [][]int, err error){
    lines := strings.Split(data, "\n")
    arr = make([][]int, (len(lines))-1)
    for i, _:= range lines{
	if lines[i] == "" { continue }
	rows := strings.Split(lines[i], " ")
	for j, _ := range rows {
		if rows[j] == "" { 
			fmt.Println("Error line", i)
		} else {
			row, err := strconv.Atoi(rows[j])
			if err != nil { panic(err)}
			arr[i] = append(arr[i], row)
		}
	}
    } 
    return arr, nil
}

func isReportSafe(data []int, debug bool) (isSafe bool, err error) {
	inc, dec := false, false
	for i, _ := range data {
		if i == ( len(data)) -1 { break }
		result := data[i] - data[i+1]
		if debug == true { fmt.Println("Iteration", i) }
		if debug == true { fmt.Println("Difference", result) }
		if i == 0 {
			if result == 0 { 
			  	if debug == true { fmt.Println("Return: no difference") }
				return false, nil
			} else if result < 0 { 
				inc = true 
				if debug == true { fmt.Println("Increasing") }
			} else { 
				dec = true 
				if debug == true { fmt.Println("Decreasing") }
			}
		} 
		if result == 0 { 
			if debug == true { fmt.Println("Return: no difference") }	
			return false, nil 
		} else if result < 0 {
			if result < -3 || inc == false {
				if debug == true { fmt.Println("Unsafe") }
				return false, nil	
			} 
		} else { 
			if result > 3 || dec == false {
				if debug == true { fmt.Println("Unsafe") }
				return false, nil
			}	
		}
	}
	return true, nil
}

func analyzeReports(data [][]int, debug bool) (safeAmount int, err error){
	safeAmount = 0
	for i, _ := range data {
		if debug==true { fmt.Println("----------------") }
		if debug==true { fmt.Println("Report", data[i]) }
		if data[i] == nil { break }
		isReportSafe, err := isReportSafe(data[i], debug)
		if err != nil { panic(err) }
		if isReportSafe == true { safeAmount++ }
		if debug==true { fmt.Println("Current safe amount:", safeAmount) }
	}
	return safeAmount, nil	
}

func main() {
	debugPtr := flag.Bool("d", false, "debug mode")
	flag.Parse()
	filename := "data"
	data, err := readFile(filename)
	if err != nil { panic(err) }
	arr, err := getArrFromFile(data)
	if err != nil { panic(err) }
	safeAmount, err := analyzeReports(arr, *debugPtr)
	fmt.Println("Amount of safe Reports:", safeAmount)
}

