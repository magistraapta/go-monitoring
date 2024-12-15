package main

import (
	"log"
	"net/http"
	"time"

	"go-backend/database"

	user "go-backend/user"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"math/rand"
)

var (
	counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "devbulls_counter",
		Help: "Counting the total number of requets being handled",
	})

	gauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "devbulls_gauge",
		Help: "Monitoring node usage",
	}, []string{"node", "namespace"})
)

func RecordMetrics() {
	go func() {
		for {
			counter.Inc()
			gauge.WithLabelValues("node-1", "namespace-b").Set(rand.Float64())
			time.Sleep(time.Second * 5)
		}
	}()
}

func main() {
	RecordMetrics()

	// Initialize db connection
	_, err := database.GetDb() // GetDb() sets the global database.DB connection
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}

	defer database.DB.Close() // Use database.DB here

	// Setup Gin
	router := gin.Default()

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Request success"})
	})

	router.GET("/users", getUsers)
	router.POST("/users", addUser)

	router.Run()
}

func getUsers(ctx *gin.Context) {

	var users []user.UserResponse
	// Check if the database connection is initialized
	if database.DB == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}

	// select all users query
	query := `SELECT username, email FROM users` // Use 'users' if it's a plural table name

	// Execute the query to insert the user
	err := database.DB.Select(&users, query)
	if err != nil {
		// Return the actual error message
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query users: " + err.Error()})
		return
	}

	// Return response with the list of users
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully show all users",
		"user":    users,
	})
}

func addUser(ctx *gin.Context) {
	var newUser user.User

	// Check binding model to JSON process
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the database connection is initialized
	if database.DB == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}

	// Create user query
	query := `INSERT INTO users (username, email, password) VALUES (:username, :email, :password)` // Use 'users' if it's a plural table name

	// Execute the query to insert the user
	_, err := database.DB.NamedExec(query, newUser)
	if err != nil {
		// Return the actual error message
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	// Return response with the created user details
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully Created User",
		"user": gin.H{
			"username": newUser.Username,
			"email":    newUser.Email,
		},
	})
}
