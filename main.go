package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sawasaki-narumi/calcal-api/routes"
	"github.com/sawasaki-narumi/calcal-api/utils"
	"log"

	"github.com/sawasaki-narumi/calcal-api/models"
)

func main() {
	db, err := utils.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{})

	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	router := gin.Default()
	api := router.Group("/api/v1")
	routes.InitializeRoutes(api)
	router.Run()
}
