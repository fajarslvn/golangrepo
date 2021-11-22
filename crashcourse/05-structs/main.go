package main

import "fmt"

type Person struct {
	FirstName, 
	LastName, 
	Country string
	Age int
}

type Book struct {
	Title, 
	Author string
	Price int
}

type DemonSlayer struct {
	Title,
	Starring,
	Seasons string
	Duration int
}

func main() {
	personOne := Person {
		FirstName: "Demon",
		LastName: "King",
		Age: 300,
		Country: "Hell",
	}

	bookOne := Book {
		Title: "Lord of the Rings",
		Author: "Tolkiens",
	}

	demonSlayer := DemonSlayer {
		Title: "End of game",
		Starring: "Tanjirou",
		Duration: 120,
	}

	// fmt.Println("His name is", getFullName(personOne)) ex 1
	fmt.Println("His name is", personOne.getFullName()) // ex 2

	fmt.Println("bookOne title:", bookOne.Title)
	bookOne.changeTitle("Harry Potter")
	fmt.Println("bookTwo title:", bookOne.Title)

	demonSlayer.animeTitle("The Beginnings")
	fmt.Println("Episode of", demonSlayer.Title)
}

/* ex 1
func getFullName(person Person) string {
	return person.FirstName + " " + person.LastName
}*/

// ex 2
func (person *Person) getFullName() string {
	return person.FirstName + " " + person.LastName
}

func (book *Book) changeTitle(newTitle string) {
 	book.Title = newTitle
}

func (anime *DemonSlayer) animeTitle(chgTitle string) {
	anime.Title = chgTitle
}