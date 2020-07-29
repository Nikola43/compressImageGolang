package main

import (
	"fmt"
	"github.com/h2non/bimg"
	"os"
)

func main() {
	inputFile := "assets/img/image.jpg"
	outputFile := "assets/img/imageOut.jpg"

	compressImage(inputFile, outputFile)

	fmt.Print(inputFile + "->" + string(getFileSize(inputFile)))
	fmt.Print(outputFile + "->" + string(getFileSize(outputFile)))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func getFileSize(filePath string) int64 {
	file, err := os.Stat(filePath)
	checkError(err)
	return file.Size()
}

func compressImage(inputFilePath, outFilePath string) {
	options := bimg.Options{
		Width:     800,
		Height:    600,
		Crop:      true,
		Quality:   95,
		Rotate:    180,
		Interlace: true,
	}

	// open file
	buffer, err := bimg.Read(inputFilePath)
	checkError(err)

	// process file
	newImage, err := bimg.NewImage(buffer).Process(options)
	checkError(err)

	// save file
	err = bimg.Write(outFilePath, newImage)
	checkError(err)
}
