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
func newParse(bytes []byte) (p parse.Parse, err error) {
	parsers := []func(bytes []byte) (parse.Parse, error){
		a.NewInput,
		b.NewInput,
	}

	for _, parser := range parsers {
		if p, err = parser(bytes); err == nil {
			return p, nil
		}
	}
	return nil, err
}
