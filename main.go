package main

import (
	"fmt"
	"net/http"
	"os"
	usercontroller "startrail/controllers"
	"startrail/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// album represents data about a record album.
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

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {

	godotenv.Load()
	_, err := database.GetDB()
	if err != nil {
		fmt.Printf("Error while initializing database: %v", err)
		os.Exit(1)
	}

	router := gin.Default()
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", usercontroller.RegisterUser)
	}

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
