package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album with the stucture which contains info
type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64  `json:"price"`
}
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	// Get the all data of albums 
    router.GET("/albums", getAlbums)
	// Get the all data of albums 
	router.GET("/albums/:id", getAlbumsByID)

	// Post the data to albums
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}

// Get the whole albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Get the albums by id
func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	// Looping through the albums to get the id for a particular id
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Content not found!"})
}

// Store the data in some var
func postAlbums(c *gin.Context) {
	var newAlbum Album;

	// Call bindJson to check the err msg
	if err := c.BindJSON(&newAlbum); err !=  nil {
		return
	}

	albums := append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, albums)
}
