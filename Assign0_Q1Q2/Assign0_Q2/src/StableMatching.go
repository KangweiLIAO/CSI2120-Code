package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var (
	empMap     map[string][]string
	stuMap     map[string][]string
	ePairs     map[string]string
	sPairs     map[string]string
	wg         sync.WaitGroup
	numOfPairs int
	done       chan bool
)

func readTable(path string) *[][]string {
	tFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%v Not Found!\n", path)
		panic(err)
	}
	tReader := csv.NewReader(strings.NewReader(string(tFile)))
	tArr, _ := tReader.ReadAll()
	return &tArr
}

func initAlgorithm(t1Arr *[][]string, t2Arr *[][]string) {
	empMap = make(map[string][]string)
	stuMap = make(map[string][]string)
	for i := 0; i < len(*t1Arr); i++ {
		empMap[(*t1Arr)[i][0]] = (*t1Arr)[i][1:]
		stuMap[(*t2Arr)[i][0]] = (*t2Arr)[i][1:]
	}
}

func offer(currE string) {
	defer wg.Done()
	if ePairs[currE] == "" {
		for _, currS := range empMap[currE] {
			if evaluate(currE, currS) == true {
				return
			}
		}
	}
}

func evaluate(emp string, stu string) bool {
	<-done
	if sPairs[stu] == "" {
		// fmt.Println(emp + ": " + stu)
		// set employer-student pair
		ePairs[emp] = stu
		sPairs[stu] = emp
		done <- true
		return true
	} else {
		prevEmp := sPairs[stu]
		buffer := stuMap[stu]
		// check if student prefer current employer or not
		if sliceIndexOf(emp, &buffer) < sliceIndexOf(prevEmp, &buffer) {
			// fmt.Println(stu + ": " + prevEmp + " <- " + emp)
			// switch previous pair to current
			sPairs[ePairs[emp]] = ""
			ePairs[emp] = stu
			sPairs[stu] = emp
			ePairs[prevEmp] = ""
			fmt.Println(sPairs[stu])
			// call offer(previous employer)
			wg.Add(1)
			done <- true
			offer(prevEmp)
			return true
		}
		// student reject employer
		// fmt.Println(stu + " reject " + emp)
		done <- true
		return false
	}
}

func MWAlgorithm() {
	ePairs = make(map[string]string)
	sPairs = make(map[string]string)
	done = make(chan bool, 1)
	for e := range empMap {
		ePairs[e] = ""
	}
	for s := range stuMap {
		sPairs[s] = ""
	}
	fmt.Println(empMap)
	fmt.Println(stuMap)
	done <- true
	for e := range ePairs {
		wg.Add(1)
		go offer(e)
	}
	wg.Wait()
}

func sliceIndexOf(val string, slice *[]string) int {
	for i, v := range *slice {
		if v == val {
			return i
		}
	}
	return -1
}

func writeOutputTable() {
	var output [][]string
	numOfPairs = len(empMap)
	fileName := "matches_go_" + strconv.Itoa(numOfPairs) + "x" + strconv.Itoa(numOfPairs) + ".csv"
	file, err := os.Create(filepath.Dir(os.Args[0]) + "/" + fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println(ePairs)
	fmt.Println(sPairs)

	tWriter := csv.NewWriter(file)
	for key, value := range ePairs {
		output = append(output, []string{key, value})
	}
	tWriter.WriteAll(output)
	tWriter.Flush()
	fmt.Println("Successfully created output table!")
}

func main() {
	t1Arr := readTable(filepath.Dir(os.Args[0]) + "/" + os.Args[1])
	t2Arr := readTable(filepath.Dir(os.Args[0]) + "/" + os.Args[2])
	initAlgorithm(t1Arr, t2Arr)
	MWAlgorithm()
	writeOutputTable()
}
