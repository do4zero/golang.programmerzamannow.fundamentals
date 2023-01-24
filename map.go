package main

import "fmt"

func main(){
	// person := map[string]string{
	// 	"name" : "Eko",
	// 	"address" : "Subang",
	// }

	var person map[string]string = make(map[string]string)

	person["name"] = "eko"
	person["address"] = "BTN Griya"
	person["nope"] ="6289507854000"

	fmt.Println(person)

	delete(person,"address")
	
	fmt.Println(person)
}