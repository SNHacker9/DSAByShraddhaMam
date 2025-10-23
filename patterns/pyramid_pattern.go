package patterns

import "fmt"

type pyramidType int

const(
	pyramidPatter pyramidType = iota
)

type pryramid struct{
	height int
	pyramidType 
}

func NewPyramidPattern(h int)PatternBoxes{
	return &pryramid{height: h}
}

func(p pryramid)PrintPattern(){
     switch p.pyramidType{
	 case pyramidPatter:p.printPyramidPattern()
	 }
}

func(p pryramid)printPyramidPattern(){
	for i:=1;i<=p.height;i++{
		// print spaces
		for j:=i;j<p.height;j++{
			fmt.Print(" ")
		}
       
		//print numbers straight
		for k:=1;k<=i;k++{
			fmt.Print(k)
		}
        
		temp:=i-1
		//print number reverse
		for l:=1;l<i;l++{
           fmt.Print(temp)
		   temp--
		}

		fmt.Print("\n")

	}
}