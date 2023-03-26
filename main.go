package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tniita/play-with-go/db"
	"github.com/tniita/play-with-go/http"
)

func main() {
	db.SetupDB()
	router := gin.Default()
	router.GET("/items", http.GetAll)
	router.GET("/items/:id", http.GetItemById)
	router.POST("/items", http.UpdateItem)
	router.DELETE("/items/:id", http.DeleteItem)
	router.Run()
}
