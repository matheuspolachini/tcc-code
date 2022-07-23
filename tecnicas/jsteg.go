package main

import (
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"lukechampine.com/jsteg"
)

func encodeImages() {
	rand.Seed(time.Now().UnixNano())
	baseDir := "/media/external/tcc_work"
	sourceDir := filepath.Join(baseDir, "original_jpg")
	dir, err := os.Open(sourceDir)
	if err != nil {
		panic(err)
	}

	files, err := dir.ReadDir(0)
	if err != nil {
		panic(err)
	}

	// shuffle file list
	for i := range files {
		j := rand.Intn(i + 1)
		files[i], files[j] = files[j], files[i]
	}

	filesCount := len(files)
	fmt.Println("filesCount: ", filesCount)

	firstHalf := files[:filesCount/2]
	fmt.Println("firstHalf: ", len(firstHalf))

	secondHalf := files[filesCount/2:]
	fmt.Println("secondHalf: ", len(secondHalf))

	targetDir := filepath.Join(baseDir, "jsteg")
	_ = os.Mkdir(targetDir, 0755)

	for _, file := range firstHalf {
		fmt.Println("Encoding file ", file.Name())
		fullPath := filepath.Join(sourceDir, file.Name())
		targetPath := filepath.Join(targetDir, fmt.Sprintf("encoded_%s", file.Name()))
		encodeRandomString(fullPath, targetPath)
	}

	for _, file := range secondHalf {
		fmt.Println("Copying file " + file.Name())
		fullPath := filepath.Join(sourceDir, file.Name())
		outputPath := filepath.Join(targetDir, file.Name())
		copyFile(fullPath, outputPath)
	}
}

func generateRandomAsciiString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func copyFile(sourcePath string, destPath string) {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
}

func encodeRandomString(coverPath string, outputFilePath string) {
	f, err := os.Open(coverPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	stats, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := stats.Size()

	stringLength := 10
	if fileSize > 10000 {
		stringLength = 500
	} else if fileSize > 4000 {
		stringLength = 100
	} else if fileSize > 3000 {
		stringLength = 50
	}

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	randomString := generateRandomAsciiString(stringLength)
	jsteg.Hide(outputFile, img, []byte(randomString), nil)
}

func main() {
	encodeImages()
}
