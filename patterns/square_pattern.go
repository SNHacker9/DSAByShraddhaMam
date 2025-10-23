package patterns

import(
	"fmt"
)

type elemenType int

const(
Integer elemenType = iota
Char
IncreasingInteger  
IncreasingChar
)

type box struct{
  rows int
  cols int
  eleType elemenType
}

func NewSquarePatternBox(rows,cols int,eleType elemenType)PatternBoxes{
	return &box{
		rows: rows,
		cols: cols,
		eleType:eleType,
	}
}

func(b *box) PrintPattern(){
	switch b.eleType{
	case Integer:b.printNumber()
    case Char:b.printChar()
	case IncreasingInteger:b.printIncreasingNumber()
	case IncreasingChar: b.printIncreasingChar()
	}
}

func(b *box) printNumber(){
	for i:=1;i<=b.rows;i++{
		for j:=1;j<=b.cols;j++ {
		 fmt.Print(j," ")
		
		}
 
		fmt.Print("\n")
	 }
}


func(b *box)printChar(){
	for i:=1;i<=b.rows;i++{
		ch:='A'
		for j:=1;j<=b.cols;j++ {
		 fmt.Print(string(ch)," ")
		 ch+=1
		}
 
		fmt.Print("\n")
	}	
}

func(b *box) printIncreasingNumber(){
	temp:=1

	for i:=1;i<=b.rows;i++{
		for j:=1;j<=b.cols;j++{
			fmt.Print(temp, " ")
			temp++
		}

		fmt.Print("\n")
	}
}

func(b *box) printIncreasingChar(){
	char:='A'

	for i:=1;i<=b.rows;i++{
		for j:=1;j<=b.cols;j++{
			fmt.Print(string(char), " ")
			char++
		}

		fmt.Print("\n")
	}
}

func(e elemenType)String()string{
	switch e {
	case Integer:
		return "int"
	case Char:
		return "char"	
  case IncreasingInteger:
    return "inc_int"  
  case IncreasingChar:
    return "inc_char"
	default:
		return "unknown"	
	}
}