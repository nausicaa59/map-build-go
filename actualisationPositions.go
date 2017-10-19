package main

import(
	"fmt"
	"time"
)

func actualisationPositions(bdd Impl) {
	start := time.Now()
	ouputFusionPath := "output/csv-fusion/fusion.csv"
	cerles := ReadCsvCircles(ouputFusionPath)
    
    nbPseudo := len(cerles)
    for i,v := range cerles {
    	x := (v.C_x0 + v.C_x1) / 2
    	y := (v.C_y0 + v.C_y1) / 2
    	pourc := (float64(i) / float64(nbPseudo)) * 100
    	pourcInt := int(pourc)
    	bdd.updateAuteurCoordinate(v.Id, x, y);
    	fmt.Println(pourcInt,v.Id, x, y)
    }
	

    end := time.Now()
    fmt.Println("Duration Initialisation : ",end.Sub(start), len(cerles))	
}
