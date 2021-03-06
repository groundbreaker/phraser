# :speech_balloon: phraser

[![GoDoc](https://godoc.org/github.com/groundbreaker/phraser?status.svg)](https://godoc.org/github.com/groundbreaker/phraser)

Phraser generates human readable passphrases from customizable word lists.

These passphrases are great for use as temporary passwords or whenever you need
a set of user-friendly random words strung together.

## Install

    go get github.com/groundbreaker/phraser

## Usage

```go
import (
  "fmt"

  "github.com/groundbreaker/phraser"
)

func main() {
  p := phraser.New()
  phrase := p.GeneratePhrase()
  fmt.Println(phrase.Passphrase())
}
```
