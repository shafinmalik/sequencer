package genetext

import (
	"fmt"
	"io/ioutil"
)

// Use this pkg to export in FASTA format

// fasta : each string in []string should have a max of 120 characters. dscr should be a max of 120 characters
type fasta struct {
	title string
	nAcid string
}

// fasta read constructor
func fread(filename string) *fasta {
	// for testing
	dat, err := ioutil.ReadFile("/tmp/" + filename)
	check(err)

	contents := string(dat)
	var start int
	foundStart := false
	var end int
	for i := 0; i < len(contents); i++ {
		// extract the title
		if contents[i] == '>' {
			start = i
			foundStart = true
		}

		if foundStart && contents[i] == '\n' {
			end = i
		}
	}

	newTitle := contents[start:end]
	fmt.Println(newTitle)
	// Now extract nAcid
	// temporary for debugging
	fmt.Println(contents, err)
	// End Debugging

	return &fasta{}
}

// TODO: fasta constructor without file input
func ftouch() *fasta {
	return &fasta{}
}

// fasta write output (possibly change return type to int or bool)
func (f *fasta) fwrite() bool {
	// insert '\n' after title
	// insert '\n' to sequence every 120 runes
	// then append sequence to title.
	dat := f.title + "\n"
	for i := 0; i < len(f.nAcid); i++ {
		if i%120 == 0 {
			f.nAcid = f.nAcid[:i] + "\n" + f.nAcid[i:]
		}
	}
	dat = dat + f.nAcid
	dump := []byte(dat)

	var filename string
	fmt.Println("Enter Filename: ")
	fmt.Scanf("%s\n", &filename)
	filename = "/tmp/" + filename

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
