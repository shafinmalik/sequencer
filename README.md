# sequencer.go 
### (Work in Progress)
 A simple bioinformatics tool for finding the similarity between amino acid sequences. 

 ### TODO: 
 - Implement main in seq.go
 - Refactor for "while" loops
 - Implement fasta format read/write in genetext.go
 - Implement similarity matrix generator (simindex.go)

# fasta.go
This package can be used to read text files formatted in 120 character fasta format.
> TODO : Possibly add variable to describe if the text is a gene or protein

## genetext.fread()
Read-in constructor for the fasta struct which will search for a file in the name provided by the user. 

## genetext.ftouch()
Constructor for the fasta struct meant to be used with seq.go with user inputs rather than files. 

## genetext.fwrite() 
> Todo : Change to void 

Write to file using a fasta in memory. 

# simindex.go
> TODO : 