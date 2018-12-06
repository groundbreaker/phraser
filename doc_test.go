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

func ExampleParseWordList() {
	phraser.ParseWordList("/path/to/words")
	// Output: []string{"people", "humans","robots", "droids"}
}
