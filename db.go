package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
    "fmt"
)


type Impl struct {
	DB *gorm.DB
}


type Auteur struct {
  ID uint 						`gorm:"primary_key"`
  Pseudo string  				`gorm:"column:pseudo"`
  Created_at time.Time  		`gorm:"column:created_at"`
  Updated_at time.Time 			`gorm:"column:updated_at"`
  Cheked_profil uint			`gorm:"column:cheked_profil"`
  Pays string					`gorm:"column:pays"`
  Nb_messages uint				`gorm:"column:nb_messages"`
  Img_lien string				`gorm:"column:img_lien"`
  Nb_relation uint				`gorm:"column:nb_relation"`
  Banni uint					`gorm:"column:banni"`
  Date_inscription time.Time 	`gorm:"column:date_inscription"`
  Coord_X float64 				`gorm:"column:coord_X"`
  Coord_Y float64 				`gorm:"column:coord_Y"`
}


type Sujet struct {
  ID uint 						`gorm:"primary_key"`
  Created_at time.Time  		`gorm:"column:created_at"`
  Updated_at time.Time 			`gorm:"column:updated_at"`
  Parcoured uint				`gorm:"column:parcoured"`
  Url string					`gorm:"column:url"`
  Title string					`gorm:"column:title"`
  Auteur uint					`gorm:"column:auteur"`
  Nb_reponses uint				`gorm:"column:nb_reponses"`
  Initialised_at time.Time		`gorm:"column:initialised_at"`
}


func (i *Impl) InitDB() {
	var err error
	i.DB, err = gorm.Open("mysql", "root:root@/scrapping?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Got error when connect database, the error is '%v'")
	}
}


func (i *Impl) InitSchema() {
	i.DB.AutoMigrate(&Auteur{}, &Sujet{})
}


func (i *Impl) Close() {
	fmt.Println("Fermeture")
	i.DB.Close()
}


func (i *Impl) GetAllPseudo() []string {
	var names []string
	i.DB.Model(&Auteur{}).Pluck("pseudo", &names)
	fmt.Println(len(names))

	var final []string
	for _,v := range names {
		if(v != "Pseudosupprimé") {
			final = append(final, v)
		}
	}

	return final	
}

func (i *Impl) GetNbAuteur() int {
	var a int
	i.DB.Table("auteurs").Count(&a)
	return a
}


func (i *Impl) GetAuteur(id int) Auteur {
	var a Auteur
	i.DB.Find(&a, id)
	return a
}

func (i *Impl) GetAuteurs(lim int, offset int) []Auteur {
	var a []Auteur
	i.DB.Limit(lim).Offset(offset).Where("pseudo != ?", "Pseudosupprimé").Order("pseudo").Find(&a)
	return a
}


func (i *Impl) GetAuteurByPseudo(id string) Auteur {
	var a Auteur
	i.DB.Where("pseudo = ?", id).First(&a)
	return a
}


func (i *Impl) GetSujetByAuteur(id int) []Sujet {
	var a []Sujet
	i.DB.Where("auteur = ?", id).Find(&a)
	return a
}


func (i *Impl) updateAuteurCoordinate(id string, x float64, y float64) {
	a := i.GetAuteurByPseudo(id)
	if a.ID != 0 {
		i.DB.Model(&a).Updates(map[string]interface{}{"coord_X": x, "coord_Y": y})
	}
}