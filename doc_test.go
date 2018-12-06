package phraser_test

import (
	"fmt"

	"github.com/groundbreaker/phraser"
)

func Example() {
	p := phraser.New()
	phrase := p.GeneratePhrase()
	fmt.Println(phrase.Passphrase())
}

func Example_override() {
	// Override the Default* vars before calling phraser.New()
	// Use your own []string values, or use phraser.ParseWordList to laod it from a file.
	phraser.DefaultSymbol = "#"
	phraser.DefaultAdjectives = []string{"awesome"}
	phraser.DefaultNouns = []string{"robots"}

	p := phraser.New()
	// You can also override values on the instance of phraser.
	p.Adjectives = []string{"amiable"}

	phrase := p.GeneratePhrase() // this sets the random values

	// You can override the choosen values here as well.
	phrase.Number = 42
	phrase.Symbol = "?"

	fmt.Println(phrase.Passphrase())
	// Output: Amiable42Robots?
}
