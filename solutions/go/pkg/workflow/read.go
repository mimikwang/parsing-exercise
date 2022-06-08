package workflow

import (
	stdio "io"
	"log"

	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/io"
)

type inputData struct {
	fileName string
	data     []byte
}

// read reads files from a directory and send bytes to a channel
func read(dirPath string, dataCh chan<- inputData, statusCh chan<- error) {
	defer close(dataCh)
	reader, err := io.NewReader(dirPath)
	if err != nil {
		statusCh <- err
		return
	}
	for {
		fileName, err := reader.FileName()
		if err == stdio.EOF {
			break
		}
		log.Printf("loading %s", fileName)
		if err != nil {
			statusCh <- err
			break
		}
		data, err := reader.Read()
		if err != nil {
			statusCh <- err
			break
		}
		dataCh <- inputData{fileName: fileName, data: data}
	}
}
