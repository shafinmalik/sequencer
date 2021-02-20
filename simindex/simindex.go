package simindex

import (
	"image"
)

type dotplot struct {
	plot   image.RGBA
	matrix [][]rune
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
