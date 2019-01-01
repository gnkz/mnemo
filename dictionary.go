package mnemo

type (
	// A Dictionary represents a list of words and a separator used to create a mnemonic
	Dictionary interface {
		// Returns a word at certain position. If the index does not exists this should return an error
		Word(index int) (string, error)

		// Returns a string used as separator for each word in the mnemonic
		Separator() string
	}
)
