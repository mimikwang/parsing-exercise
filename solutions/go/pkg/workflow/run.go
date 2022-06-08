package workflow

func Run(inputDir, outputDir string) error {
	inputCh := make(chan inputData, 5)
	outputCh := make(chan outputData, 5)
	doneCh := make(chan error)

	go read(inputDir, inputCh, doneCh)
	go process(inputCh, outputCh, doneCh)
	go write(outputDir, outputCh, doneCh)

	err := <-doneCh
	if err != nil {
		return err
	}
	return nil
}
