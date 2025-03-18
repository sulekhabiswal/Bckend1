package IntermediateServices

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/gofiber/fiber/v2"
)

// InsertToBigQuery inserts the filtered data into the specified BigQuery table
func InsertToBigQuery(filteredData map[string]interface{}, tableName string, bookMark string) (map[string]interface{}, error) {
	ctx := context.Background()

	// Initialize BigQuery client
	client, err := bigquery.NewClient(ctx, "BigQueryProjectID") 
	if err != nil {
		fmt.Println("ERROR:", bookMark, "Failed to create BigQuery client:", err)
		return map[string]interface{}{
			"status":       1,
			"message":      "BigQuery client initialization failed",
			"data":         nil,
			"errorMessage": err.Error(),
		}, err
	}
	defer client.Close()

	// Select dataset and table
	table := client.Dataset("BigQueryDataset").Table(tableName)

	// Prepare data for insertion
	inserter := table.Inserter()
	rows := []map[string]interface{}{filteredData} // Convert filteredData into a slice of maps

	// Insert data
	err = inserter.Put(ctx, rows)
	if err != nil {
		fmt.Println("ERROR:", bookMark, "BigQuery insertion failed:", err)
		return map[string]interface{}{
			"status":       1,
			"message":      "Catch block error. Insertion failed",
			"data":         nil,
			"errorMessage": err.Error(),
		}, err
	}

	// Success response
	fmt.Println("INFO:", bookMark, "Successfully inserted into", tableName)
	return map[string]interface{}{
		"status":       0,
		"message":      fmt.Sprintf("Successfully inserted into %s.", tableName),
		"data":         filteredData,
		"errorMessage": "",
	}, nil

	
}

func InsertIntoBigquery(c *fiber.Ctx, data map[string]interface{}, messageId string) (string, error) {
	recMessage, messageExists := data["message"].(map[string]interface{})
	if !messageExists {
		fmt.Println("ERROR:", messageId, "Invalid message format.")
		return "", fmt.Errorf("invalid message format")
	}

	var bookMark string

	column1Str, col1Exists := recMessage["id"].(string)
	if !col1Exists {
		if column1Num, ok := recMessage["id"].(json.Number); ok {
			column1Str = column1Num.String()
			col1Exists = true
		}
	}

	if col1Exists {
		bookMark = column1Str
	} else {
         fmt.Println("Id column is not exist")
	}

	fmt.Println("Value of bookmark:", bookMark)
	return bookMark, nil
}


func InsertIntoBigqueryTwo(c *fiber.Ctx, data map[string]interface{}, messageId string) (string, error) {
	recMessage, messageExists := data["message"].(map[string]interface{})
	if !messageExists {
		fmt.Println("ERROR:", messageId, "Invalid message format.")
		return "", fmt.Errorf("invalid message format")
	}

	var bookMark string

	column1Str, col1Exists := recMessage["user_id"].(string)
	if !col1Exists {
		if column1Num, ok := recMessage["user_id"].(json.Number); ok {
			column1Str = column1Num.String()
			col1Exists = true
		}
	}

	if col1Exists {
		bookMark = column1Str
	} else {
         fmt.Println("UserId column is not exist")
	}

	fmt.Println("Value of bookmark:", bookMark)
	return bookMark, nil
}
