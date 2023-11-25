package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	possibleCharacters := "ABCDEFGHJKLMNPQRSTUVWXYZ0123456789"
	//lengthOfSet := 1

	// get length of set from user
	fmt.Printf("Enter length of set: ")
	var lengthOfSet int
	fmt.Scanln(&lengthOfSet)

	// Calculate the total number of characters
	numCharacters := len(possibleCharacters)

	// Calculate number of combinations
	numCombinations := int(math.Pow(float64(numCharacters), float64(lengthOfSet)))

	fmt.Printf("Number of Combinations: %d\n", numCombinations)

	<-time.After(5 * time.Second)

	lap := 0
	// Iterate through all combinations
	for i := 0; i < numCombinations; i++ {
		lap++
		combination := make([]rune, lengthOfSet)
		index := i

		// Generate a combination by selecting characters at different positions
		for j := 0; j < lengthOfSet; j++ {
			charIndex := index % numCharacters
			combination[lengthOfSet-j-1] = rune(possibleCharacters[charIndex]) // Fill the slice backwards
			index /= numCharacters
		}

		fmt.Printf("%s: %d\n", string(combination), lap)
	}
}
