package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Load the configuration
	if err := loadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	// Set up the Gin router
	router := gin.Default()

	// Define the endpoints
	router.GET("/hello", handleHello)
	router.GET("/new", handleNewEndpoint)

	// Start the server
	port := "127.0.0.1:8080"
	fmt.Printf("Server is running on %s\n", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start the server: %s", err)
	}
}

func loadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func handleHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func handleNewEndpoint(c *gin.Context) {
	// Check if the "new_endpoint" feature is enabled
	if viper.GetBool("new_endpoint") {
		c.JSON(200, gin.H{
			"message": "This is the new endpoint!",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "This feature is not available.",
		})
	}
}
