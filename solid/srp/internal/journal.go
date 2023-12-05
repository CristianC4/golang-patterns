package internal

import (
	"fmt"
	"time"
)

type Journal struct {
	id      uint
	entries []string
}

func (j *Journal) AddEntry(text string) {
	j.id++
	newEntry := fmt.Sprintf("[%d] (%v) \t\t%s\n", j.id, time.Now().UTC().Format("04-12-2023"), text)
	j.entries = append(j.entries, newEntry)
}

func (j *Journal) String() string {
	return fmt.Sprint(j.entries)
}

func (j *Journal) RemoveEntry(index int) {
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
	// in this case we leave the compiler counts for us the length of slice ...
}

func NewJournal(id uint) *Journal {
	return &Journal{
		id:      id,
		entries: []string{},
	}
}
