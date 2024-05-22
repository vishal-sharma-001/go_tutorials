package main

import (
	"fmt"
)

type owner struct{
	name string
}

type cars struct{
	model string
	cost int
	owner owner
}

func test(){
	
	fmt.Println("Hi")

	//create an array

	// var arr[5] int = [5]int{1, 2,}

	// //slice
	// var sc[] string 
	
	// var mymap map[string]int = make(map[string]int)

	// mymap["hi"] = 1

	// for keys, vals := range mymap {

	// }

	var mycar cars = cars{"Tata", 12, owner{""}}

	mycar.owner.name = "goku"

	anonymusStruct := struct{
		names string
		age int
	}{"hi", 20}

	fmt.Println(anonymusStruct)

}