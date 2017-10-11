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
    generateInput(bdd, 3)*/

    cleanOutput()
    Parse()
	Fusion()
	initialisation()
	Draw()
}