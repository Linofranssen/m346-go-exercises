package main

import "fmt"

func main() {
	var modules = map[int]string{
		104: "Datenmodell implementieren",
		117: "Informatik- und Netzinfrastruktur für ein kleines Unternehmen realisieren",
		346: "Cloud Lösungen konzipieren und realisieren",
	}
	// TODO: create a map called "modules"

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117)
	// TODO: delete one
	modules[322] = "Benutzerschnittstellen entwerfen und implementieren"
	// TODO: add one
	modules[322] = "Benutzerschnittstellen erstellen sowie implementieren"
	// TODO: replace one
	fmt.Println(modules)
}
