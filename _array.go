package main

import "fmt"

func main(){
	var names [4]string

	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy"
	names[3] = "Homan"

	var values = [3]int{
		90,
		100,
		90,
	}

	fmt.Println(values)
	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])
	fmt.Println(len(values))
	fmt.Println(len(names))
}