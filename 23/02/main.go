package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Raw input
var rawValues []string

// Map to load each games Cube Actuals into
var cubeActuals map[string]int

// Map to compare Cube actuals against, the "rules"
var cubeRules = map[string]int{
    "Red": 12,
    "Green": 13,
    "Blue": 14,
}

func main(){
    fmt.Println("Day 02")

    rawValues = loadValues()
    iterateGames(rawValues)
}

// Read from input.txt
func loadValues() []string {
    var rawValues []string
    
    file, err := os.Open("input.txt")
    if err != nil {log.Fatal(err)}
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {rawValues = append(rawValues, scanner.Text())}
    if err := scanner.Err(); err != nil {log.Fatal(err)}

    defer file.Close()
    return rawValues
}

// Iterate through each game, pulling out the Rounds and Game ID
// TODO: STORE IN AN OBJECT
func iterateGames(games []string) {
    for idx, game := range games {
        fmt.Println(idx + 1, game)
        gameWithoutID := strings.Split(game, ":")
        gameRounds := strings.Split(gameWithoutID[1], ";")
        fmt.Println("Rounds: ")
        for _, round := range gameRounds {
            fmt.Println(round)
        }
    }
}
