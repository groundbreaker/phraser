package phraser

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PhraserTestSuite struct {
	suite.Suite
	Adjectives []string
	Nouns      []string
	Phraser    *Phraser
	Phrase     *Phrase
}

func (suite *PhraserTestSuite) SetupTest() {
	suite.Adjectives = []string{"adaptable", "amiable", "empathetic"}
	suite.Nouns = []string{"people", "humans", "robots"}
	suite.Phraser = New("words/test_adjectives", "words/test_nouns")
	suite.Phrase = &Phrase{
		Adjective: "amiable",
		Noun:      "robots",
		Number:    42,
		Symbol:    "?",
	}
}

func (suite *PhraserTestSuite) TestParseWordList() {
	suite.Equal(suite.Adjectives, ParseWordList("words/test_adjectives"), "it returns the correct words")
}

func (suite *PhraserTestSuite) TestNew() {
	suite.IsType(&Phraser{}, suite.Phraser, "it returns an instance of Phraser")
	suite.IsType(&rand.Rand{}, suite.Phraser.RNG, "it returns an instance of rand.Rand")
	suite.Equal(suite.Adjectives, suite.Phraser.Adjectives, "it returns the correct Adjectives")
	suite.Equal(suite.Nouns, suite.Phraser.Nouns, "it returns the correct Nouns")
}

func (suite *PhraserTestSuite) TestGeneratePhrase() {
	phrase := suite.Phraser.GeneratePhrase()
	suite.IsType(&Phrase{}, phrase, "it returns an instance of Phrase")
	suite.Equal(DefaultSymbol, phrase.Symbol, "it returns the correct Symbol")
}

func (suite *PhraserTestSuite) TestRandomAdjective() {
	suite.Contains(suite.Adjectives, suite.Phraser.RandomAdjective(), "it returns one of the adjectives on the list")
}

func (suite *PhraserTestSuite) TestRandomNoun() {
	suite.Contains(suite.Nouns, suite.Phraser.RandomNoun(), "it returns one of the nouns on the list")
}

func (suite *PhraserTestSuite) TestPassphrase() {
	suite.Equal("Amiable42Robots?", suite.Phrase.Passphrase(), "it returns the correct passphrase")
}

func TestPhraserTestSuite(t *testing.T) {
	suite.Run(t, new(PhraserTestSuite))
}
