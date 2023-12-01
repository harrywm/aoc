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

var digitizedValues []string
var calibrationValues []int64 

// Day 01
// Sum all calibration values.
// input.txt contains muddled calibration values, one per line, where a calibration value is the first digit and last on each line.
func main() {
    fmt.Println("Day 01")
    
    rawValues := loadValues()


    // Digitize rawValues
    for _, rawValue := range rawValues {
        digitizedValues = append(digitizedValues, digitize(rawValue))
    }

    for _, value := range digitizedValues {
        calibrationValues = append(calibrationValues, unMuddle(value))
    }

    // sum Calibration Values
    sum := sumCalibrationValues(calibrationValues)
    fmt.Println(sum)
}

// Read from input.txt
func loadValues() []string {
    fmt.Println("Loading values...")
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    var rawValues []string
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        rawValues = append(rawValues, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    defer file.Close()
    return rawValues
}

func digitize(line string) string {
    numbers := map[string]int64{
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    var digitized []string
    for i := 0; i < len(line); i++ {
        idxPlus := string(line[i:4])
        if _, err := strconv.ParseInt(string(line[i]), 10, 64); err == nil {
            digitized = append(digitized, string(line[i]))
       }
       for num, digit := range numbers { 
           if idxPlus == num {
               digitized = append(digitized, strconv.FormatInt(digit, 10))
           }
       }
   }

    // split line by match in numbers
    returnDigitized := strings.Join(digitized, "")
    // Return string
    return returnDigitized
}

func unMuddle(v string) int64 {

    re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
    i := re.FindAllString(v, -1)
    
    squashed := strings.Join(i, "")

    var candidateDigits []string = []string{squashed[:1], squashed[len(squashed)-1:]}

    squashedDigits := strings.Join(candidateDigits, "")

    intDigits, err := strconv.ParseInt(squashedDigits, 10, 64)
    if err != nil {
        log.Fatal(err)
    }

    return intDigits
}

func sumCalibrationValues(c []int64) int64 {
    var sum int64 = 0
    for _, v := range c {
        sum += v
    }
    return sum
}
