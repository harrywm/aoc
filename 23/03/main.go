package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Raw input
var raw []string

func main(){
    fmt.Println("Day 03")
    raw = loadInput()

    // Raw is split by NewLines
    for _, line := range raw {
        fmt.Println(line)
    }
}

// Read from input.txt
func loadInput() []string {
    var in []string

    file, err := os.Open("input.txt")
    if err != nil {log.Fatal(err)}
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {in = append(in, scanner.Text())}
    if err := scanner.Err(); err != nil {log.Fatal(err)}

    defer file.Close()
    return in
}



