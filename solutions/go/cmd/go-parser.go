package main

import (
	"log"
	"os"

	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/workflow"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		log.Fatalln("exactly 2 arguments required")
	}
	inputDir := args[1]
	outputDir := args[2]

	log.Println("starting workflow...")
	if err := workflow.Run(inputDir, outputDir); err != nil {
		log.Fatalln(err)
	}
	log.Println("success!")
}
