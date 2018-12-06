package phraser

import (
	"fmt"
	"math/rand"
	"strings"
)

// DefaultSymbol will be used when making Phrases.
var DefaultSymbol = "!"

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
