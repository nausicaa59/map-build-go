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

func scaleCircles(c []CercleFinal, fact int) {
	for i,_ := range c {
		c[i].X *= float64(fact)
		c[i].Y *= float64(fact)
		c[i].R *= float64(fact)
		c[i].C_x0 *= float64(fact)
		c[i].C_y0 *= float64(fact)
		c[i].C_x1 *= float64(fact)
		c[i].C_y1 *= float64(fact)
	}
}


func generateDim(dim int) {
	start := time.Now()
	oldDim := dim - 1
	oldDimPath := getPathDim(oldDim)
	oldFiles := searchAllFiles(oldDimPath)
	oldFilesInfo := searchAllInfo(oldFiles)
	pixelBySection := 256
	
	for _,f := range oldFilesInfo {
		cercles := ReadCsvCircles(f.path)
		if cercles == nil {
			continue
		}

		scaleCircles(cercles, 2)		
		for l := 0; l < 2; l++ {
			for c := 0; c < 2; c++ {
				ligne := (f.ligne * 2) + l
				colonne := (f.colonne * 2) + c
				pathSection := generatePathSection(dim, ligne, colonne)				
				section := Section{}
				section.x0 = float64(colonne) * float64(pixelBySection)
				section.x1 = section.x0 + float64(pixelBySection)
				section.y0 = float64(ligne) * float64(pixelBySection)
				section.y1 = section.y0 + float64(pixelBySection)
				selections := GetCerclesInSection(cercles, section)
				WriteCsvOutput(pathSection, selections)
			}			
		}
	}

    end := time.Now()
    fmt.Println("Duration build dim : ",end.Sub(start))
}