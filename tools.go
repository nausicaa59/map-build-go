package main

import(
	"os"
	"github.com/gocarina/gocsv"
	"fmt"
	"path/filepath"
)

func WriteCsvOutput(path string, p []CercleFinal) {
	clientsFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	err = gocsv.MarshalFile(&p, clientsFile)
	if err != nil {
		panic(err)
	}
}


func ReadCsvCircles(path string) []CercleFinal {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Erreur lecture")
		panic(err)
	}
	defer file.Close()

	
	var cercles []CercleFinal
	if err := gocsv.UnmarshalFile(file, &cercles); err != nil {
		fmt.Println("Erreur parse")
		panic(err)
	}

	return cercles
}

func cleanOutput() {
	//clean partial
	ouputPartial := "output/csv-partial/"
    os.RemoveAll(ouputPartial)

    if _, err := os.Stat(ouputPartial); os.IsNotExist(err) {
	    os.Mkdir(ouputPartial, os.ModePerm)
	}

	//clean partial
	ouputFusion := "output/csv-fusion/"
    os.RemoveAll(ouputFusion)

    if _, err := os.Stat(ouputFusion); os.IsNotExist(err) {
	    os.Mkdir(ouputFusion, os.ModePerm)
	}

	//clean partial
	ouputMap := "output/csv-map/"
    os.RemoveAll(ouputMap)

    if _, err := os.Stat(ouputMap); os.IsNotExist(err) {
	    os.Mkdir(ouputMap, os.ModePerm)
	}
}


func searchAllFiles(path string) []string {
  	fileList := []string{}

	if _, err := os.Stat(path); os.IsNotExist(err) {
	    return fileList
	}

	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		fileInfo, err := os.Stat(path)
		if(!fileInfo.IsDir()) {
			fileList = append(fileList, path)
		}
		return nil
	})

  	return fileList 
}