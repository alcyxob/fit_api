package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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

// A global variable that will hold a reference to the MongoDB client
var mongoClient *mongo.Client

// The init function will run before our main function to establish a connection to MongoDB. If it cannot connect it will fail and the program will exit.
func init() {
	if err := connect_to_mongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}

// Our implementation logic for connecting to MongoDB
func connect_to_mongodb() error {
	uri := os.Getenv("MONGODB_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client
	return err
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.GET("/excercises", getExcercises)
	router.GET("/excercises/:id", getExcersiseByID)
	router.POST("/excercises", postExcercise)

	router.Run("localhost:8080")
}

// getExcercises responds with the list of all albums as JSON.
func getExcercises(c *gin.Context) {
	// Find movies
	cursor, err := mongoClient.Database("fitnnes").Collection("excercises").Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Map results
	var excercises []bson.M
	if err = cursor.All(context.TODO(), &excercises); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return movies
	c.JSON(http.StatusOK, excercises)
}

// postExcercises adds an excersise from JSON received in the request body.
func postExcercise(c *gin.Context) {
	var newExcercise excercise

	// Call BindJSON to bind the received JSON to
	// newExcercise.
	if err := c.BindJSON(&newExcercise); err != nil {
		return
	}

	coll := mongoClient.Database("fitnnes").Collection("excercises")
	result, err := coll.InsertOne(context.TODO(), newExcercise)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusCreated, result)
}

// getExcersiseByID locates the excercise whose ID value matches the id
// parameter sent by the client, then returns that excercise as a response.
func getExcersiseByID(c *gin.Context) {
	// Get excercise ID from URL
	id := c.Param("id")
	// Find excercise by ID
	var excercise bson.M
	err := mongoClient.Database("fitnnes").Collection("excercises").FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&excercise)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return excercise
	c.JSON(http.StatusOK, excercise)
}
