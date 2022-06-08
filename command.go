package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

// this file consists of logic processing for the command flag (<command> <csv_directory>)

// parse command from cli
func parseCommand(cmd string) (res func(), err error) {
	cmdVal = cmd

	switch cmd {
	case "read":
		return readData, nil
	case "modify":
		return modifyData, nil
	default:
		return nil, errors.New("invalid command")
	}
}

func modifyData() {
	// TODO logic to modify targeted fields to specific value
	// we may want to add some conditional parsing? (i.e x <= y, x > z)
	// may want to target multiple fields...? how would that affect the conditional parsing
	fmt.Println("modifying data...")
}

func readData() {
	var targetIndex int
	switch typeof(targetField) {
	case "string":
		for headerIndex, header := range dataSlice[0].headers {
			if header == targetField {
				targetIndex = headerIndex
				break
			}
		}
	case "int":
		targetIndex = targetField.(int)
	}

	for _, data := range dataSlice {
		for _, record := range data.values {
			fmt.Println(record[targetIndex])
		}
	}
}

// utility function to parse all cli arguments, returns
func cliArgParse() (func(), string) {

	flag.Parse()

	arguments := flag.Args()
	cmdFn, err := parseCommand(arguments[0])
	if err != nil {
		sugar.Panicw("Command passed does not exist",
			"function", "parseCommand()",
			"err", err.Error(),
		)
	}

	path := arguments[1:]

	if len(path) > 1 {
		log.Panic("Multiple filepaths provided")
	} else if len(path) == 0 {
		log.Panic("No filepath provided")
	}
	sugar.Infow("Retrieved directory pathname from cli arguments",
		"function", "cliArgparse()",
		"directory_pathname", path[0],
	)
	fmt.Printf("%s -> %s\n", path[0], cmdVal)

	return cmdFn, path[0]
}
