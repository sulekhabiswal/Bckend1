// package IntermediateServices

// import (
// 	"fmt"
// 	"github.com/gofiber/fiber/v2"
// )

// // InsertService handles inserting or updating records in different tables
// func InsertService(c *fiber.Ctx, decodedData map[string]interface{}, bookMark string) error {
// 	recMessage, messageExists := decodedData["message"].(map[string]interface{})
// 	if !messageExists {
// 		fmt.Println("ERROR:", bookMark, "Invalid message format.")
// 		return fmt.Errorf("invalid message format")
// 	}

// 	// Step 1: Ensure Key == "PREPAIDCARD"
// 	productName, keyExists := decodedData["key"].(string)
// 	if !keyExists || productName != "PREPAIDCARD" {
// 		fmt.Println("ERROR:", bookMark, "Key must be PREPAIDCARD")
// 		return fmt.Errorf("key must be PREPAIDCARD")
// 	}

// 	// Step 2: Extract Operation (CREATE/UPDATE)
// 	operation, operationExists := recMessage["column36"].(string)
// 	if !operationExists {
// 		fmt.Println("ERROR:", bookMark, "Missing operation type (CREATE/UPDATE).")
// 		return fmt.Errorf("missing operation type")
// 	}

// 	fmt.Println("INFO:", bookMark, "Processing operation:", operation)

// 	// Step 3: Handle logic based on `column6` (Table Type)
// 	switch recMessage["column6"].(string) {
// 	case "customer":
// 		// Filter Customer Data
// 		filteredData, err := FilterCustomerData(recMessage, bookMark)
// 		if err != nil {
// 			return err
// 		}

// 		// Prepare Query
// 		prepareQueryObj := map[string]interface{}{
// 			"data":      filteredData,
// 			"operation": operation,
// 			"tableName": "customer",
// 		}
// 		postgreQuery, err := PrepareQuery(prepareQueryObj, bookMark)
// 		if err != nil {
// 			return err
// 		}

// 		// Execute Query
// 		_, err = RunPostgresQuery(postgreQuery, bookMark)
// 		if err != nil {
// 			return err
// 		}

// 		return nil

// 	case "card":
// 		// Filter Card Data
// 		filteredData, err := FilterCardData(recMessage, bookMark)
// 		if err != nil {
// 			return err
// 		}

// 		// Prepare Query
// 		prepareQueryObj := map[string]interface{}{
// 			"data":      filteredData,
// 			"operation": operation,
// 			"tableName": "card",
// 		}
// 		postgreQuery, err := PrepareQuery(prepareQueryObj, bookMark)
// 		if err != nil {
// 			return err
// 		}

// 		// Execute Query
// 		_, err = RunPostgresQuery(postgreQuery, bookMark)
// 		if err != nil {
// 			return err
// 		}

// 		// If MCC List Exists, Convert []interface{} to []map[string]interface{}
// 		rawMccList, ok := recMessage["column40"].([]interface{})
// 		if ok && len(rawMccList) > 0 {
// 			var cardMccList []map[string]interface{}
// 			for _, item := range rawMccList {
// 				if mccMap, valid := item.(map[string]interface{}); valid {
// 					cardMccList = append(cardMccList, mccMap)
// 				}
// 			}

// 			// Process MCC Data
// 			filteredMccData, err := FilterMccData(cardMccList, bookMark)
// 			if err != nil {
// 				return err
// 			}

// 			// Prepare Bulk Query
// 			prepareBulkQueryObj := map[string]interface{}{
// 				"keyValueList": filteredMccData["keyValueList"],
// 				"valueList":    filteredMccData["valueList"],
// 				"operation":    operation,
// 				"tableName":    "card_mcc",
// 			}

// 			postgreBulkQuery, err := PrepareBulkQuery(prepareBulkQueryObj, bookMark)
// 			if err != nil {
// 				return err
// 			}

// 			// Execute Bulk Query
// 			_, err = RunPostgresQuery(postgreBulkQuery, bookMark)
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		return nil

// 	default:
// 		fmt.Println("ERROR:", bookMark, "Invalid table type")
// 		return fmt.Errorf("invalid table type")
// 	}
// }

package IntermediateServices

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// InsertService handles inserting or updating records in different tables
func InsertService(c *fiber.Ctx, decodedData map[string]interface{}, bookMark string) error {
	recMessage, messageExists := decodedData["message"].(map[string]interface{})
	if !messageExists {
		fmt.Println("ERROR:", bookMark, "Invalid message format.")
		return errors.New("invalid message format")
	}

	// Step 1: Ensure Key == "PREPAIDCARD"
	productName, keyExists := decodedData["key"].(string)
	if !keyExists || productName != "PREPAIDCARD" {
		fmt.Println("ERROR:", bookMark, "Key must be PREPAIDCARD")
		return SendResponse(c, false, "Key must be PREPAIDCARD")
	}

	// Step 2: Extract Operation (CREATE/UPDATE)
	operation, operationExists := recMessage["column36"].(string)
	if !operationExists {
		fmt.Println("ERROR:", bookMark, "Missing operation type (CREATE/UPDATE).")
		return SendResponse(c, false, "Missing operation type")
	}

	fmt.Println("INFO:", bookMark, "Processing operation:", operation)

	switch recMessage["column6"].(string) {
	//customer case

	case "customer":
		// Check if user_id (column1) exists
		if _, userExists := recMessage["column1"]; !userExists {
			fmt.Println("ERROR:", bookMark, "user_id should not be empty.")
			return SendResponse(c, false, "user_id should not be empty.")
		}

		// Filter customer data
		filteredData, err := FilterCustomerData(recMessage, bookMark)
		if err != nil {
			fmt.Println("ERROR:", bookMark, "Error in data filtering.")
			return SendResponse(c, false, "Error in data filtering.")
		}

		// Prepare the query object
		prepareQueryObj := map[string]interface{}{
			"data":      filteredData,
			"operation": operation,
			"tableName": "customer",
		}

		// Prepare the SQL query
		postgreQuery, err := PrepareQuery(prepareQueryObj, bookMark)
		if err != nil {
			fmt.Println("ERROR:", bookMark, "Error in preparing query.")
			return SendResponse(c, false, "Error in preparing query.")
		}

		// Execute the query in PostgreSQL
		_, err = RunPostgresQuery(postgreQuery, bookMark)
		if err != nil {
			fmt.Println("ERROR:", bookMark, "Error executing query.")
			return SendResponse(c, false, "Error executing query.")
		}

		fmt.Println("INFO:", bookMark, "Customer operation executed successfully.")
		return SendResponse(c, true, "Customer operation executed successfully.")

	// Card Case only
	case "card":

		cardMccList, mccExists := recMessage["column40"].([]interface{})
		_, refExists := recMessage["column1"]

		//  CREATE Operation for Card
		if operation == "CREATE" {
			if !refExists || !mccExists {
				fmt.Println("ERROR:", bookMark, "card_refnumber or card MCC list is missing")
				return SendResponse(c, false, "card_refnumber or card MCC list is missing")
			}

			filteredData, err := FilterCardData(recMessage, bookMark)
			if err != nil {
				return SendResponse(c, false, "Error filtering card data")
			}

			prepareQueryObj := map[string]interface{}{
				"data":      filteredData,
				"operation": operation,
				"tableName": "card",
			}
			postgreQuery, err := PrepareQuery(prepareQueryObj, bookMark)
			if err != nil {
				return SendResponse(c, false, "Error preparing card query")
			}

			_, err = RunPostgresQuery(postgreQuery, bookMark)
			if err != nil {
				return SendResponse(c, false, "Error executing card query")
			}

			//  Process MCC Data if Exists
			if mccExists && len(cardMccList) > 0 {
				var formattedMccList []map[string]interface{}
				for _, item := range cardMccList {
					if mccMap, valid := item.(map[string]interface{}); valid {
						if message, exists := mccMap["message"].(map[string]interface{}); exists {
							formattedMccList = append(formattedMccList, message)
						}
					}
				}

				filteredMccData, err := FilterMccData(formattedMccList, bookMark)
				if err != nil {
					return SendResponse(c, false, "Error filtering MCC data")
				}

				prepareBulkQueryObj := map[string]interface{}{
					"keyValueList": filteredMccData["keyValueList"],
					"valueList":    filteredMccData["valueList"],
					"operation":    operation,
					"tableName":    "card_mcc",
				}

				postgreBulkQuery, err := PrepareBulkQuery(prepareBulkQueryObj, bookMark)
				if err != nil {
					return SendResponse(c, false, "Error preparing MCC bulk query")
				}

				_, err = RunPostgresQuery(postgreBulkQuery, bookMark)
				if err != nil {
					return SendResponse(c, false, "Error executing MCC bulk query")
				}
			}

			return SendResponse(c, true, "Card and MCC data processed successfully.")
		} else {

			operation = "UPDATE"

			//  If both cardRefNumber and cardMccList are missing, return an error
			if !mccExists && !refExists {
				fmt.Println("ERROR:", bookMark, "card Mcc list or card ref number is not present for card data.")
				return SendResponse(c, false, "card Mcc list or card ref number is missing")
			}

			// If only MCC List is missing, update only Card table
			if !mccExists {
				filteredCardData, err := FilterCardData(recMessage, bookMark)
				if err != nil {
					return SendResponse(c, false, "Error filtering card data")
				}

				prepareQueryObj := map[string]interface{}{
					"data":            filteredCardData,
					"operation":       operation,
					"tableName":       "card",
					"primaryKeyName":  "card_refnumber",
					"primaryKeyValue": recMessage["column1"],
				}
				postgreQuery, err := PrepareQuery(prepareQueryObj, bookMark)
				if err != nil {
					return SendResponse(c, false, "Error preparing card update query")
				}

				_, err = RunPostgresQuery(postgreQuery, bookMark)
				if err != nil {
					return SendResponse(c, false, "Error executing card update query")
				}

				return SendResponse(c, true, "Card data updated successfully.")
			}

			// If only CardRefNumber is missing, update only MCC table
			if !refExists {
				var formattedMccList []map[string]interface{}
				for _, item := range cardMccList {
					if mccMap, valid := item.(map[string]interface{}); valid {
						if message, exists := mccMap["message"].(map[string]interface{}); exists {
							formattedMccList = append(formattedMccList, message)
						}
					}
				}

				filteredCardMccList, err := FilterMccData(formattedMccList, bookMark)
				if err != nil {
					return SendResponse(c, false, "Error filtering MCC data")
				}

				prepareBulkQueryObj := map[string]interface{}{
					"keyValueList": filteredCardMccList["keyValueList"],
					"operation":    operation,
					"tableName":    "card_mcc",
				}

				postgreBulkQuery, err := PrepareBulkQuery(prepareBulkQueryObj, bookMark)
				if err != nil {
					return SendResponse(c, false, "Error preparing MCC update query")
				}

				_, err = RunPostgresQuery(postgreBulkQuery, bookMark)
				if err != nil {
					return SendResponse(c, false, "Error executing MCC update query")
				}

				return SendResponse(c, true, "MCC table updated successfully.")
			}

			// If both cardRefNumber & MCC List exist, update both tables
			filteredCardData, err := FilterCardData(recMessage, bookMark)
			if err != nil {
				return SendResponse(c, false, "Error filtering card data")
			}

			prepareQueryObj := map[string]interface{}{
				"data":            filteredCardData,
				"operation":       operation,
				"tableName":       "card",
				"primaryKeyName":  "card_refnumber",
				"primaryKeyValue": recMessage["column1"],
			}
			postgreQuery, err := PrepareQuery(prepareQueryObj, bookMark)
			if err != nil {
				return SendResponse(c, false, "Error preparing card update query")
			}

			_, err = RunPostgresQuery(postgreQuery, bookMark)
			if err != nil {
				return SendResponse(c, false, "Error executing card update query")
			}

			//  Update MCC table
			var formattedMccList []map[string]interface{}
			for _, item := range cardMccList {
				if mccMap, valid := item.(map[string]interface{}); valid {
					if message, exists := mccMap["message"].(map[string]interface{}); exists {
						formattedMccList = append(formattedMccList, message)
					}
				}
			}

			filteredCardMccList, err := FilterMccData(formattedMccList, bookMark)
			if err != nil {
				return SendResponse(c, false, "Error filtering MCC data")
			}

			prepareBulkQueryObj := map[string]interface{}{
				"keyValueList": filteredCardMccList["keyValueList"],
				"operation":    operation,
				"tableName":    "card_mcc",
			}

			postgreBulkQuery, err := PrepareBulkQuery(prepareBulkQueryObj, bookMark)
			if err != nil {
				return SendResponse(c, false, "Error preparing MCC update query")
			}

			_, err = RunPostgresQuery(postgreBulkQuery, bookMark)
			if err != nil {
				return SendResponse(c, false, "Error executing MCC update query")
			}

			return SendResponse(c, true, "Card and MCC data updated successfully.")
		}
		//car_acs
	case "card_acs":
		// Check if encrypted_card_number exists
		if _, exists := recMessage["column1"]; !exists {
			fmt.Println("ERROR:", bookMark, "encrypted_card_number should not be empty.")
			return SendResponse(c, false, "encrypted_card_number should not be empty.")
		}

		// Filter Card ACS data
		filteredData, err := FilterCardAcsData(recMessage, bookMark)
		if err != nil {
			return SendResponse(c, false, "Error filtering Card ACS data.")
		}

		// Prepare Query Object
		prepareQueryObj := map[string]interface{}{
			"data":      filteredData,
			"operation": operation,
			"tableName": "card_acs",
		}

		postgreQuery, err := PrepareQuery(prepareQueryObj, bookMark)
		if err != nil {
			return SendResponse(c, false, "Error preparing Card ACS query.")
		}

		_, err = RunPostgresQuery(postgreQuery, bookMark)
		if err != nil {
			return SendResponse(c, false, "Error executing Card ACS query.")
		}

		fmt.Println("INFO:", bookMark, "Card ACS operation executed successfully.")
		return SendResponse(c, true, "Card ACS operation executed successfully.")

	default:
		return SendResponse(c, false, "Invalid table type")
	}
}
