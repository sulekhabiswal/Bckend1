package IntermediateServices

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// InsertService handles insertion logic for different product types
func InsertServiceBigquery(c *fiber.Ctx, data map[string]interface{}, bookMark string) error {
	recMessage, recExists := data["message"].(map[string]interface{})
	if !recExists {
		fmt.Println("ERROR:", bookMark, "Missing 'message' in request data")
		return errors.New("missing 'message' in request data")
	}

	productName, productExists := data["key"].(string)
	if !productExists {
		fmt.Println("ERROR:", bookMark, "Missing 'key' in request data")
		return errors.New("missing 'key' in request data")
	}

	switch productName {
	case "PREPAIDCARD_TRANSACTION":
		filteredData, err := FilterDataBigquery(recMessage, bookMark)
		if err != nil {
			fmt.Println("ERROR:", bookMark, "Error in data filtering:", err)
			return errors.New("error in data filtering")
		}

		transactionCreatedDate, ok := filteredData["created_date"].(string)
		if !ok || transactionCreatedDate == "" {
			fmt.Println("ERROR:", bookMark, "Missing 'created_date' in filtered data")
			return errors.New("missing 'created_date' in filtered data")
		}

		tableName, err := GetTableName(transactionCreatedDate, bookMark)
		if err != nil {
			fmt.Println("ERROR:", bookMark, "Error in creating table name:", err)
			return errors.New("error in creating table name")
		}

		_, err = InsertToBigQuery(filteredData, tableName, bookMark)
		if err != nil {
			fmt.Println("ERROR:", bookMark, "Error in BigQuery insertion:", err)
			return errors.New("error in BigQuery insertion")
		}

		fmt.Println("INFO:", bookMark, "Successfully inserted into BigQuery")
		return SendResponse(c, true, "Successfully inserted.")

	case "PREPAIDCARD_USER":
		userID, userExists := recMessage["user_id"].(string)
		if !userExists || userID == "" {
			fmt.Println("ERROR:", bookMark, "Missing input parameter(user_id)")
			return errors.New("missing input parameter(user_id)")
		}

		filteredData, err := FilterDataForUserCreation(recMessage, bookMark)
		if err != nil {
			fmt.Println("ERROR:", bookMark, "Error in data filtering:", err)
			return errors.New("error in data filtering")
		}

		_,err = InsertToBigQuery(filteredData, "usercreation_bigqueryTableName", bookMark)
		if err != nil {
			fmt.Println("ERROR:", bookMark, "Error in BigQuery insertion:", err)
			return errors.New("error in BigQuery insertion")
		}

		fmt.Println("INFO:", bookMark, "Successfully inserted into BigQuery")
		return SendResponse(c, true, "Successfully inserted.")

	default:
		fmt.Println("ERROR:", bookMark, "Invalid Product name:", productName)
		return errors.New("invalid Product name")
	}
}
