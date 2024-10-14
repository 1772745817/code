package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getalbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postalbum(c *gin.Context) {
	var newalbums album
	if err := c.BindJSON(&newalbums); err != nil {
		return
	}
	albums = append(albums, newalbums)
	c.IndentedJSON(http.StatusCreated, newalbums)

}

func getalbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	//gin.H : map[string]interface{}别名 键为massage,值为“album not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album not found"})
}

func main() {
	route := gin.Default()

	route.GET("/albums", getalbum)

	route.POST("/albums", postalbum)

	route.GET("/albums/:id", getalbumById)
	route.Run("localhost:8080")

}
