# See-ESVI, a CSV parser and modifier utility tool

A utility tool to parse large amount of CSVs and modify targeted fields.

## How to run the tool

### Building
Build the binary file with
```shell
$ go build
```


### Running
Execute the binary file with the necessary arguments
```shell
$ ./see-esvi <flags> <command> <directory_with_csv_files>
```

Or simply run it
```shell
$ go run see-esvi.go <flags> <command> <directory_with_csv_files>
```

### Tests
Run tests using
```shell
$ go test -v
```
If you don't have gcc, you can run the test with
```shell
$ go test -v -vet=off
```
or you can simply set your `CGO_ENABLED=0` in your environment
```shell
$ export CGO_ENABLED=0
```

For information on why the cgo tool is required, you can refer to this [link](https://pkg.go.dev/cmd/cgo#:~:text=The%20cgo%20tool%20is%20enabled,to%200%20to%20disable%20it.)


## Commands
| Command | Description |
| --- | --- |
| read | Simple read command on csvs, will return target field values |
| modify | Modify command on csvs, will modify target field value to specified value | 

## CLI arguments
| Flag | Description |
| --- | --- |
| -r | Look for .csv files recursively in the given directory
| -tf | Target field index or name to retrieve values from 