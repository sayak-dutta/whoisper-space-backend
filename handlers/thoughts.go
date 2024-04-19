package handlers

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

type Thought struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	// Add any other fields you need
}

func CreateThought(svc *dynamodb.DynamoDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var thought Thought
		if err := c.ShouldBindJSON(&thought); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Save thought to DynamoDB
		input := &dynamodb.PutItemInput{
			Item: map[string]*dynamodb.AttributeValue{
				"id":      {S: aws.String(thought.ID)},
				"content": {S: aws.String(thought.Content)},
				// Add any other fields
			},
			TableName: aws.String("thoughts"), // Replace with your DynamoDB table name
		}

		_, err := svc.PutItem(input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, thought)
	}
}

func ListThoughts(svc *dynamodb.DynamoDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Scan DynamoDB table for all thoughts
		input := &dynamodb.ScanInput{
			TableName: aws.String("thoughts"), // Replace with your DynamoDB table name
		}

		result, err := svc.Scan(input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		thoughts := []Thought{}
		err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &thoughts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, thoughts)
	}
}
