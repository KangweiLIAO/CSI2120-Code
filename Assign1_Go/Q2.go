// @Title  			Q2.go
// @Description 	A program that simulates a Neural Network where each neuron runs concurrently.
// @Author  		Kangwei Liao (8568800)

package main

import (
	"fmt"
	"math"
	"sync"
)

// @Title  			calcX1()
// @Description 	function for X1, pour 3 return values in to channel for Z1, Z2, Z3
// @param			int, int, chan float64
func calcX1(k int, n int, ch chan float64) {
	for i := 0; i < 3; i++ {
		ch <- math.Sin((2.0 * math.Pi * (float64(k) - 1.0)) / float64(n))
	}
}

// @Title  			calcX2()
// @Description 	function for X2, pour 3 return values in to channel for Z1, Z2, Z3
// @param			int, int, chan float64
func calcX2(k int, n int, ch chan float64) {
	for i := 0; i < 3; i++ {
		ch <- math.Cos((2.0 * math.Pi * (float64(k) - 1.0)) / float64(n))
	}
}

// @Title  			calcZ1()
// @Description 	function for Z1, pour return value in to channel for T1
// @param			chan float64, chan float64, chan float64
func calcZ1(chX1 chan float64, chX2 chan float64, chZ1 chan float64) {
	chZ1 <- 1 / (1 + math.Pow(math.E, -(0.1+0.3*<-chX1+0.4*<-chX2)))
}

// @Title  			calcZ2()
// @Description 	function for Z2, pour return value in to channel for T1
// @param			chan float64, chan float64, chan float64
func calcZ2(chX1 chan float64, chX2 chan float64, chZ2 chan float64) {
	chZ2 <- 1 / (1 + math.Pow(math.E, -(0.5+0.8*<-chX1+0.3*<-chX2)))
}

// @Title  			calcZ3()
// @Description 	function for Z3, pour return value in to channel for T1
// @param			chan float64, chan float64, chan float64
func calcZ3(chX1 chan float64, chX2 chan float64, chZ3 chan float64) {
	chZ3 <- 1 / (1 + math.Pow(math.E, -(0.7+0.6*<-chX1+0.6*<-chX2)))
}

// @Title  			calcT1()
// @Description 	function for T1, print the value
// @param			chan float64, chan float64, chan float64, *sync.WaitGroup
func calcT1(chZ1 chan float64, chZ2 chan float64, chZ3 chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("T1 = %v\n", 1/(1+math.Pow(math.E, -(0.5+0.3*<-chZ1+0.7*<-chZ2+0.1*<-chZ3))))
}

// ---------- Main Function ----------

func main() {
	var n int
	var wg sync.WaitGroup
	fmt.Printf("Please enter the input N: ")
	fmt.Scanf("%d", &n)
	chX1 := make(chan float64)
	chX2 := make(chan float64)
	chZ1 := make(chan float64)
	chZ2 := make(chan float64)
	chZ3 := make(chan float64)
	for k := 0; k < n; k++ {
		wg.Add(1)
		go calcX1(k, n, chX1)
		go calcX2(k, n, chX2)
		go calcZ1(chX1, chX2, chZ1)
		go calcZ2(chX1, chX2, chZ2)
		go calcZ3(chX1, chX2, chZ3)
		go calcT1(chZ1, chZ2, chZ3, &wg)
		wg.Wait()
	}
}
