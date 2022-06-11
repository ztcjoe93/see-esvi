package main

import (
	"encoding/csv"
	"flag"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"go.uber.org/zap"
)

var (
	isRecursive = flag.Bool("r", false, "should recursively look for files or not")
	valFlag     = flag.String("v", "", "value to be used in commands")
	// if no arguments are provided, initialize to look at index 0 of csv
	targetField interface{} = 0
	cmdVal      string

	sugar     *zap.SugaredLogger
	dataSlice []*Data
)

func init() {
	// checking if -tf argument is a field index or field name
	flag.Func("tf", "target field by index or name for value lookup", func(stringVal string) error {
		intVal, err := strconv.Atoi(stringVal)
		if err != nil {
			targetField = stringVal
		} else {
			targetField = intVal
		}
		return nil
	})
}

type Data struct {
	name    string
	headers []string
	values  [][]string
}

// convenience function to initialize a Data struct
func newData(path string, csvRecord [][]string) *Data {

	data := Data{name: filepath.Base(path)}
	data.headers = csvRecord[0]
	data.values = csvRecord[1:]

	return &data
}

// parses all files retrieved sequentially
func parseCsv(files []string) {
	for _, file := range files {
		filePath, err := os.Open(file)
		if err != nil {
			sugar.Errorw("Error opening .csv to parse",
				"error_message", err,
				"filePath_path", filePath,
			)
		}

		defer filePath.Close()
		csvReader := csv.NewReader(filePath)
		records, err := csvReader.ReadAll()
		if err != nil {
			sugar.Errorw("Error parsing .csv filePath",
				"error_message", err,
				"filePath_path", filePath,
			)
		}

		parsedRecord := newData(file, records)
		dataSlice = append(dataSlice, parsedRecord)

		sugar.Infow("Success in parsing .csv file",
			"file_path", parsedRecord.name,
			"headers", parsedRecord.headers,
			"values", parsedRecord.values,
		)
	}
}

// fetches all .csvs given a directory path
func fetchCsvs(directory string, isRecursive bool) []string {
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
			sugar.Errorw("Error when recursively visiting directories",
				"error_message", err,
			)
		}

		sugar.Infow("Retrieved all .csv files successfully.",
			"recursive", isRecursive,
			"files", files,
		)

		return files
	}

	items, _ := ioutil.ReadDir(directory)
	for _, item := range items {
		if !item.IsDir() {
			path := filepath.Join(directory, item.Name())
			matched, _ := filepath.Match("*.csv", filepath.Base(path))
			if matched {
				files = append(files, path)
			}
		}
	}
	sugar.Infow("Retrieved all .csv files successfully.",
		"recursive", isRecursive,
		"files", files,
	)

	return files
}

func main() {
	cfg := zap.NewDevelopmentConfig()

	outputPath := filepath.Join(".", "output")
	_ = os.MkdirAll(outputPath, os.ModePerm)

	cfg.OutputPaths = []string{"./debug.log", "stderr"}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugar = logger.Sugar()
	cmd, path := cliArgParse()
	csvFiles := fetchCsvs(path, *isRecursive)
	parseCsv(csvFiles)
	cmd()
}
