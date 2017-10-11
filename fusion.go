package main

import(
	"math"
	"io/ioutil"
	"log"
	"fmt"
	"time"
)

type PartialFile struct {
	path string
	xCorrection int
	yCorrection int
}

func Correction(path string, xCorrection int , yCorrection int) []CercleFinal {
	cercles := ReadCsvCircles(path)
	for i,_ := range cercles {
		cercles[i].X    	+= float64(xCorrection)
		cercles[i].Y    	+= float64(yCorrection)
		cercles[i].C_x0 	+= float64(xCorrection)
		cercles[i].C_y0 	+= float64(yCorrection)
		cercles[i].C_x1 	+= float64(xCorrection)
		cercles[i].C_y1 	+= float64(yCorrection)
	}

	return cercles	
}

func CarreApproximation(l int) int {
	sq := math.Sqrt(float64(l))
	a := math.Floor(sq)
	r := sq - a
	
	if r > 0 {
		return int(a) + 1
	}
	
	return int(a)
}


func Fusion() {
	//init
	start := time.Now()
	ouputDir := "output/csv-partial/"
	ouputFusionPath := "output/csv-fusion/fusion.csv"

	//read directory
    files, err := ioutil.ReadDir(ouputDir)
    if err != nil {
        log.Fatal(err)
    }

    //determination carre
    nbFile := len(files)
    c := CarreApproximation(nbFile)
    compteur := 0

    //determine correction
    var partials []PartialFile
    for i := 0; i < c; i++ {
    	for j := 0; j < c; j++ {
    		if compteur < nbFile {
    			tempo := PartialFile{}
    			tempo.path = ouputDir + files[compteur].Name()
    			tempo.xCorrection = j*2000
    			tempo.yCorrection = i*2000
    			partials = append(partials, tempo)    			
    		} 
    		compteur += 1   		
    	}
    }

    //apply correction
    var final []CercleFinal
    for _,p := range partials {
    	cercles := Correction(p.path, p.xCorrection, p.yCorrection)
    	for _,c := range cercles {
    		final = append(final, c)
    	}    	 
    }

    WriteCsvOutput(ouputFusionPath, final)

    
    end := time.Now()
    fmt.Println("Duration fusion : ",end.Sub(start))	
}