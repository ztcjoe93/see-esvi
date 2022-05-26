# See-ESVI, a CSV parser and modifier utility tool

A utility tool to parse large amount of CSVs and modify targeted fields.

## How to run the tool
Build the binary file with
```shell
$ go build
```

Execute the binary file with the necessary arguments
```shell
$ ./see-esvi <flags> <directory_with_csv_files>
```

Or simply run it
```shell
$ go run see-esvi.go <flags> <directory_with_csv_files>
```

## CLI arguments
| Flag | Description |
| --- | --- |
| -r | Look for .csv files recursively in the given directory
| -tf | Target field index to retrieve values from 