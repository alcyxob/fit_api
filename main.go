package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// album represents data about a record album.
type excercise struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Muscles string `json:"muscle"`
	Type    string `json:"type"`
}

// albums slice to seed record album data.
var excercises = []excercise{
	{ID: "1", Title: "Horisontall Pull", Muscles: "Back", Type: "Pull"},
	{ID: "2", Title: "Push ups", Muscles: "Chest", Type: "Push"},
	{ID: "3", Title: "Pull ups", Muscles: "Back", Type: "Pull"},
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()
	router.GET("/excercises", getExcercises)
	router.GET("/excercises/:id", getExcersiseByID)
	router.POST("/excercises", postExcercise)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getExcercises(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, excercises)
}

// postExcercises adds an excersise from JSON received in the request body.
func postExcercise(c *gin.Context) {
	var newExcercise excercise

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newExcercise); err != nil {
		return
	}

	// Add the new album to the slice.
	excercises = append(excercises, newExcercise)
	c.IndentedJSON(http.StatusCreated, newExcercise)
}

// getExcersiseByID locates the excercise whose ID value matches the id
// parameter sent by the client, then returns that excercise as a response.
func getExcersiseByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range excercises {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "excercise not found"})
}
