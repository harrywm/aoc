package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Raw input
var rawValues []string

func main(){
    fmt.Println("Day 03")
    rawValues = loadValues()
}

// Read from input.txt
func loadValues() []string {
    
    file, err := os.Open("input.txt")
    if err != nil {log.Fatal(err)}
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {rawValues = append(rawValues, scanner.Text())}
    if err := scanner.Err(); err != nil {log.Fatal(err)}

    defer file.Close()
    return rawValues
}



