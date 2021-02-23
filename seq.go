package main

import (
	"fmt"
	"unicode"
)

// Declaration for codon type : c := codon{'G', 'T', 'A'}
type codon [3]rune

func (c *codon) translate() rune {
	tc := map[codon]rune{
		{'G', 'C', 'U'}: 'A', {'G', 'C', 'C'}: 'A', {'G', 'C', 'A'}: 'A', {'G', 'C', 'G'}: 'A',
		{'U', 'G', 'U'}: 'C', {'U', 'G', 'C'}: 'C',
		{'G', 'A', 'U'}: 'D', {'G', 'A', 'C'}: 'D',
		{'G', 'A', 'A'}: 'E', {'G', 'A', 'G'}: 'E',
		{'U', 'U', 'U'}: 'F', {'U', 'U', 'C'}: 'F',
		{'G', 'G', 'U'}: 'G', {'G', 'G', 'C'}: 'G', {'G', 'G', 'A'}: 'G', {'G', 'G', 'G'}: 'G',
		{'C', 'A', 'U'}: 'H', {'C', 'A', 'C'}: 'H',
		{'A', 'U', 'A'}: 'I', {'A', 'U', 'U'}: 'I',
		{'A', 'U', 'C'}: 'I', {'A', 'A', 'A'}: 'K',
		{'A', 'A', 'G'}: 'K',
		{'U', 'U', 'A'}: 'L', {'U', 'U', 'G'}: 'L', {'C', 'U', 'U'}: 'L', {'C', 'U', 'C'}: 'L',
		{'C', 'U', 'A'}: 'L', {'C', 'U', 'G'}: 'L',
		{'A', 'U', 'G'}: 'M', {'A', 'A', 'U'}: 'N', {'A', 'A', 'C'}: 'N',
		{'C', 'C', 'U'}: 'P', {'C', 'C', 'C'}: 'P', {'C', 'C', 'A'}: 'P', {'C', 'C', 'G'}: 'P',
		{'C', 'A', 'A'}: 'Q', {'C', 'A', 'G'}: 'Q',
		{'C', 'G', 'U'}: 'R', {'C', 'G', 'C'}: 'R', {'C', 'G', 'A'}: 'R', {'C', 'G', 'G'}: 'R', {'A', 'G', 'A'}: 'R', {'A', 'G', 'G'}: 'R',
		{'U', 'C', 'U'}: 'S', {'U', 'C', 'C'}: 'S', {'U', 'C', 'A'}: 'S', {'U', 'C', 'G'}: 'S', {'A', 'G', 'U'}: 'S', {'A', 'G', 'C'}: 'S',
		{'A', 'C', 'U'}: 'T', {'A', 'C', 'C'}: 'T', {'A', 'C', 'A'}: 'T', {'A', 'C', 'G'}: 'T',
		{'G', 'U', 'U'}: 'V', {'G', 'U', 'C'}: 'V', {'G', 'U', 'A'}: 'V', {'G', 'U', 'G'}: 'V',
		{'U', 'G', 'G'}: 'W',
		{'U', 'A', 'U'}: 'Y', {'U', 'A', 'C'}: 'Y',
		{'U', 'A', 'A'}: '_', {'U', 'A', 'G'}: '_', {'U', 'G', 'A'}: '_',
	}

	// DEBUG Check syntax
	value, ok := tc[*c]
	fmt.Println(value, ok)
	if ok {
		return value
	}

	return '!'
}

// Entry Library -> May refactor later
type Entry struct {
	nucAcid  DNA
	rnucAcid RNA
	aaSeq    Protein
}

func newEntry(mechanism rune) *Entry {
	var title string
	var seq string

	// TODO: Get user input vs get from fasta

	fmt.Println("Enter title: ")
	fmt.Scanf("%s\n", &title)
	fmt.Println("Enter Sequence: ")
	fmt.Scanf("%s\n", &seq)

	dna := newDNA(title, seq)
	rna := transcribe(*dna)
	aa := translateToProtein(*rna)

	e := Entry{nucAcid: *dna, rnucAcid: *rna, aaSeq: *aa}
	return &e
}

// Print Details
func (e *Entry) printEntry() {

}

// DNA Section
type DNA struct {
	name            string
	sequence        []rune
	complement      []rune
	size            int
	frequency       map[rune]int // count, not percentage
	knownComplement bool         // True if user entered a known complement. Enables mutation checking. False if complement was calculated.
	subsequences    []rune       // noted subsequences, if added by user
}

func newDNA(title string, seq string) *DNA {
	d := DNA{name: title}
	d.sequence = []rune(seq)

	d.validateDna()
	d.reverseComplement('y')
	d.gcContent()

	return &d
}

//	Ensure all runes are capital, all unknowns are 'N'
func (s *DNA) validateDna() {
	s.size = len(s.sequence)
	for i := 0; i < s.size; i++ {
		if s.sequence[i] == 'a' || s.sequence[i] == 'c' || s.sequence[i] == 't' || s.sequence[i] == 'g' {
			s.sequence[i] = unicode.ToUpper(s.sequence[i])
		}
	}
}

//	choice should take in a user choice. Also set known_complement here
func (s *DNA) reverseComplement(auto rune) {
	if auto == 'y' || auto == 'Y' {
		for i := 0; i < s.size; i++ {
			if s.sequence[i] == 'A' {
				s.complement = append(s.complement, 'T')
			} else if s.sequence[i] == 'T' {
				s.complement = append(s.complement, 'A')
			} else if s.sequence[i] == 'G' {
				s.complement = append(s.complement, 'C')
			} else if s.sequence[i] == 'C' {
				s.complement = append(s.complement, 'G')
			} else {
				s.complement = append(s.complement, s.sequence[i])
			}
		}
	} else {
		var insert string
		fmt.Println("Enter the complementary strand: ")
		fmt.Scanln("Enter the complementary strand: ")
		s.complement = []rune(insert)
	}
}

// TODO
func (s *DNA) gcContent() float64 {
	return 0.0
}

// Print Ladder format
func (s *DNA) printStrands() {
	for i := 0; i < s.size; i++ {
		fmt.Printf("%c - %c\n", s.sequence[i], s.complement[i])
	}
}

// RNA Section
type RNA struct {
	name     string
	sequence []rune
	mRNA     []rune
	codons   []codon
}

// RNA constructor
func transcribe(d DNA) *RNA {
	r := RNA{name: d.name}
	var choice rune
	var start int
	var end int

	fmt.Printf("Transcribe (F)ull strand or (S)ubsequence?\n")
	for {
		fmt.Scanf("%c\n", &choice)
		// fmt.Println(choice)

		if choice == 'f' || choice == 's' {
			break
		}
	}

	// transcribe full sequence first
	for i := 0; i < d.size; i++ {
		if d.sequence[i] == 'A' {
			r.sequence = append(r.sequence, 'U')
		} else if d.sequence[i] == 'T' {
			r.sequence = append(r.sequence, 'A')
		} else if d.sequence[i] == 'G' {
			r.sequence = append(r.sequence, 'C')
		} else if d.sequence[i] == 'C' {
			r.sequence = append(r.sequence, 'G')
		} else {
			r.sequence = append(r.sequence, d.sequence[i])
		}
		// fmt.Println(r.sequence)
	}

	// then extract mRNA subsequence
	if choice == 'f' {
		r.mRNA = r.sequence

	} else if choice == 's' {
		// valid_range := false
		validRange := false
		for validRange != true {
			fmt.Println("Enter starting nucleotide: ")
			n, err := fmt.Scanf("%d\n", &start)
			if err != nil || n != 1 {
				fmt.Println(n, err)
			}

			fmt.Println("Enter ending nucleotide: ")
			n, err = fmt.Scanf("%d\n", &end)
			if err != nil || n != 1 {
				fmt.Println(n, err)
			}

			if start >= 0 && end <= d.size && start < end {
				validRange = true
			} else {
				fmt.Println("Invalid range, please re-enter")
			}
		}
		r.mRNA = r.sequence[start:end]
	}
	r.extractCodons()

	return &r
}

// needs to be rewritten
func (r *RNA) extractCodons() {
	k := 0
	insert := codon{}
	for k < len(r.mRNA) {
		insert = codon{}
		i := 0
		for i < 3 {
			// Refactor
			if k+i >= len(r.mRNA) {
				break
			}
			insert[i] = r.mRNA[k+i]
			//fmt.Println(r.mRNA[k+i])
			i++
		}
		//fmt.Println("Printing current insert")
		//fmt.Println(insert)
		r.codons = append(r.codons, insert)
		k += 3
	}
}

func (r *RNA) mrnaPrint() {
	fmt.Println("Printing mRNA: ")
	fmt.Println(r.mRNA)
	for i := 0; i < len(r.mRNA); i++ {
		fmt.Printf("%c", r.mRNA[i])
	}
}

// Protein Section (Amino Acid sequences)
type Protein struct {
	name      string
	aminoacid string
	size      int
}

func translateToProtein(m RNA) *Protein {
	p := Protein{name: m.name}
	var tempAA []rune

	for i := 0; i < len(m.codons); i++ {
		//fmt.Println(m.codons[i])
		tempResidue := m.codons[i].translate()
		//fmt.Println(tempResidue)
		tempAA = append(tempAA, tempResidue)
	}

	// remove last element if it is invalid
	if tempAA[len(tempAA)-1] == '!' {
		tempAA = tempAA[:len(tempAA)-1]
	}

	p.aminoacid = string(tempAA)
	p.size = len(p.aminoacid)
	return &p
}

func main() {
	Library := make([]Entry, 0)
	exit := 'c'
	run(&Library, exit)
	// TODO
}

func run(e *[]Entry, c rune) {
	for {
		if c == 'q' || c == 'Q' {
			break
		} else if c == 'a' || c == 'A' {
			// Add stuff into list
		} else if c == 'e' || c == 'E' {
			// Remove stuff from list
		} else if c == 'w' {
			// Write using genetext.go
		} else if c == 'r' {
			// Read using genetext.go
		}
	}
	// fmt.Println(e)
}

func addEntryToList(n *[]Entry) {
	*n = append(*n, *newEntry('1'))
}

func removeEntryFromList(n []Entry, name string) {

}

func outputList(n []Entry) {
	fmt.Println("Printng List")
}
