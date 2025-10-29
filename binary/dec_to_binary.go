package binary


func DecToBinary(n int) int{

	ans:=0
	pow:=1 // 10^0
  for n >0 {
	rem:=n%2
	n=n/2

	ans+=(pow * rem)

	pow = pow * 10
  }

  return ans
}