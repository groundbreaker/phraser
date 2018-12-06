package phraser_test

import (
	"math/rand"

	"github.com/groundbreaker/phraser"
)

func (suite *PhraserTestSuite) TestParseWordList() {
	suite.Equal(suite.Adjectives, phraser.ParseWordList("_test/wordlist"), "it returns the correct words")
}

func (suite *PhraserTestSuite) TestNewRNG() {
	suite.IsType(&rand.Rand{}, phraser.NewRNG(), "it returns an instance of rand.Rand")
}
