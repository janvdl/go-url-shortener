package main

import (
	"math/rand/v2"
	"os"
	"strings"
)

var adjectives []string
var animals []string

func init() {
	// load adjectives and populate in memory
	data_adjectives, err := os.ReadFile("./words/adjectives.txt")
	if err != nil {
		panic(err)
	}
	adjectives = strings.Split(string(data_adjectives), "\n")

	// load animal names and populate in memory
	data_animals, err := os.ReadFile("./words/animals.txt")
	if err != nil {
		panic(err)
	}
	animals = strings.Split(string(data_animals), "\n")
}

func makeShortLink() string {
	return strings.Join([]string{getAdjective(), getAdjective(), getAnimal()}, "-")
}

// return random adjective
func getAdjective() string {
	max_len := len(adjectives)
	r := rand.IntN(max_len)

	return adjectives[r]
}

// return random animal
func getAnimal() string {
	max_len := len(animals)
	r := rand.IntN(max_len)

	return animals[r]
}
