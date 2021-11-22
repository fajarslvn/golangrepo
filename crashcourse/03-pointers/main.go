package main

import "fmt"

func main(){
	var myFavFruit string
	myFavFruit = "Durian"
	
	fmt.Println("My favourite fruit is", myFavFruit)
	
	myFavFruit = "Orange"
	fmt.Println("My favourite fruit is", myFavFruit)
	
	changefavFruit(&myFavFruit, "Apple")
	fmt.Println("My new favourite fruit is", myFavFruit)
}

func changefavFruit(fruitpointer *string, newFavFruit string){
	fmt.Println("The address is:", fruitpointer, "The value is:", *fruitpointer)
	*fruitpointer = newFavFruit
}