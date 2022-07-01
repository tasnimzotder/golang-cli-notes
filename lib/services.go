package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

var notes_file string = "/tmp/notes/.notes"

type NotesT struct {
	LoggedAt time.Time `json:"loggedAt"`
	Note     string    `json:"note"`
}

func printArgLine(arg string, description string, example string) {
	fmt.Println("\t", arg, "\t: ", description, "\t\te.g.  ", example)
}

func printHelp() {
	fmt.Println("\nA CLI app to take notes")
	fmt.Print("\n")

	fmt.Println("\tnote\t: take notes")

	fmt.Print("\n")
	fmt.Println("arguments:")

	arguments := [][]string{
		{"-h", "Show all the commands", "note -h"},
		{"-l", "Show last 5 notes    ", "note -l"},
		{"-la", "Show all the notes  ", "note -la"},
		{"-ex", "Export all the notes", "note -ex notes"},
		{"-del", "Delete all the notes", "note -del"},
		{"-{n}", "SShow last {n} notes", "note -13"},
	}

	for _, arg := range arguments {
		printArgLine(arg[0], arg[1], arg[2])
	}
}

func fileExistCheck(filename string) {
	if _, err := os.Stat(filename); err != nil {
		os.Create(filename)
	}

}

func takeNote(note string) {
	fileExistCheck(notes_file)

	file, err := ioutil.ReadFile(notes_file)
	if err != nil {
		log.Println("Error reading", err)
	}

	// check if notes_file is present

	var notes []NotesT
	json.Unmarshal(file, &notes)

	nNote := NotesT{time.Now(), note}

	notes = append(notes, nNote)

	json, err := json.MarshalIndent(notes, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	err = ioutil.WriteFile(notes_file, json, 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func printAllNotes() {
	printLastNotes(math.MaxInt)
}

func printLastNotes(lines int) {
	fileExistCheck(notes_file)

	file, err := ioutil.ReadFile(notes_file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var notes []NotesT
	json.Unmarshal(file, &notes)

	if len(notes) < lines {
		lines = len(notes)
	}

	for i := len(notes) - lines; i < len(notes); i++ {
		fmt.Println(notes[i].LoggedAt.Format(time.RFC822), ": ", notes[i].Note)
	}
}

func deleteNotes() {
	file, err := ioutil.ReadFile(notes_file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var notes []NotesT
	json.Unmarshal(file, &notes)

	notes = []NotesT{}

	json, err := json.MarshalIndent(notes, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	err = ioutil.WriteFile(notes_file, json, 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Notes deleted")
}

func exportNotes(filename string) {
	file, err := ioutil.ReadFile(notes_file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var notes []NotesT
	json.Unmarshal(file, &notes)

	json, err := json.MarshalIndent(notes, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	err = ioutil.WriteFile(filename, json, 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Notes exported")
}

func getNumber(arg string) int {
	if len(arg) == 0 {
		return math.MaxInt
	}

	number, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Error converting to number:", err)
		return math.MaxInt
	}

	return number
}
