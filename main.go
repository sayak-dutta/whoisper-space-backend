package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/sayak-dutta/whoisper-space-backend/handlers"
)

func main() {
	// Set up AWS DynamoDB connection
	creds := credentials.NewEnvCredentials()
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: creds,
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}
	svc := dynamodb.New(sess)

	r := gin.Default()

	// Set up routes
	r.GET("/thoughts", handlers.ListThoughts(svc))
	r.POST("/thoughts", handlers.CreateThought(svc))

	// Add routes for other endpoints

	r.Run() // listen and serve on 0.0.0.0:8080
}
