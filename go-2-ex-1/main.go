package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	DayOfBirht   int
	MonthOfBirth int
	YearOfBirth  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	NumberOfSiblings byte
	ZodiacSign       rune
	FullName         FullName
	BirthDate        BirthDate
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		NumberOfSiblings: 2,        // TODO: adjust
		ZodiacSign:       '\u2651', // TODO: adjust
		FullName: FullName{
			"Noah", "Krummenacher",
		},
		BirthDate: BirthDate{
			26, 12, 2005,
		},
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings += 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
