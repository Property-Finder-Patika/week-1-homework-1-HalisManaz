package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name string
}

func main() {

	fmt.Println("Selam :)")

	var name string = "Halis"
	var greeting = createGreet(name)
	fmt.Printf("%s\n", greeting)

	greet(name)

	fmt.Printf("\n\n")

	var greeter = Person{name: "Halis"}
	var myfriend = Person{name: "Zeynep"}

	greeting = greeter.greet()
	fmt.Printf("%s\n", greeting)

	greeting = greeter.greetSomebody(myfriend)
	fmt.Printf("%s\n", greeting)

	fmt.Printf("\n\n")

	greetPrinter(createGreetInTurkish, "Hatice")
	greetPrinter(createGreetInEnglish, "Mary")
	greetPrinter(convertToUpperCase, "Naber?")

	fmt.Printf("\n\n")

	closure := func(name string) {
		greeting := "Selam " + name + " :)"
		fmt.Printf("%s", greeting)
	}
	closure("Fatma")
}

func greet(name string) {
	fmt.Printf("Selam %s :)\n", name)
}

func createGreet(name string) string {
	greeting := "Selam " + name + " :)"
	return greeting
}

func (p Person) greet() string {
	return "Selam " + p.name + " :)"
}

func (p Person) greetSomebody(q Person) string {
	return "Selam from " + p.name + " to " + q.name + " :)"
}

func createGreetInTurkish(name string) string {
	return "Selam " + name + " :)"
}

func createGreetInEnglish(name string) string {
	return "Hi " + name + " :)"
}

func convertToUpperCase(name string) string {
	return strings.ToUpper(name)
}

func greetPrinter(function func(it string) string, name string) {
	var greeting = function(name)
	fmt.Printf("%s\n", greeting)
}
