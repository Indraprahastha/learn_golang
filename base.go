package main

import (
	"./config"
	controllers "./controller"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// ini adalah file yg ertama kali akan dijalankan

func main() {
	// mengarah ke file config, dengan function DBInit
	db := config.DBInit()
	// mengarah ke file grom
	inDB := &controllers.InDB{DB: db}
	// default dari librarry / framework
	router := gin.Default()

	// kumpulan routernya
	router.GET("/person/:id", inDB.GetPerson)
	router.GET("/persons", inDB.GetPersons)
	router.POST("/person", inDB.CreatePerson)
	router.PUT("/person", inDB.UpdatePerson)
	router.DELETE("/person/:id", inDB.DeletePerson)
	// jalan di 3000
	router.Run(":3000")
}
