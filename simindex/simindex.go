package simindex

import (
	"image"
)

type dotplot struct {
	plot   image.RGBA
	matrix [][]rune
	seq0   []rune
	seq1   []rune
	gene   bool // false if amino acid sequence
}

// TODO
func makeDotplot() *dotplot {
	return &dotplot{}
}

// Produces matrix
func (plt *dotplot) makeMatrix(nrows int, mcols int) {

}

func (plt *dotplot) makePlot(seq1 string, seq2 string) {

}
