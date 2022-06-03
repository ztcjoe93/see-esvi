package main

import (
	"encoding/csv"
	"flag"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

var (
	isRecursive = flag.Bool("r", false, "should recursively look for files or not")
	targetField = flag.Int("tf", 0, "target field by index for value lookup")
	sugar       *zap.SugaredLogger
	dataSlice   []*Data
)

type Data struct {
	name    string
	headers []string
	values  [][]string
}

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

func main() {
	cfg := zap.NewDevelopmentConfig()

	_, err := os.Stat(".")
	cfg.OutputPaths = []string{"./debug.log", "stderr"}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugar = logger.Sugar()
	path := cliArgParse()
	csvFiles := fetchCsvs(path, *isRecursive)
	parseCsv(csvFiles)
}
