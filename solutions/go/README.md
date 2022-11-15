# parsing-exercise golang

Golang solution to [parsing-exercise](https://github.com/mimikwang/parsing-exercise).

## Requirements

* `go` version 1.18 and above

## Installation

First, run tests by running the following from this directory.

```
go test ./...
```

Make sure all tests are passing.

Then, build the command line took by running the following from this directory.

```
go build -o parse ./cmd
```

An executable `parse` will be created in this directory.

## Run on Example Files

Run the following to parse the example files.

```
mkdir output
./parse ../../files ./output
```

You should see the following logged to `stdout`:

```
2022/11/15 09:47:43 starting workflow...
2022/11/15 09:47:43 loading input001.json
2022/11/15 09:47:43 loading input002.json
2022/11/15 09:47:43 loading input003.json
2022/11/15 09:47:43 processing input001.json
2022/11/15 09:47:43 processing input002.json
2022/11/15 09:47:43 writing input001.json
2022/11/15 09:47:43 processing input003.json
2022/11/15 09:47:43 error processing input003.json
2022/11/15 09:47:43 writing input002.json
2022/11/15 09:47:43 success!
```

You should also find `input001.json` and `input002.json` in the `output` folder of the parsed output.

## Clean Up

Run the following to clean up.

```
rm -rf ./output
rm parse
```
