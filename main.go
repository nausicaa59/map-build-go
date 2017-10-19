package main

import(
	"fmt"
)


func main(){
	fmt.Println("start")
    bdd := Impl{}
    bdd.InitDB()
    bdd.InitSchema()
    
    /*cleanInput()
    generateInput(bdd, 3)

    cleanOutput()
    Parse()
	Fusion()
	initialisation()
	Draw(0,2)*/

	//generateAllDim()
	actualisationPositions(bdd)
}

func generateAllDim(){
	for i := 1; i < 8; i++ {
		generateDim(i)
		Draw(i,2)
	}
}