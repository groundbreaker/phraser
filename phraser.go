package phraser

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// New returns a new instance of Phraser using the DefaultAdjectives and
// DefaultNouns. You can override these exported vars before calling New() to
// use your own word list.
func New() *Phraser {
	return &Phraser{
		Adjectives: DefaultAdjectives,
		Nouns:      DefaultNouns,
		RNG:        NewRNG(),
	}
}

// Phraser has two word lists, one for Adjectives and one for Nouns. It uses
// these words when creating passphrases.
type Phraser struct {
	Adjectives []string
	Nouns      []string
	RNG        *rand.Rand
}

// GeneratePhrase genereates and returns a new Phrase.
func (phraser *Phraser) GeneratePhrase() *Phrase {
	return &Phrase{
		Adjective: phraser.RandomWord(phraser.Adjectives),
		Number:    phraser.RNG.Intn(99),
		Noun:      phraser.RandomWord(phraser.Nouns),
		Symbol:    DefaultSymbol,
	}
}

// RandomWord returns a random word from the given list.
func (phraser *Phraser) RandomWord(words []string) string {
	return words[phraser.RNG.Intn(len(words))]
}

// Phrase contains the phrase parts and the ability to concatentate the parts
// into a passphrase.
type Phrase struct {
	Adjective string
	Number    int
	Noun      string
	Symbol    string
}

// Passphrase contcatenates the phrase parts and returns a passphrase in the
// form of: AdjectiveNumberNounSymbol, e.g. "Amiable42Robots!".
func (phrase *Phrase) Passphrase() string {
	return fmt.Sprintf("%s%v%s%s", strings.Title(phrase.Adjective), phrase.Number, strings.Title(phrase.Noun), phrase.Symbol)
}

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
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
