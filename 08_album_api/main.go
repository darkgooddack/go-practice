package main

import (
	"hello-go/08_album_api/config"
	"hello-go/08_album_api/db"
	"hello-go/08_album_api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	cfg := config.LoadConfig()
	db.InitDB(cfg)

	// Автоматическая миграция
	if err := db.DB.AutoMigrate(&models.Album{}); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	err := router.Run("localhost:8000")
	if err != nil {
		return
	}
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// как импортировать этот Album из ./models/album.go
	var albums []models.Album
	db.DB.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Генерируем UUID для нового альбома
	newAlbum.ID = uuid.New().String()

	db.DB.Create(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id parameter sent by the client
func getAlbumByID(c *gin.Context) {
	var album models.Album
	if err := db.DB.First(&album, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Альбом не найден"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}
