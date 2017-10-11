package main

import (
  "fmt"
  "time"
  "github.com/fogleman/gg"
)


func placeText(c CercleFinal, dc *gg.Context, X float64, Y float64) {
  nbCaract := len(c.Id)
  largeurCaract := (c.C_x1 - c.C_x0)/float64(nbCaract)
  size := ((1.9755*largeurCaract) -0.0127)

  if(size >= 10) {
    if err := dc.LoadFontFace("input/font/nasalization-rg.ttf", size); err != nil {
      panic(err)
    }

    width, height := dc.MeasureString(c.Id)
    dc.SetRGB255(255, 255, 255)
    dc.DrawString(c.Id, X - width/2, Y + height/2)
    dc.Fill()
    fmt.Println("trace !")    
  }
}


func traceCercle(f SectionFile, cercles []CercleFinal, quality int) {
    correctionX := f.colonne * 256 * quality
    correctionY := f.ligne * 256 * quality
    outputFile := generatePathImgOutput(f.dim, f.ligne, f.colonne)
    fmt.Println(outputFile, correctionX, correctionY)
    
    dc := gg.NewContext(256*quality, 256*quality)
    dc.SetRGB255(0, 0, 0)
    dc.DrawRectangle(0,0,256*float64(quality),256*float64(quality))
    dc.Fill()

    for _, c := range cercles {
      X:= (c.X*float64(quality)) - float64(correctionX)
      Y:= (c.Y*float64(quality)) - float64(correctionY)
      dc.DrawCircle(X, Y, c.R*float64(quality))

      if(f.dim == 0) {
        dc.SetRGB255(c.I1, c.I2, c.I3)
      } else {
        dc.SetRGB255(c.C1, c.C2, c.C3)
      }
      
      dc.Fill()
      placeText(c, dc, X, Y)
    }
    
    dc.SavePNG(outputFile)
}


func Draw() {
    start := time.Now()
    searchDir := "output/csv-map/0"
    fileList := searchAllFiles(searchDir)
    fileInfos := searchAllInfo(fileList)

    for _,v := range fileInfos {
        cercles := ReadCsvCircles(v.path)
        traceCercle(v, cercles, 2)
    }

    end := time.Now()
    fmt.Println("Duration Draw : ",end.Sub(start))
}