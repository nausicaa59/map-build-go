package main

import(
	"fmt"
	"time"
)

func getExtremite(c []CercleFinal) (float64, float64) {
	X1 := 0.0
	Y1 := 0.0

	for _,v := range c {
		if v.C_x1 > X1 {
			X1 = v.C_x1
		}

		if v.C_y1 > Y1 {
			Y1 = v.C_y1
		}
	}

	return X1, Y1
}


func defineNbLigneColonne(c []CercleFinal, cPixel int) (int,int) {
	x, y := getExtremite(c)
	nbLigne := y / float64(cPixel)
	nbColonne := x / float64(cPixel)
	return int(nbLigne)+1, int(nbColonne)+1
}


func initialisation() {
	pixelBySection := 256
	start := time.Now()
	ouputFusionPath := "output/csv-fusion/fusion.csv"
	cerles := ReadCsvCircles(ouputFusionPath)
    
	ligne, colonne := defineNbLigneColonne(cerles, pixelBySection)
	for l := 0; l < ligne; l++ {
		for c := 0; c < colonne; c++ {
			section := Section{}
			section.x0 = float64(c) * float64(pixelBySection)
			section.x1 = section.x0 + float64(pixelBySection)
			section.y0 = float64(l) * float64(pixelBySection)
			section.y1 = section.y0 + float64(pixelBySection)
			pathSection := generatePathSection(0, l, c)
			selections := GetCerclesInSection(cerles, section)
			WriteCsvOutput(pathSection, selections)			
		}
	}

    end := time.Now()
    fmt.Println("Duration Initialisation : ",end.Sub(start), len(cerles))	
}