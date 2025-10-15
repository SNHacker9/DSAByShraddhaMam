package patterns

import "fmt"

type pType int

const(
	Triangle pType = iota
	TraingleNumbers
	TraingleChar
	TriangleNumberReset
	FloydTrianglePattern
	ReverseTriangleNumberReset
	InvertedTrianglePattern
)

type star struct{
	rows int
	patternType pType
}
	
func NewStarPattern(rows int,pType pType)PatternBoxes{
	return &star{rows: rows,patternType: pType}
}

func(s star)PrintPattern(){
    switch s.patternType{
	case Triangle: s.printTrianglePattern()
	case TraingleNumbers: s.printTriangleNumberPattern()
	case TraingleChar:s.printTriangleCharPattern()
	case TriangleNumberReset: s.printTriangleNumberResetPattern()
	case FloydTrianglePattern: s.printFloydTrianglePattern()
	case ReverseTriangleNumberReset: s.printReverseTriangleNumberResetPattern()
	case InvertedTrianglePattern: s.printInvertedTrianglePattern()
	}
}

func(s star)printTrianglePattern(){
	for i:=1;i<=s.rows;i++{
		for j:=0;j<i;j++{
		  fmt.Print("*"," ")
		}
  
		fmt.Print("\n")
	  }
}

func(s star)printTriangleNumberPattern(){
	for i:=1;i<=s.rows;i++{
		for j:=0;j<i;j++{
		  fmt.Print(i," ")
		}
  
		fmt.Print("\n")
	  }
}

func(s star)printTriangleCharPattern(){
	char:='A'

	for i:=1;i<=s.rows;i++{
		for j:=0;j<i;j++{
		  fmt.Print(string(char)," ")
		}
        char++
		fmt.Print("\n")
	  }
}

func(s star)printTriangleNumberResetPattern(){
	for i:=1;i<=s.rows;i++{
		for j:=1;j<i+1;j++{
		  fmt.Print(j," ")
		}
  
		fmt.Print("\n")
	  }
}

func(s star)printFloydTrianglePattern(){
	temp:=1
	for i:=1;i<=s.rows;i++{
		for j:=1;j<i+1;j++{
		  fmt.Print(temp," ")
		  temp++
		}
  
		fmt.Print("\n")
	  }
}

func(s star)printReverseTriangleNumberResetPattern(){
	for i:=1;i<=s.rows;i++{
		for j:=i;j>=1;j--{
		  fmt.Print(j," ")
		}
  
		fmt.Print("\n")
	  }
}

// Note: This is important can be revised.
func(s star)printInvertedTrianglePattern(){
	for i:=0;i<s.rows;i++{
		// print spaces
		for j:=0;j<i;j++{
		  fmt.Print(" ")
		}

		// print numbers
		for j:=0;j<s.rows-i;j++{
			fmt.Print(i+1)
		}
  
		fmt.Print("\n")
	  }
}