package main

import (
	"flag"
	"log"
)

// pass interface into the function to return a string representation of type
func typeof(v interface{}) string {
	switch v.(type) {
	case string:
		return "string"
	case int:
		return "int"
	default:
		return "unknown"
	}
}

// utility function to parse all cli arguments
func cliArgParse() string {

	flag.Parse()
	targetPath := flag.Args()

	if len(targetPath) > 1 {
		log.Panic("Multiple filepaths provided")
	} else if len(targetPath) == 0 {
		log.Panic("No filepath provided")
	}
	sugar.Infow("Retrieved directory pathname from cli arguments",
		"directory", targetPath[0],
	)
	return targetPath[0]
}
