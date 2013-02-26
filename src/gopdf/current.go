package gopdf

import (
	
)

type Current struct{
	X float64
	Y float64
	
	//font
	IndexOfFontObj int
	CountOfFont int
	CountOfL int
	
	Font_Size int
	Font_Style string
	Font_IFont IFont
	Font_FontCount int
	
	//page
	IndexOfPageObj int
}