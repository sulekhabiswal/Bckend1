package IntermediateServices

import (
	"fmt"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)


func InsertIntoPostgre(c *fiber.Ctx, data map[string]interface{}, messageId string) (string, error) {
	recMessage, messageExists := data["message"].(map[string]interface{})
	if !messageExists {
		fmt.Println("ERROR:", messageId, "Invalid message format.")
		return "", fmt.Errorf("invalid message format")
	}

	
	var bookMark string

	
	column1Str, col1Exists := recMessage["column1"].(string)
	if !col1Exists {
		if column1Num, ok := recMessage["column1"].(json.Number); ok {
			column1Str = column1Num.String() 
			col1Exists = true
		}
	}

	
	if col1Exists {
		bookMark = column1Str
	} else if column40, ok := recMessage["column40"].([]interface{}); ok && len(column40) > 0 {
		if col40Map, isMap := column40[0].(map[string]interface{}); isMap {
			if col40Message, exists := col40Map["message"].(map[string]interface{}); exists {
				if column40Num, exists := col40Message["column2"].(json.Number); exists {
					bookMark = column40Num.String()
				}
			}
		}
	}

	fmt.Println("Value of bookmark:", bookMark)
	return bookMark, nil
}
