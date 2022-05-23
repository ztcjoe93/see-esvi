package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

var (
	targetDirectory *string
	isRecursive     *bool
	// targetFieldIndex  *int
	// targetHeaderField *string
)

func parseCsv(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to read input file"+path, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file csv "+path, err)
	}

	return records
}

// fetches all .csvs given a directory path
func fetchCsvs(directory string, isRecursive bool) {
	if directory == "" {
		directory, err := os.Getwd()
		if err != nil {
			log.Fatal("Unable to get working directory"+directory, err)
		}
	}

	var files []string

	if isRecursive {
		err := filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() {
				if matched, err := filepath.Match("*.csv", filepath.Base(path)); err != nil {
					return err
				} else if matched {
					files = append(files, path)
				}
			}
			return nil
		})

		if err != nil {
			log.Fatal("Directory provided is invalid"+directory, err)
		}

		return
	}

	fmt.Println(files)
}

// utility function to parse all cli arguments
func cliArgParse() {
	targetDirectory = flag.String("t", "", "target directory to look for csv files to parse")
	isRecursive = flag.Bool("r", false, "should recursively look for files or not")
	// targetFieldIndex = flag.Int("f", 0, "target field (index) to extract value from")
	// targetHeaderField = flag.String("h", "", "target header to extract values from")
	flag.Parse()

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("%s: %s\n", f.Name, f.Value)
	})

	testValue := flag.Args()

	var i interface{}
	i = testValue
	fmt.Println(reflect.TypeOf(testValue))
	fmt.Println(testValue)
	if _, isList := i.([]string); isList {
		log.Fatal("Multiple filepaths provided")
	}
}

func main() {
	cliArgParse()

	// fetchCsvs(*targetDirectory, *isRecursive)

	// exPath := filepath.Dir(parentPath)
}
