package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Raw input
var rawValues []string

// Answer Outputs
var possibleIDs []int64
var powerOfMins []int64

// Map to compare Cube actuals against, the "rules"
var cubeRules = map[string]int64{
    "Red": 12,
    "Green": 13,
    "Blue": 14,
}

func main(){
    fmt.Println("Day 02")

    powerOfMins = iterateGames(loadValues())

    fmt.Println("Powers: ", sumValues(powerOfMins))
    fmt.Println("Sum: ", sumValues(possibleIDs))
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

func iterateGames(games []string) []int64 {

    var minPowerList []int64
    
    for idx, game := range games {
    
        possibleGame := true
        minRolls := []int64{0, 0, 0}
        
        gameWithoutID := strings.Split(game, ":")
        gameRounds := strings.Split(gameWithoutID[1], ";")

        for _, round := range gameRounds {
            for _, roll := range strings.Split(round, ",") {

                re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
                i := re.FindAllString(roll, -1)
                
                rollNum, err := strconv.ParseInt(strings.Join(i, ""), 10, 64)
                if err != nil {log.Fatal(err)} 
            
                if strings.Contains(roll, "red") {
                    if rollNum > cubeRules["Red"] {     
                        possibleGame = false
                    }

                    if rollNum > minRolls[0] {
                        minRolls[0] = rollNum
                    }
                }
                
                if strings.Contains(roll, "green") {
                    if rollNum > cubeRules["Green"] {     
                        possibleGame = false
                    }
                    if rollNum > minRolls[1] {
                        minRolls[1] = rollNum
                    }
                }
                
                if strings.Contains(roll, "blue") {
                    if rollNum > cubeRules["Blue"] {     
                        possibleGame = false
                    }
                    if rollNum > minRolls[2] {
                        minRolls[2] = rollNum
                    }
                }
            }
        }
        minPowerList = append(minPowerList, multiplyValues(minRolls))
        if possibleGame { possibleIDs = append(possibleIDs, int64(idx + 1)) }
    }
    return minPowerList
}

func sumValues(c []int64) int64 {
    var sum int64 = 0
    for _, v := range c {
        sum += v
    }
    return sum
}

func multiplyValues(c []int64) int64 {
    out := c[0] * c[1] * c[2]
    return out 
}
