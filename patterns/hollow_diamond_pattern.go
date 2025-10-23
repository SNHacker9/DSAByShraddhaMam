package patterns

import "fmt"

type hDiamond int

const(
	HollowDiamond hDiamond= iota
)
type diamond struct{
	height int
	hDiamond
} 

func NewDiamond(h int , t hDiamond)PatternBoxes{
	return &diamond{
        height: h,
		hDiamond: t,
	}
}

func(d diamond)PrintPattern(){
  switch d.hDiamond{
  case HollowDiamond:d.printHollowDiamond()
  }
}

func(d diamond)printHollowDiamond(){
     //outer loop

	 //top
	 for i:=0;i<d.height;i++{
      // spaces
      for j:=0;j<d.height-i-1;j++{
		fmt.Print(" ")
	  }

	  fmt.Print("*")

	  if i!=0{
		// spaces
		for j:=0;j<2*i-1;j++{
			fmt.Print(" ")
		}

		fmt.Print("*")
	  }

	  fmt.Print("\n")

	 }

	 //bottom

	 for i:=0;i<d.height-1;i++{
		//spaces 
        for j:=0;j<i+1;j++{
			fmt.Print(" ")
		}

		fmt.Print("*")
       
		if i!=d.height-2{
			//spaces
            for j:=0;j<2*(d.height-i)-5;j++{
              fmt.Print(" ")
			}
		fmt.Print("*")
		}


		fmt.Print("\n")
		
	 }
}