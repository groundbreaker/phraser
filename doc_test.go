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
