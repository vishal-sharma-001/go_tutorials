package main

import (
	"errors"
	"fmt"
	"strings"
)

type Cars struct{

}

func main() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String
	
	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
	printMe()
	printName("vishal")
	ans, rem, err := divide(14,2)
	if(err != nil){
		fmt.Printf("Res: %v   Rem: %v", ans,rem)
	}
	
	//Arrays, slices, maps, Looping Control Structures

	var intArr[3] int
	intArr[0] = 0  // assign value
	fmt.Println(intArr[0]) //to print
	fmt.Println(intArr[1:3]) //to print multiple elements
	fmt.Println(&intArr[0])  //TO PRINT MEMORY ADDRESS


	//slices
	var slice1[] int   //declaration
	slice1 = append(slice1, 1245)  //to append
	fmt.Println(slice1)  //to print

	//declaration + initialization
	var slice2[] int= []int{1,2,3}
	fmt.Printf("This is length of the slice %v.", len(slice2))
	fmt.Printf("This is capacity of the slice %v.", cap(slice2))

	
	slice1 = append(slice1, slice2...)

	// var myslice []int = make([]int, 3, 8)
	
	//map 

	var mymap map[string]int = make(map[string]int)   //declaration
	var mymap2 = map[string]int{"Visz": 25, "Hello": 21}  //declaration + intitalization

	val, ok := mymap["hi"];
	if ok {
		fmt.Println(val);
	}else{
		fmt.Println("The value does not exist")
	}
	delete(mymap2,"Visz")

	for val:= range mymap2{
		fmt.Println(val)
	} 

	for idx, v := range  slice1 {
		fmt.Printf("idx %v | val %v \n",idx, v)
	}
	
	i :=0
	for i < 10{
		i++;
	}
	for j:=0; j<10 ; j++{

	}

	str := "résumé"
	fmt.Printf("character: %v  | Type: %T\n", str[0], str[0])


	for i, v := range str{
		fmt.Printf("%v   %v\n", i, v)
	}

	name := "vishal"
	lastname := "sharma"

	name= name + lastname

	var sb strings.Builder
	sb.WriteString(name)
	sb.WriteString(lastname)

	// var resString = sb.String()

	//Structs and Interfaces in golang 

	test()
}

// functions
func printMe() {
	fmt.Println("HI");
}


//function parameters
func printName(name string) {
	fmt.Println(name);
}

//Handling errors
func divide(a int, b int) (int, int, error) {
	var err error  // by default err value is nil
	if(b != 0){
		return a/b, a%b, err
	}else{
		err = errors.New("divide by zero error")
	}
	
	return 0, 0, err
}
