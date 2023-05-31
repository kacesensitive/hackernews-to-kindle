package main

import (
	"fmt"
	"os"
	"os/exec"
)

func writeToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}

func convertHTMLToMobi(htmlFileName string) error {
	cmd := exec.Command("/opt/homebrew/bin/ebook-convert", htmlFileName, mobiFileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error converting HTML to MOBI: %v, output: %s", err, output)
	}
	return nil
}
