package genetext

// Use this pkg to export in FASTA format

// each string in []string should have a max of 120 characters
type fasta struct {
	title     string
	dscr      string
	aminoacid []string
}

func (*fasta) formatfasta(name string, protein string) {

}

func (f *fasta) fread() {

}

func (f *fasta) fwrite() {

}
