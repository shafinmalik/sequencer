package genetext

import (
	"fmt"
	"io/ioutil"
)

// Use this pkg to export in FASTA format

// fasta : each string in []string should have a max of 120 characters. dscr should be a max of 120 characters
type fasta struct {
	title string
	dscr  string
	nAcid []string
}

// fasta read constructor
func fread() *fasta {
	dat, err := ioutil.ReadFile("/tmp/dat")

	// temporary for debugging
	fmt.Printf(string(dat), err)
	// End Debugging

	return &fasta{}
}

// TODO: fasta constructor without file input
func ftouch() *fasta {
	return &fasta{}
}

// fasta write output (possibly change return type to int or bool)
func (f *fasta) fwrite() {

}

// formatting output
// place all components into 1 string, then insert '\n' newlines
func (*fasta) formatfastaR(name string, nuc string) {

}

// panic() if error in output
func check(e error) {
	if e != nil {
		panic(e)
	}
}
