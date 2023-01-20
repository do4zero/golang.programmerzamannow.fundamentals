package main

import "fmt"

func main(){
	var name string

	name = "Luthfi Aziz Nugraha"
	fmt.Println(name)

	name = "Aziz Nugraha"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	// jika tidak menggunakan var gunakan := untuk deklarasi variable
	hello := "Say Hello"
	fmt.Println(hello)

	country := "Indonesia"
	fmt.Println(country)

	// multiple variable
	var (
		firstName = "Luthfi"
		lastName = "Aziz"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}