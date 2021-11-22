package main

import "fmt"

func main(){
	var number int
	number = 15

	if number % 2 == 0{
		fmt.Println(number, "is even")
	}else if number % 3 == 0{
		fmt.Println(number, "is a multiple of 3")
	}else{
		fmt.Println(number, "is odd")
	}

	favFruit := "Banana"

	switch favFruit{
	case "Apple":
		fmt.Println("my favourite fruit is", favFruit)
	case "Banana":
		fmt.Println("I love", favFruit)
	default:
		fmt.Println("I don't think I like it..")
	}

	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "is even")
		// } else if i % 3 == 0 {
		// 	fmt.Println(i, "is a multiple of 3")
		} else {
			fmt.Println(i,"is odd")
		}
	}
}