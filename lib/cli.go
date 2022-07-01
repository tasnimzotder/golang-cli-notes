package lib

import (
	"os"
	"strings"
)

func getArguments() []string {
	return os.Args

}

func argsController() {
	args := getArguments()
	args = args[1:]

	if len(args) == 1 {
		arg := args[0]

		if arg[0] == '-' {
			if arg == "-h" {
				printHelp()
			} else if arg == "-l" {
				printLastNotes(5)
			} else if arg == "-la" {
				printAllNotes()
			} else if arg == "-ex" {
				exportNotes(args[1])
			} else if arg == "-del" {
				deleteNotes()
			} else if strings.Contains(arg, "-") && getNumber(arg[1:]) > 0 {
				printLastNotes(getNumber(arg[1:]))
			} else {
				takeNote(arg)
			}
		}
	}

	if len(args) >= 1 && args[0][0] != '-' {
		note := strings.Join(args, " ")

		takeNote(note)
	}

	if len(args) == 0 {
		printLastNotes(5)
	}
}

func CLI() {
	argsController()
}
