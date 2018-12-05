// Package phraser generates human readable passphrases from customizable word
// lists. These passphrases are great for use as temporary passwords or
// whenever you need a set of user-friendly random words.
package phraser

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// VERSION returns the SemVer for Package phraser.
const VERSION = "v1.0.0"

var (
	// DefaultSymbol will be used when making Phrases.
	DefaultSymbol = "!"
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

// New returns a new instance of Phraser. Requires the paths to each word
// list it needs to Adjectives and Nouns.
func New(adjPath string, nounPath string) *Phraser {
	return &Phraser{
		Adjectives: ParseWordList(adjPath),
		Nouns:      ParseWordList(nounPath),
		RNG:        rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Phraser has two word lists, one for Adjectives and one for Nouns. It uses
// these words when creating passphrases.
type Phraser struct {
	Adjectives []string
	Nouns      []string
	RNG        *rand.Rand
}

// GeneratePhrase genereates and returns a new phrase
func (phraser *Phraser) GeneratePhrase() *Phrase {
	return &Phrase{
		Adjective: phraser.RandomAdjective(),
		Number:    phraser.RNG.Intn(99),
		Noun:      phraser.RandomNoun(),
		Symbol:    DefaultSymbol,
	}
}

// RandomAdjective returns a random word from the Adjectives list.
func (phraser *Phraser) RandomAdjective() string {
	return phraser.Adjectives[phraser.RNG.Intn(len(phraser.Adjectives))]
}

// RandomNoun returns a random word from the Nouns list.
func (phraser *Phraser) RandomNoun() string {
	return phraser.Nouns[phraser.RNG.Intn(len(phraser.Nouns))]
}

// Phrase contains the phrase parts and the ability to concatentate the parts
// into a passphrase.
type Phrase struct {
	Adjective string
	Number    int
	Noun      string
	Symbol    string
}

// Passphrase contcatenates the phrase parts and returns a passphrase.
func (phrase *Phrase) Passphrase() string {
	return fmt.Sprintf("%s%v%s%s", strings.Title(phrase.Adjective), phrase.Number, strings.Title(phrase.Noun), phrase.Symbol)
}
