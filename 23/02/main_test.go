package main

import (
	"testing"
)

func TestIterate(t *testing.T) {
    var testInput = []string{
            "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
            "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
            "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
            "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
            "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
        }
    var testOutput = []int64{
            48,
            12,
            1560,
            630,
            36,
    }

    out := iterateGames(testInput)
    for idx, v := range out {
        if v != testOutput[idx] {
            t.Errorf("iterateGames(%q) = %d", testInput[idx], testOutput[idx])
        }
    }
}

func BenchmarkIterate(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var testInput = []string{
            "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
            "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
            "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
            "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
            "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
        }
        iterateGames(testInput)
    }
}
