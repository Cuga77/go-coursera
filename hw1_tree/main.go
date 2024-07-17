package hw1_tree

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

const (
	firstLineSeparator  = "├───"
	mediumLineSeparator = "│"
	lastLineSeparator   = "└───"
)

// TODO:  ДЗ: придумать итеративный алгоритм. Без глобальных переменных.
func dirTree(output io.Writer, currDir string, printFiles bool) error {
	recursionPrintServise("", output, currDir, printFiles)
	return nil
}

func recursionPrintServise(prependingString string, output io.Writer, currDir string, printFiles bool) {
	//Взаимодействие с папкой
	fileObj, err := os.Open(currDir)
	defer fileObj.Close()
	if err != nil {
		log.Fatalf("Could not open %s: %s", currDir, err)
	}
	fileName := fileObj.Name()
	files, err := ioutil.ReadDir(fileName)
	if err != nil {
		log.Fatalf("Could not read directory %s: %s", currDir, err)
	}

	//Создаем и сортируем список файлов
	var fileMap map[string]os.FileInfo = map[string]os.FileInfo{}
	var unsortedNames []string = []string{}
	for _, file := range files {
		unsortedNames = append(unsortedNames, file.Name())
		fileMap[file.Name()] = file
	}
	sort.Strings(unsortedNames)
	var sortedFilesArr []os.FileInfo = []os.FileInfo{}
	for _, stringName := range unsortedNames {
		sortedFilesArr = append(sortedFilesArr, fileMap[stringName])
	}
	files = sortedFilesArr

	//Переходим к новому списку файлов
	var newFileList []os.FileInfo = []os.FileInfo{}
	var length int
	if !printFiles {
		for _, file := range files {
			if file.IsDir() {
				newFileList = append(newFileList, file)
			}
		}
		files = newFileList
	}
	length = len(files)

	//Печать папок и файлов
	for i, file := range files {
		if file.IsDir() {
			var stringPrepending string
			if length > i+1 {
				fmt.Fprintf(output, prependingString+firstLineSeparator+"%s\n", file.Name())
				stringPrepending = prependingString + mediumLineSeparator + "\t"
			} else {
				fmt.Fprintf(output, prependingString+lastLineSeparator+"%s\n", file.Name())
				stringPrepending = prependingString + "\t"
			}
			newDir := filepath.Join(currDir, file.Name())
			recursionPrintServise(stringPrepending, output, newDir, printFiles)
		} else if printFiles {
			if file.Size() > 0 {
				if length > i+1 {
					fmt.Fprintf(output, prependingString+firstLineSeparator+"%s (%vb)\n", file.Name(), file.Size())
				} else {
					fmt.Fprintf(output, prependingString+lastLineSeparator+"%s (%vb)\n", file.Name(), file.Size())
				}
			} else {
				if length > i+1 {
					fmt.Fprintf(output, prependingString+firstLineSeparator+"%s (empty)\n", file.Name())
				} else {
					fmt.Fprintf(output, prependingString+lastLineSeparator+"%s (empty)\n", file.Name())
				}
			}
		}
	}
}
