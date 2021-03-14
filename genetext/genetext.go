package genetext

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Use this pkg to export in FASTA format

// Constants
const linewidth int = 120

// fasta : each string in []string should have a max of 120 characters. dscr should be a max of 120 characters
type fasta struct {
	title string
	nAcid string
}

// Debugging for testing Read/Constructors
func (f fasta) GetVars() {
	fmt.Println("Name below: ")
	fmt.Println(f.title)
	fmt.Println("Sequence below: ")
	fmt.Println(f.nAcid)
}

// fasta read constructor
func Fread(filename string) *fasta {

	var name string
	var sequence string
	// for testing
	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileStrings []string

	for scanner.Scan() {
		fileStrings = append(fileStrings, scanner.Text())
	}

	file.Close()

	sequence = ""
	if len(fileStrings) >= 2 && fileStrings[0][0] == '>' {
		name = fileStrings[0]
		for i := 1; i < len(fileStrings); i++ {
			sequence = sequence + fileStrings[i]
		}
	} else {
		fmt.Println("Invalid Fasta File.")
	}

	return &fasta{title: name, nAcid: sequence}
}

// TODO: fasta constructor without file input
func Ftouch() *fasta {
	return &fasta{}
}

// fasta write output (possibly change return type to int or bool)
func (f *fasta) Fwrite() bool {
	// insert '\n' after title
	// insert '\n' to sequence every 120 runes
	// then append sequence to title.
	dat := f.title // + "\n"
	for i := 0; i < len(f.nAcid); i++ {
		if i%linewidth == 0 {
			f.nAcid = f.nAcid[:i] + "\n" + f.nAcid[i:]
		}
	}
	dat = dat + f.nAcid
	dump := []byte(dat)

	var filename string
	fmt.Println("Enter Filename: ")
	fmt.Scanf("%s\n", &filename)
	// filename = "/tmp/" + filename

	err := ioutil.WriteFile(filename, dump, 0644)
	// Write to file - return true if successful
	check(err)

	return true
}

// panic() if error in output
func check(e error) {
	if e != nil {
		panic(e)
	}
}
