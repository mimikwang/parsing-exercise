package workflow

import (
	"log"

	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse"
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse/a"
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse/b"
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse/output"
)

type outputData struct {
	fileName string
	output   output.Output
}

// process input bytes into output
func process(inputCh <-chan inputData, outputCh chan<- outputData, statusCh chan<- error) {
	defer close(outputCh)
	for input := range inputCh {
		log.Printf("processing %s", input.fileName)
		parser, err := newParse(input.data)
		if err != nil {
			log.Printf("error processing %s", input.fileName)
			continue
		}
		output, err := parser.Parse()
		if err != nil {
			statusCh <- err
			return
		}
		outputCh <- outputData{fileName: input.fileName, output: output}
	}
}

// newParse is a factory method for a parse.Parse
func newParse(bytes []byte) (parse.Parse, error) {
	parserA, err := a.NewInput(bytes)
	if err == nil {
		return parserA, nil
	}
	parserB, err := b.NewInput(bytes)
	if err == nil {
		return parserB, nil
	}
	return nil, err
}
