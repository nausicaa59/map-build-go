package main

import (
  "fmt"
  "time"
  "github.com/fogleman/gg"
  "path/filepath"
  "os"
)


func placeText(c CercleFinal, dc *gg.Context, X float64, Y float64, forced bool) {
    nbCaract := len(c.Id)
    largeurCaract := (c.C_x1 - c.C_x0)/float64(nbCaract)
    size := ((1.9755*largeurCaract) -0.0127)

    if(size >= 8 || forced) {
        if(forced && size < 18) {
            size = 18
        }

        if(size > 400) {
            size = 400
        }

        if err := dc.LoadFontFace("input/font/nasalization-rg.ttf", size); err != nil {
            panic(err)
        }

        width, height := dc.MeasureString(c.Id)
        dc.SetRGB255(255, 255, 255)
        dc.DrawString(c.Id, X - width/2, Y + height/2)
        dc.Fill()
    }
}


func traceCercle(f SectionFile, cercles []CercleFinal, quality int) {
    correctionX := f.colonne * 256 * quality
    correctionY := f.ligne * 256 * quality
    outputFile := generatePathImgOutput(f.dim, f.ligne, f.colonne)

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

        forceText := false
        if(f.dim > 5) {
            forceText = true
        }

        placeText(c, dc, X, Y, forceText)
    }
    
    dc.SavePNG(outputFile)
    dc = nil
}


func DrawFile(path string, quality int) {
    f := infoFile(path)
    cercles := ReadCsvCircles(f.path)
    
    if cercles != nil {
      traceCercle(f, cercles, quality)
      fmt.Println("Trace", path)
    } else {
      fmt.Println("Ignore")
    }
}


func Draw(dim int, quality int) {
    start := time.Now()
    searchDir := getPathDim(dim)
    
    filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
        fileInfo, err := os.Stat(path)
        if(!fileInfo.IsDir()) {
            DrawFile(path, quality)
        }

        return nil
    })

    end := time.Now()
    fmt.Println("Duration Draw : ",end.Sub(start))
}