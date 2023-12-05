package internal

import (
	"fmt"
	"os"
)

type PersistenceManager interface {
	Save(fileName, extensionName string) error
	LoadAll(fileName, extensionName string) error
}

type PersistenceFile struct {
	fileName      string
	extensionName string
	journal       *Journal
}

func (pf *PersistenceFile) Save(fileName, extensionName string) error {
	err := os.WriteFile(fileName+extensionName, []byte(pf.journal.String()), 0644)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}
	return nil
}

func (pf *PersistenceFile) LoadAll(fileName, extensionName string) error {
	fmt.Printf(" .: Loading from %s:. \n", fileName)
	entriesLoaded, err := os.ReadFile(fileName + extensionName)
	if err != nil {
		return fmt.Errorf("error loading file: %v", err)
	}
	fmt.Printf(" %s \n", entriesLoaded)
	return nil
}

func NewPersistenceFile(fileName, extensionName string, journal *Journal) PersistenceManager {
	return &PersistenceFile{fileName: fileName, extensionName: extensionName, journal: journal}
}
