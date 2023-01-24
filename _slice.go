package main

import "fmt"

func main(){
	// var months = [...]string{
	// 	"Januari",
	// 	"Februari",
	// 	"Maret",
	// 	"April",
	// 	"Mei",
	// 	"Juni",
	// 	"Juli",
	// 	"Agustus",
	// 	"September",
	// 	"Oktober",
	// 	"November",
	// 	"Desember",
	// }

	// var slice1 = months[4:7]

	// fmt.Println(months)
	// fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei ( Update )"
	// fmt.Println(months)

	// var slice2 = months[11:]
	// fmt.Println(slice2)

	// var slice3 = append(slice2, "Due")
	// fmt.Println(slice3)
	// slice3[1] = "Desember Noted"
	// fmt.Println(slice3)

	// fmt.Println(slice2)
	// fmt.Println(months)

	// var nameSlice [5]string;
	// nameSlice[0] = "Luthfi"
	// fmt.Println(nameSlice)

	// newSlice := make([]string, 2,5)
	// newSlice[0] = "Eko"
	// newSlice[1] = "Kurniawan"

	// fmt.Println(newSlice)
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))

	// copySlice := make([]string, len(newSlice), cap(newSlice))
	// copy(copySlice, newSlice)
	// fmt.Println(copySlice)

	iniArray := [3]string{"Luthfi"}
	iniSlice := []string{"Luthfi"}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}