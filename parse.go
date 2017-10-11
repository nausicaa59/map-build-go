package main

import(
	"fmt"
	"os"
	"encoding/xml"
	"io/ioutil"
	"strings"
	"strconv"
	"log"
	"time"
)


type CercleFinal struct {
	Id 		string
	X 		float64
	Y 		float64
	R 		float64
	C1 		int
	C2 		int
	C3 		int
	C_x0 	float64
	C_y0 	float64
	C_x1 	float64
	C_y1 	float64
	I1 		int
	I2 		int
	I3 		int
}


type Svg struct {
	XMLName xml.Name `xml:"svg"`
	G   []G   `xml:"g"`
}


type G struct {
	XMLName xml.Name 	 `xml:"g"`
	Class    	string   `xml:"class,attr"`
	Transform   string   `xml:"transform,attr"`
	Circle 		Circle 	 `xml:"circle"`
	ClipPath 	string 	 `xml:"clipPath"`
	Title 		string 	 `xml:"title"`
}


type Circle struct {
	XMLName 	xml.Name `xml:"circle"`
	Id    		string   `xml:"id,attr"`
	R   		string   `xml:"r,attr"`
	Style    	string   `xml:"style,attr"`
}


func ParseTranslate(t string) (float64, float64) {
	clean := strings.Replace(t, "translate(", "", -1)
	clean = strings.Replace(clean, ")", "", -1)
	fragments := strings.Split(clean, ",")
	x, _ := strconv.ParseFloat(fragments[0], 64)
	y, _ := strconv.ParseFloat(fragments[1], 64)
	return x,y
}


func ParseColor(s string) (int, int, int) {
	clean := strings.Replace(s, "fill: rgb(", "", -1)
	clean  = strings.Replace(clean, ");", "", -1)
	clean  = strings.Replace(clean, " ", "", -1)
	fragments := strings.Split(clean, ",")
	c1, _ := strconv.Atoi(fragments[0])
	c2, _ := strconv.Atoi(fragments[1])
	c3, _ := strconv.Atoi(fragments[2])
	return c1, c2, c3
} 


func ParseCercle(g G, imgM *ImageMatrix) CercleFinal {
	tempo := CercleFinal{}
	tempo.Id = g.Circle.Id
	tempo.X, tempo.Y = ParseTranslate(g.Transform)
	tempo.R, _ = strconv.ParseFloat(g.Circle.R, 64)
	tempo.C1, tempo.C2, tempo.C3 = ParseColor(g.Circle.Style)
	tempo.I1, tempo.I2, tempo.I3 = imgM.getColorByCoord(int(tempo.X), int(tempo.Y))
	tempo.C_x0 = tempo.X - tempo.R
	tempo.C_y0 = tempo.Y - tempo.R
	tempo.C_x1 = tempo.X + tempo.R
	tempo.C_y1 = tempo.Y + tempo.R	
	return tempo
}


func ParseFile(pathSVG string, pathImg string) []CercleFinal {
	xmlFile, err := os.Open(pathSVG)
	if err != nil {
		panic("Error opening file:")
	}

	defer xmlFile.Close()
	var svg Svg
	var final []CercleFinal
	img := getImageMatrix(pathImg)
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &svg)

	for i := 0; i < len(svg.G); i++ {
		t := ParseCercle(svg.G[i], &img)
		final = append(final, t)
	}

	return final	
}



func Parse() {
	start := time.Now()
	pathDir := "input/svg/"
	ouputDir := "output/csv-partial/"

    files, err := ioutil.ReadDir(pathDir)
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
    	filePath := pathDir + f.Name()
    	prefixe := strings.Replace(f.Name(), ".xml", "", -1)
    	outputName := prefixe + ".csv"
    	outputPath := ouputDir + outputName
    	cercles := ParseFile(filePath, "input/images/image.png")
		WriteCsvOutput(outputPath, cercles)
    }    


    end := time.Now()
    fmt.Println("Duration parse :", end.Sub(start))	
}