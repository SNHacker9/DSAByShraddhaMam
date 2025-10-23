package main

import "DSAByShradhaMam/patterns"

func main(){
	// b:=patterns.NewSquarePatternBox(3,4,patterns.IncreasingChar)

	// b.PrintPattern()

	// s:=patterns.NewStarPattern(4,patterns.InvertedTrianglePattern)

	// s.PrintPattern()

	// py:=patterns.NewPyramidPattern(4)

	// py.PrintPattern()

	d:=patterns.NewDiamond(4,patterns.HollowDiamond)

	d.PrintPattern()
}