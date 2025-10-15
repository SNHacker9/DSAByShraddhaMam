package main

import "DSAByShradhaMam/patterns"

func main(){
	// b:=patterns.NewSquarePatternBox(3,4,patterns.IncreasingChar)

	// b.PrintPattern()

	s:=patterns.NewStarPattern(4,patterns.InvertedTrianglePattern)

	s.PrintPattern()
}