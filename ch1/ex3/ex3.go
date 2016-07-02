// Package main provides ...
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo() // First call shouldn't count
	echo()
	echo2()
	echo3()
	// To check if order of execution makes any difference
	echo3()
	echo2()
	echo()
}

/*
hi there you whats up
Echo: 41.444µs
hi there you whats up
Echo: 4.142µs
hi there you whats up
Echo2: 4.136µs
hi there you whats up
Echo3: 3.39µs
hi there you whats up
Echo3: 2.716µs
hi there you whats up
Echo2: 3.88µs
hi there you whats up
Echo: 3.101µs
*/

func echo() {
	startTime := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("Echo:", time.Since(startTime))
}

func echo2() {
	startTime2 := time.Now()
	s2, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s2)
	fmt.Println("Echo2:", time.Since(startTime2))
}

func echo3() {
	startTime3 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("Echo3:", time.Since(startTime3))
}
