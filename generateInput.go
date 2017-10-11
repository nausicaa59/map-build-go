package main

import(
	"fmt"
	"time"
	"strconv"
	"os/exec"
	"os"
)

type SegmentsCercles struct {
	start int
	nb int
}

func writeFile(path string, auteurs []Auteur) {
	var file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Erreur de l'ouverture de ", path, err)
		panic("arret")
	}

	defer file.Close()

	_, err = file.WriteString("id,value\n")
	if err != nil {
		fmt.Println("Erreur lors de l'écriture des headers ", path, err)
		panic("arret")
	}


	for _,a := range auteurs {
		_, err = file.WriteString(a.Pseudo + "," + strconv.Itoa(int(a.Nb_messages)) + "\n")
		if err != nil {
			fmt.Println("Erreur lors de l'écriture sur ", path, err)
			panic("arret")
		}		
	}

	// save changes
	err = file.Sync()
	if err != nil {
		fmt.Println("Erreur lors de la synchronisation de l'écriture sur ", path, err)
		panic("arret")
	}
}


func generateListe(bdd Impl, segments []SegmentsCercles) {
	for i,v := range segments {
		auteurs := bdd.GetAuteurs(v.nb, v.start)
		pathFile := "input/listePseudo/liste" + strconv.Itoa(i+1) +".txt"
		writeFile(pathFile, auteurs)
	}
}


func getSegments(bdd Impl, nbCercle int) []SegmentsCercles {
    var decoupes []SegmentsCercles
    nbAuteur := bdd.GetNbAuteur()
    reste := nbAuteur%nbCercle
    pseudoByCercle := int((nbAuteur - reste)/nbCercle)

    for i := 1; i <= nbCercle; i++ {
    	t := SegmentsCercles{}
    	t.start = (i-1)*pseudoByCercle
    	t.nb 	= pseudoByCercle
    	
    	if i == nbCercle {
    		t.nb += reste
    	}

    	decoupes = append(decoupes, t)
    }

    return decoupes
}

func generateInput(bdd Impl, nbCercle int) {
	start := time.Now()    
    segments := getSegments(bdd, nbCercle)
    generateListe(bdd, segments)

	cmd := exec.Command("python","input/script.py")
	err := cmd.Run()
  	if err != nil {
  		fmt.Println("Erreur !", err)
  	}


    end := time.Now()
    fmt.Println("Duration Generate Input : ",end.Sub(start))
}

