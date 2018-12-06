package phraser_test

import (
	"math/rand"
	"testing"

	"github.com/groundbreaker/phraser"
	"github.com/stretchr/testify/suite"
)

type PhraserTestSuite struct {
	suite.Suite
	Adjectives []string
	Nouns      []string
	Phraser    *phraser.Phraser
	Phrase     *phraser.Phrase
}

func (suite *PhraserTestSuite) SetupTest() {
	suite.Adjectives = []string{"adaptable", "amiable", "empathetic"}
	suite.Nouns = []string{"people", "humans", "robots"}
	suite.Phraser = phraser.New()
	suite.Phrase = &phraser.Phrase{
		Adjective: "amiable",
		Noun:      "robots",
		Number:    42,
		Symbol:    "?",
	}
}

func (suite *PhraserTestSuite) TestNew() {
	suite.IsType(&phraser.Phraser{}, suite.Phraser, "it returns an instance of Phraser")
	suite.IsType(&rand.Rand{}, suite.Phraser.RNG, "it returns an instance of rand.Rand")
	suite.Equal(phraser.DefaultAdjectives, suite.Phraser.Adjectives, "it returns the correct Adjectives")
	suite.Equal(phraser.DefaultNouns, suite.Phraser.Nouns, "it returns the correct Nouns")
}

func (suite *PhraserTestSuite) TestGeneratePhrase() {
	phrase := suite.Phraser.GeneratePhrase()
	suite.IsType(&phraser.Phrase{}, phrase, "it returns an instance of Phrase")
	suite.Equal(phraser.DefaultSymbol, phrase.Symbol, "it returns the correct Symbol")
}

func (suite *PhraserTestSuite) TestRandomWord() {
	suite.Contains(suite.Adjectives, suite.Phraser.RandomWord(suite.Adjectives), "it returns one of the adjectives on the list")
}

func (suite *PhraserTestSuite) TestPassphrase() {
	suite.Equal("Amiable42Robots?", suite.Phrase.Passphrase(), "it returns the correct passphrase")
}

func TestPhraserTestSuite(t *testing.T) {
	suite.Run(t, new(PhraserTestSuite))
}
