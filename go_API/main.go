package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}

}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
