package binary

func BinaryToDecimal(n int)int{
	ans:=0
	pow:=1
	for n >0{
		rem:=n%10
		n=n/10

		ans+=(rem*pow)

		pow=pow*2
	}

	return ans
}