package main

import (
	"fmt"
	"github.com/CristianC4/golang-patterns/solid/srp/internal"
)

func main() {
	journal := internal.NewJournal(1)
	journal.AddEntry("I studied today")
	journal.AddEntry("I created a bug")
	file := internal.NewPersistenceFile("journal", ".txt", journal)
	err := file.Save("journal", ".txt")
	if err != nil {
		panic(err)
	}
	err = file.LoadAll("journal", ".txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Finish Program\n")
}
