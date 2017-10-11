package main

import (
    "fmt"
    "image"
    "image/png"
    "os"
    "io"
)

type ImageMatrix struct {
	src string
	m [][]Pixel
}

type Pixel struct {
    R int
    G int
    B int
    A int
}


func getImageMatrix(src string) ImageMatrix {
    image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
    file, err := os.Open(src)
    if err != nil {
        fmt.Println("Error: File could not be opened")
        os.Exit(1)
    }

    defer file.Close()
    i := ImageMatrix{}
    i.src = src
    i.m, err = i.buildMapPixels(file)

    if err != nil {
        fmt.Println("Error: Image could not be decoded")
        os.Exit(1)
    }

    return i
}


func (i ImageMatrix) buildMapPixels(file io.Reader) ([][]Pixel, error) {
    img, _, err := image.Decode(file)
    if err != nil {
        return nil, err
    }

    bounds := img.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y
    var pixels [][]Pixel
    for y := 0; y < height; y++ {
        var row []Pixel
        for x := 0; x < width; x++ {
            row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
        }
        pixels = append(pixels, row)
    }

    return pixels, nil
}


func (i *ImageMatrix) getColorByCoord(x int, y int) (int,int,int) {
	R := 35
	G := 78
	B := 154


	if(y < 0 || y > len(i.m) - 1) {
		return R,G,B
	}

	if(x < 0 || x > len(i.m[y]) - 1) {
		return R,G,B
	}

	p := i.m[y][x]
    if(p.A == 0) {
        return 35,78,154
    }

	return p.R,p.G,p.B
}


func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
    return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}