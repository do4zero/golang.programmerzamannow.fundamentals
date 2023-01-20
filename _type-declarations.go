package main

import "fmt"

func main(){
	type NoKTP string
	type Married bool
	
	var noKtpLuthfi NoKTP = "829918882899"
	var marriedStatus Married = true
	
	fmt.Println(noKtpLuthfi)
	fmt.Println(marriedStatus)
}