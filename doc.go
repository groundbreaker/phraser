// Package phraser generates human readable passphrases from customizable word
// lists. These passphrases are great for use as temporary passwords or
// whenever you need a set of user-friendly random words.
package phraser

// VERSION returns the SemVer for Package phraser.
const VERSION = "v1.0.2"

// DefaultSymbol will be used when making Phrases.
var DefaultSymbol = "!"

// DefaultAdjectives is a list of adjectives that phraser will use by default.
var DefaultAdjectives = []string{"adaptable", "adventurous", "affable", "affectionate", "agreeable", "ambitious", "amiable", "amusing", "brave", "bright", "calm", "careful", "charming", "communicative", "compassionate ", "conscientious", "considerate", "convivial", "courageous", "courteous", "creative", "daring", "decisive", "determined", "diligent", "diplomatic", "discreet", "dynamic", "easygoing", "emotional", "empathetic", "energetic", "enthusiastic", "exuberant", "exuberant", "faithful", "fearless", "forceful", "friendly", "funny", "generous", "gentle", "gregarious", "helpful", "honest", "humorous", "imaginative", "impartial", "independent", "intelligent", "intuitive", "inventive", "kind", "loving", "loyal", "modest", "neat", "nice", "passionate", "patient", "persistent", "persistent ", "philosophical", "pioneering", "placid", "plucky", "polite", "powerful", "practical", "profitable", "quick", "rational", "reliable", "reserved", "resourceful", "sensible", "sensitive", "shy", "sincere", "sociable", "strong", "sympathetic", "thoughtful", "tidy", "unassuming", "understanding", "versatile", "warmhearted", "willing", "witty"}

// DefaultNouns is a list of nouns that phraser will use by default.
var DefaultNouns = []string{"animals", "artists", "baskets", "bottles", "brothers", "champs", "children", "circles", "dangers", "details", "doors", "droids", "dwarves", "engines", "events", "experts", "feelings", "freezers", "friends", "geese", "glasses", "graffiti", "groups", "harbors", "humans", "islands", "jellies", "journeys", "judges", "letters", "linens", "loaves", "machines", "minerals", "minutes", "moments", "months", "mornings", "nations", "numbers", "offspring", "opinions", "papers", "people", "pictures", "places", "planes", "plants", "purposes", "quizzes", "reasons", "rewards", "rivers", "roads", "robots", "scarves", "scissors", "shades", "sisters", "sounds", "soups", "spoons", "springs", "squares", "stamps", "stones", "students", "summers", "systems", "tables", "tails", "tastes", "teachers", "things", "thinkers", "thoughts", "times", "trains", "trucks", "values", "watches", "weeks", "whistles", "windows", "winters", "wires", "worlds", "worms", "years", "zebras"}
