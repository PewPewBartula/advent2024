package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
    "sort"
    "slices"
)

type NumCount struct{
	number int
	count int
}

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
	rows := strings.Split(lines[i], "   ")
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

func getNumberCount(arr [][]int)(result []NumCount, err error){
	left := make([]int, len(arr))
	right := make([]int, len(arr))
	rightCopy := make([]int, len(arr))
	var numCounts []NumCount 
	for i, _:= range arr {
		if arr[i] == nil { continue }
		left[i] = arr[i][0]
		right[i] = arr[i][1]
	}
	sort.Slice(left, func(i, j int) bool {
    		return left[i] > left[j]
	})
	sort.Slice(right, func(i, j int) bool {
    		return right[i] > right[j]
	})
	for i, _ := range left {
		copy(rightCopy, right)
		var numCount NumCount
		numCount.number = left[i]
		numCount.count = 0
		var idx = 0
		for idx != -1 {
			idx = slices.IndexFunc(rightCopy, func(n int) bool {
				return n == left[i]
			})
			if idx != -1 {
				rightCopy[idx] = -1
				numCount.count += 1
			}
		}
		if numCount.count > 0 {
			numCounts = append(numCounts, numCount)
			copy(rightCopy, right)
		}
	}
	return numCounts, nil
}

func main() {
	var filename = "data"
	//var filename = "test"
	data, err := readFile(filename)
    	if err != nil { panic(err) }
	arr, err := getArrFromFile(data)
	if err != nil { panic(err) }
	numCount, err := getNumberCount(arr)
	var result, amount = 0, 0
	for i, _ := range numCount {
		amount = numCount[i].number * numCount[i].count 
		result += amount
	}
	fmt.Println("Result is", result)
}
