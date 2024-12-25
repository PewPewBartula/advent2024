package main

import (
	"fmt"
    "io/ioutil"
    "strconv"
    "strings"
    "sort"
    "math"
)

func readFile(fname string) (data string, err error) {
    dataFileFormat, err := ioutil.ReadFile(fname)
    if err != nil { return "", err }
    data = string(dataFileFormat)
    return data, nil
}

func getArrFromFile(data string) (arr [][]int, err error){

    lines := strings.Split(data, "\n")
    arr = make([][]int, len(lines))
    for i, _:= range lines{
	if lines[i] == "" { continue }
	rows := strings.Split(lines[i], "   ")
	for j, _ := range rows {
		if rows[j] == "" { 
			fmt.Println("Error line ", i)
		} else {
			row, err := strconv.Atoi(rows[j])
			if err != nil { panic(err)}
			arr[i] = append(arr[i], row)
		}
	}
	rows = nil
    } 

    return arr, nil
}

func getLowestNumbers(arr [][]int, index int)(result []int, err error){
	if index > (len(arr)-1) { panic(err) }
	left := make([]int, len(arr))
	right := make([]int, len(arr))
	for i, _:= range arr {
		if arr[i] == nil { continue }
		left = append(left, arr[i][0])
		right = append(right, arr[i][1])
	}
	sort.Slice(left, func(i, j int) bool {
    		return left[i] > left[j]
	})
	sort.Slice(right, func(i, j int) bool {
    		return right[i] > right[j]
	})
	return []int{ left[0+index], right[0+index] }, nil
}

func getDifference(arr []int)(result int, err error){
	return int(math.Abs((float64(arr[0]))-(float64(arr[1])))), nil	
}

func main() {
	var filename = "data"
	//var filename = "test"
	var difference = 0
	data, err := readFile(filename)
    	if err != nil { panic(err) }
	arr, err := getArrFromFile(data)
	if err != nil { panic(err) }
	//iterations := int(len(arr))
	//diffArr := make([]int, len(arr))
	var result = 0
	for i, _ := range arr {
		lowest, err := getLowestNumbers(arr,i)
		if err != nil { panic(err) }
		difference, err = getDifference(lowest)
		if err != nil { panic(err) }
		result = result + difference
	}
	fmt.Println("Result is ", result)
}
