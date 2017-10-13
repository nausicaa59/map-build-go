package main

import(
	"strconv"
	"os"
	"strings"
	"fmt"
)

type Section struct {
	x0 float64
	x1 float64
	y0 float64
	y1 float64
}

type SectionFile struct {
    dim int
    colonne int
    ligne int
    path string
}


func CercleInSection(c CercleFinal, sect Section) bool {
	scond_x1 := (c.C_x0 > sect.x0) && (c.C_x1 < sect.x1)
	scond_x2 := (c.C_x0 < sect.x0) && (c.C_x1 > sect.x1)
	scond_x3 := (c.C_x0 < sect.x0) && (c.C_x1 > sect.x0)
	scond_x4 := (c.C_x0 < sect.x1) && (c.C_x1 > sect.x1)	
	condition_x := scond_x1 || scond_x2 || scond_x3 || scond_x4	

	scond_y1 := (c.C_y0 > sect.y0) && (c.C_y1 < sect.y1)
	scond_y2 := (c.C_y0 < sect.y0) && (c.C_y1 > sect.y1)
	scond_y3 := (c.C_y0 < sect.y0) && (c.C_y1 > sect.y0)
	scond_y4 := (c.C_y0 < sect.y1) && (c.C_y1 > sect.y1)
	condition_y := scond_y1 || scond_y2 || scond_y3 || scond_y4	

	return condition_x && condition_y
}


func GetCerclesInSection(c []CercleFinal, sect Section)[]CercleFinal {
	var s []CercleFinal

	for _,v := range c {
		if CercleInSection(v, sect) {
			s = append(s, v)
		}
	}

	return s
}

func getPathDim(dim int) string {
	strDim 	:= strconv.Itoa(dim)
	return "output/csv-map/" + strDim + "/"
}

func generatePathSection(dim int, l int, c int) string {
	strDim 	:= strconv.Itoa(dim)
	strL 	:= strconv.Itoa(l)
	strC 	:= strconv.Itoa(c)
	folderDim := "output/csv-map/" + strDim + "/"
	folderLig := folderDim + strL + "/"

	if _, err := os.Stat(folderDim); os.IsNotExist(err) {
    	os.Mkdir(folderDim, os.ModePerm)
	}

	if _, err := os.Stat(folderLig); os.IsNotExist(err) {
    	os.Mkdir(folderLig, os.ModePerm)
	}

	return folderLig + strL + "-" + strC + ".csv"
}

func generatePathImgOutput(dim int, l int, c int) string {
	strDim 	:= strconv.Itoa(dim)
	strL 	:= strconv.Itoa(l)
	strC 	:= strconv.Itoa(c)
	folderDim := "output/img-map/" + strDim + "/"
	folderLig := folderDim + strL + "/"

	if _, err := os.Stat(folderDim); os.IsNotExist(err) {
    	os.Mkdir(folderDim, os.ModePerm)
	}

	if _, err := os.Stat(folderLig); os.IsNotExist(err) {
    	os.Mkdir(folderLig, os.ModePerm)
	}

	return folderLig + strL + "-" + strC + ".png"
}


func infoFile(path string) SectionFile {
  clean := strings.Replace(path, ".csv", "", -1)
  fragments := strings.Split(clean, "\\")
  fileFragment := strings.Split(fragments[len(fragments)-1], "-")

  valDim, err := strconv.Atoi(fragments[len(fragments)-3]);
  if(err !=  nil){
    fmt.Println("erreur sur la dimension !")
  }

  valLigne, err := strconv.Atoi(fileFragment[0]);
  if(err !=  nil){
    fmt.Println("erreur sur la ligne !")
  }

  valColonne, err := strconv.Atoi(fileFragment[1]);
  if(err !=  nil){
    fmt.Println("erreur sur la colonne !")
  }
  
  return SectionFile{
    dim: valDim,
    colonne: valColonne,
    ligne: valLigne,
    path: path}
}

func searchAllInfo(paths []string) []SectionFile {
  fileInfo := []SectionFile{}

  for _, file := range paths {
      info := infoFile(file)
      fileInfo = append(fileInfo, info)
  } 

  return fileInfo
}