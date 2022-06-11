package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

// this file consists of logic processing for the command flag (<command> <csv_directory>)

// parse command from cli
func parseCommand(param string) (res func(), err error) {
	cmdVal := param

	switch cmdVal {
	case "read":
		return readData, nil
	case "modify":
		if *valFlag == "" {
			return nil, errors.New("no value provided to valFlag")
		}
		return modifyData, nil
	default:
		return nil, errors.New("invalid command")
	}
}

func getTargetField() int {
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

	return targetIndex
}

func modifyData() {
	// TODO logic to modify targeted fields to specific value
	// we may want to add some conditional parsing? (i.e x <= y, x > z)
	// may want to target multiple fields...? how would that affect the conditional parsing

	targetIndex := getTargetField()

	for _, data := range dataSlice {
		f, err := os.Create("./output/" + data.name)
		if err != nil {
			log.Panic(err)
		}
		defer f.Close()

		writer := csv.NewWriter(f)
		defer writer.Flush()

		writer.Write(data.headers)

		for _, record := range data.values {
			record[targetIndex] = *valFlag
			writer.Write(record)
		}
	}

}

func readData() {
	targetIndex := getTargetField()

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
	if len(arguments) == 0 {
		log.Panic("No arguments provided")
	}

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
