package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"example.com/web-service-gin/data/albumsql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Errorf("Error loading .env file")
	}

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	albumsql.InitDb()

	router.Run("localhost:" + os.Getenv("API_PORT"))
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	albums, err := albumsql.AllAlbums()
	if err != nil {
		fmt.Errorf("getAlbums: %v", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	idstr := c.Param("id")

	inid, err := strconv.Atoi(idstr)

	if err != nil {
		fmt.Errorf("could not convert %v to int", idstr)
	}

	id64 := int64(inid)

	al, err := albumsql.AlbumById(id64)

	if err != nil {
		fmt.Errorf("not good %v", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}

	if al.ID != 0 {
		c.IndentedJSON(http.StatusOK, al)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum albumsql.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	id, err := albumsql.AddAlbum(newAlbum)
	if err != nil {
		fmt.Errorf("postAlbums: %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not add album"})
	}
	newAlbum.ID = id
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
