package main

import "fmt"

var firstName = "Noah"
var lastName = "Krummenacher"
var dayOfBirth = 26
var monthOfBirth = 12
var yearOfBirth = 2005
var numberOfSiblings = 2
var heightInMeters = 1.77
var zodiacSign = '\u2651'

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
