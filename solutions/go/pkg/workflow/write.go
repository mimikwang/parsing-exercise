package workflow

import (
	"log"
	"os"
	"path"
)

func write(outDir string, inputCh <-chan outputData, doneCh chan<- error) {
	defer close(doneCh)
	for input := range inputCh {
		log.Printf("writing %s", input.fileName)
		file, err := os.Create(path.Join(outDir, input.fileName))
		if err != nil {
			file.Close()
			doneCh <- err
			return
		}
		_, err = file.Write(input.output.ToBytes())
		if err != nil {
			file.Close()
			doneCh <- err
			return
		}
		file.Close()
	}
	doneCh <- nil
}
