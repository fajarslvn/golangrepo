package main

import "fmt"

type User struct {
	FirstName, LastName string
	Age int
}

func main() {
	grades := make(map[string]int)

	grades["John doe"] = 33
	grades["George foo"] = 12
	grades["Mike king"] = 56

	delete(grades, "Mike king")
	fmt.Println("grades:", grades)

	userDatabase := make(map[string]*User)

	userOne := User{FirstName: "John", LastName: "Doe", Age: 45}
	userTwo := User{FirstName: "George", LastName: "Foo", Age: 51}
	userThree := User{FirstName: "Mike", LastName: "King", Age: 42}

	userDatabase["1"] = &userOne
	userDatabase["2"] = &userTwo
	userDatabase["3"] = &userThree

	fmt.Println("User database:", userDatabase)
	fmt.Println(userDatabase["1"].FirstName, userOne.FirstName)

	for key, value := range userDatabase {
		fmt.Println("user id:", key, "user:", value.FirstName)
	}
}