package main

import (
	"fmt"
	"unicode"
)

// Declaration for codon type : c := codon{'G', 'T', 'A'}
type codon [3]rune

func (c *codon) translate() rune {
	tc := map[codon]rune{
		{'G', 'C', 'T'}: 'A', {'G', 'C', 'C'}: 'A', {'G', 'C', 'A'}: 'A', {'G', 'C', 'G'}: 'A',

		{'T', 'G', 'T'}: 'C', {'T', 'G', 'C'}: 'C',

		{'G', 'A', 'T'}: 'D', {'G', 'A', 'C'}: 'D',

		{'G', 'A', 'A'}: 'E', {'G', 'A', 'G'}: 'E',

		{'T', 'T', 'T'}: 'F', {'T', 'T', 'C'}: 'F',

		{'G', 'G', 'T'}: 'G', {'G', 'G', 'C'}: 'G', {'G', 'G', 'A'}: 'G', {'G', 'G', 'G'}: 'G',

		{'C', 'A', 'T'}: 'H', {'C', 'A', 'C'}: 'H',

		{'A', 'T', 'A'}: 'I', {'A', 'T', 'T'}: 'I',
		{'A', 'T', 'C'}: 'I', {'A', 'A', 'A'}: 'K',

		{'A', 'A', 'G'}: 'K',

		{'T', 'T', 'A'}: 'L', {'T', 'T', 'G'}: 'L', {'C', 'T', 'T'}: 'L', {'C', 'T', 'C'}: 'L',
		{'C', 'T', 'A'}: 'L', {'C', 'T', 'G'}: 'L',

		{'A', 'T', 'G'}: 'M', {'A', 'A', 'T'}: 'N', {'A', 'A', 'C'}: 'N',
		{'C', 'C', 'T'}: 'P', {'C', 'C', 'C'}: 'P', {'C', 'C', 'A'}: 'P', {'C', 'C', 'G'}: 'P',
		{'C', 'A', 'A'}: 'Q', {'C', 'A', 'G'}: 'Q',
		{'C', 'G', 'T'}: 'R', {'C', 'G', 'C'}: 'R', {'C', 'G', 'A'}: 'R', {'C', 'G', 'G'}: 'R', {'A', 'G', 'A'}: 'R', {'A', 'G', 'G'}: 'R',
		{'T', 'C', 'T'}: 'S', {'T', 'C', 'C'}: 'S', {'T', 'C', 'A'}: 'S', {'T', 'C', 'G'}: 'S', {'A', 'G', 'T'}: 'S', {'A', 'G', 'C'}: 'S',

		{'A', 'C', 'T'}: 'T', {'A', 'C', 'C'}: 'T', {'A', 'C', 'A'}: 'T', {'A', 'C', 'G'}: 'T',

		{'G', 'T', 'T'}: 'V', {'G', 'T', 'C'}: 'V', {'G', 'T', 'A'}: 'V', {'G', 'T', 'G'}: 'V',
		{'T', 'G', 'G'}: 'W',
		{'T', 'A', 'T'}: 'Y', {'T', 'A', 'C'}: 'Y',
		{'T', 'A', 'A'}: '_', {'T', 'A', 'G'}: '_', {'T', 'G', 'A'}: '_',
	}

	// DEBUG Check syntax
	value, ok := tc[*c]
	if ok {
		return value
	}

	return '!'
}

// Entry Library -> May refactor later
type Entry struct {
	nucAcid  DNA
	rnucAcid RNA
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

func (s *DNA) gcContent() float64 {
	return 0.0
}

// only use if the user requests a set of subsequences
func (s *DNA) subseqGcContent(entry int, end int) float64 {
	return 0.0
}

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
		fmt.Scanf("%c", &choice)
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

		for {
			fmt.Println("Enter starting nucleotide followed by ending nucleotide: ")

			//fmt.Scanf("%d\n", &start)
			fmt.Scanf("%d\n%d\n", &start, &end)

			//fmt.Println(start)
			//fmt.Println(end)
			if start >= 0 && end < d.size && start < end {
				break
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
			fmt.Println(r.mRNA[k+i])
			i++
		}
		fmt.Println("Printing current insert")
		fmt.Println(insert)
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
	for _, resCodon := range m.codons {
		residue := resCodon.translate()
		tempAA = append(tempAA, residue)
	}

	p.aminoacid = string(tempAA)
	return &p
}

func main() {

	// Testing DNA component
	fmt.Println("Enter Sequence: ")
	var seq string
	fmt.Scanf("%s", &seq)
	fmt.Println(seq)

	testDNA := newDNA("test", seq)
	testDNA.printStrands()

	// Testing RNA component
	testRNA := transcribe(*testDNA)
	testRNA.mrnaPrint()
	// Testing Protein Component
}
