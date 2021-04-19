package automata

// Class BoyerMoore and Associated Functions:
type BoyerMoore struct {
	alphabet  []rune
	pattern   []rune
	f         []int
	s         []int
	occurence map[rune]int
}

// Add a constructor tbd

// Preprocessing - called by constructor
func (bm *BoyerMoore) preprocess() {
	bm.processBCR()
	bm.processGSR()
}

// BCR processing - called by constructor
func (bm *BoyerMoore) processBCR() {
	bm.occurence = map[rune]int{}

	// parse through runes in alphabet
}

// GSR processing - called by constructor
func (bm *BoyerMoore) processGSR() {

}

// BoyerMoore Algorithm
func (bm BoyerMoore) SearchPattern(text string) []rune {
	var result []rune
	return result
}

// Class Automata and Associated Functions:
type Automata struct {
	numstates       int
	alphabet        rune
	transitionTable map[rune]rune // Check/Adjust map type

}

func AutomataConstructor() *Automata {
	return &Automata{numstates: 0}
}

// Pattern Search Functions (No Struct):
// Returns the position of the first occurence of given pattern. Additionally prints results vertically.
func SearchFirstOccurance(seq []rune, pattern string) int {
	return -1
}

// Return a list with all initial positions of the pattern's occurences. Additionally prints results vertically.
func SearchAllOccurences(seq []rune, pattern string) []int {
	var result []int
	return result
}

// Allows user to test pattern occurence without visual display. Void return.
func TestPatternSearch() {
}
