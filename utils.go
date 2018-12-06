package phraser

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// ParseWordList reads the words from the given path to a word list file. It
// returns a slice containing the words found in that file.
//
// A word list should just be a plain text file with one word per line.
func ParseWordList(path string) []string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	words := strings.Split(string(bytes), "\n")
	return words[:len(words)-1]
}

// NewRNG returns a newly seeded random number generator.
func NewRNG() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().Unix()))
}
