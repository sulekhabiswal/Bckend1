// package IntermediateServices

// import (
// 	"fmt"
// 	"strings"
// )

// // PrepareBulkQuery generates a bulk INSERT or UPDATE query
// func PrepareBulkQuery(prepareQueryObj map[string]interface{}, bookMark string) (*QueryObject, error) {
// 	// Extract required values
// 	tableName, tableExists := prepareQueryObj["tableName"].(string)
// 	keyValueList, keyValueExists := prepareQueryObj["keyValueList"].([]map[string]interface{})
// 	valueList, _ := prepareQueryObj["valueList"].([][]interface{})
// 	operation, operationExists := prepareQueryObj["operation"].(string)

// 	if !tableExists || !keyValueExists || !operationExists {
// 		return nil, fmt.Errorf("invalid query preparation object")
// 	}

// 	if operation == "CREATE" {
// 		// Define column names for INSERT
// 		columnList := []string{"created_date", "is_applicable", "mcc_id", "updated_date", "card_card_refnumber"}
// 		values := []interface{}{}

// 		// Flatten valueList into a single list of values
// 		for _, row := range valueList {
// 			values = append(values, row...)
// 		}

// 		// Create placeholders for each row
// 		var placeholders []string
// 		for rowIndex := range keyValueList {
// 			var rowPlaceholders []string
// 			for colIndex := range columnList {
// 				rowPlaceholders = append(rowPlaceholders, fmt.Sprintf("$%d", rowIndex*len(columnList)+colIndex+1))
// 			}
// 			placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", ")))
// 		}

// 		// Construct INSERT query
// 		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;",
// 			tableName, strings.Join(columnList, ", "), strings.Join(placeholders, ", "))

// 		// Construct QueryObject
// 		return &QueryObject{
// 			Query:  query,
// 			Values: values,
// 		}, nil

// 	} else if operation == "UPDATE" {
// 		// Define primary keys and update columns
// 		primaryKey := "mcc_id"
// 		foreignKey := "card_card_refnumber"
// 		updateColumn := "is_applicable"
// 		updateDateColumn := "updated_date" // New column for timestamp update
// 		values := []interface{}{}

// 		// Construct "VALUES" clause for bulk update
// 		var valuePlaceholders []string
// 		for i, row := range keyValueList {
// 			offset := i * 4 // 4 columns per row (mcc_id, card_card_refnumber, is_applicable, updated_date)
// 			valuePlaceholders = append(valuePlaceholders,
// 				fmt.Sprintf("($%d::BIGINT, $%d::BIGINT, $%d::BOOLEAN, $%d::TIMESTAMP)", offset+1, offset+2, offset+3, offset+4))

// 			// Append actual values
// 			values = append(values, row[primaryKey], row[foreignKey], row[updateColumn], row[updateDateColumn])
// 		}

// 		// Construct UPDATE query
// 		query := fmt.Sprintf(
// 			`UPDATE %s AS t 
// 			 SET %s = v.%s, %s = v.%s
// 			 FROM (VALUES %s) AS v (%s, %s, %s, %s)
// 			 WHERE t.%s = v.%s AND t.%s = v.%s;`,
// 			tableName, updateColumn, updateColumn, updateDateColumn, updateDateColumn,
// 			strings.Join(valuePlaceholders, ", "), primaryKey, foreignKey, updateColumn, updateDateColumn,
// 			primaryKey, primaryKey, foreignKey, foreignKey)


// 			fmt.Println(query)
// 		// Construct QueryObject
// 		return &QueryObject{
// 			Query:  query,
// 			Values: values,
// 		}, nil
// 	}

// 	return nil, fmt.Errorf("invalid operation type")
// }



// package IntermediateServices

// import (
// 	"fmt"
// 	"strings"
// )

// // PrepareBulkQuery generates a bulk INSERT or UPDATE query
// func PrepareBulkQuery(prepareQueryObj map[string]interface{}, bookMark string) (*QueryObject, error) {
// 	// Extract required values
// 	tableName, tableExists := prepareQueryObj["tableName"].(string)
// 	keyValueList, keyValueExists := prepareQueryObj["keyValueList"].([]map[string]interface{})
// 	valueList, _ := prepareQueryObj["valueList"].([][]interface{})
// 	operation, operationExists := prepareQueryObj["operation"].(string)

// 	if !tableExists || !keyValueExists || !operationExists {
// 		return nil, fmt.Errorf("invalid query preparation object")
// 	}

// 	if operation == "CREATE" {
// 		// Define column names for INSERT
// 		columnList := []string{"created_date", "is_applicable", "mcc_id", "updated_date", "card_card_refnumber"}
// 		values := []interface{}{}

// 		// Flatten valueList into a single list of values
// 		for _, row := range valueList {
// 			values = append(values, row...)
// 		}

// 		// Create placeholders for each row
// 		var placeholders []string
// 		for rowIndex := range keyValueList {
// 			var rowPlaceholders []string
// 			for colIndex := range columnList {
// 				rowPlaceholders = append(rowPlaceholders, fmt.Sprintf("$%d", rowIndex*len(columnList)+colIndex+1))
// 			}
// 			placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", ")))
// 		}

// 		// Construct INSERT query
// 		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;",
// 			tableName, strings.Join(columnList, ", "), strings.Join(placeholders, ", "))

// 		// Construct QueryObject
// 		return &QueryObject{
// 			Query:  query,
// 			Values: values,
// 		}, nil

// 	} else if operation == "UPDATE" {
// 		// Define primary keys and update columns
// 		primaryKey := "mcc_id"
// 		foreignKey := "card_card_refnumber"
// 		updateColumn := "is_applicable"
// 		values := []interface{}{}

// 		// Construct "VALUES" clause for bulk update (3 columns per row)
// 		var valuePlaceholders []string
// 		for i, row := range keyValueList {
// 			offset := i * 3 // Only 3 columns per row now
// 			valuePlaceholders = append(valuePlaceholders,
// 				fmt.Sprintf("($%d::BIGINT, $%d::BIGINT, $%d::BOOLEAN)", offset+1, offset+2, offset+3))

// 			// Append actual values
// 			values = append(values, row[primaryKey], row[foreignKey], row[updateColumn])
// 		}

// 		// Construct UPDATE query
// 		query := fmt.Sprintf(
// 			`UPDATE %s AS t 
// 			 SET %s = v.%s
// 			 FROM (VALUES %s) AS v (%s::BIGINT, %s::BIGINT, %s::BOOLEAN)
// 			 WHERE t.%s = v.%s AND t.%s = v.%s;`,
// 			tableName, updateColumn, updateColumn,
// 			strings.Join(valuePlaceholders, ", "), primaryKey, foreignKey, updateColumn,
// 			primaryKey, primaryKey, foreignKey, foreignKey)

// 		// Construct QueryObject
// 		return &QueryObject{
// 			Query:  query,
// 			Values: values,
// 		}, nil
// 	}

// 	return nil, fmt.Errorf("invalid operation type")
// }



package IntermediateServices

import (
    "errors"
    "fmt"
    "strings"
)

// PrepareBulkQuery generates a bulk INSERT or UPDATE query for MCC
func PrepareBulkQuery(prepareQueryObj map[string]interface{}, bookMark string) (*QueryObject, error) {
    // Extract required values
    tableName, tableExists := prepareQueryObj["tableName"].(string)
    keyValueList, keyValueExists := prepareQueryObj["keyValueList"].([]map[string]interface{})
    valueList, _ := prepareQueryObj["valueList"].([][]interface{})
    operation, operationExists := prepareQueryObj["operation"].(string)

    if !tableExists || !keyValueExists || !operationExists {
        return nil, errors.New("invalid query preparation object")
    }

    // CREATE Operation**
    if operation == "CREATE" {
        // Define column names for INSERT
        columnList := []string{"created_date", "is_applicable", "mcc_id", "updated_date", "card_card_refnumber"}
        values := []interface{}{}

        // Flatten valueList into a single list of values
        for _, row := range valueList {
            values = append(values, row...)
        }

        // Create placeholders for each row
        var placeholders []string
        for rowIndex := range keyValueList {
            var rowPlaceholders []string
            for colIndex := range columnList {
                rowPlaceholders = append(rowPlaceholders, fmt.Sprintf("$%d", rowIndex*len(columnList)+colIndex+1))
            }
            placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", ")))
        }

        // Construct INSERT query
        query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;",
            tableName, strings.Join(columnList, ", "), strings.Join(placeholders, ", "))

        fmt.Println("PostgresOperation : Query prepared successfully for card MCC.")

        // Construct QueryObject
        return &QueryObject{
            Query:  query,
            Values: values,
        }, nil

        //  **UPDATE Operation (Fixed Missing Column Issue)**
    } else if operation == "UPDATE" {
        // Define primary keys and update columns
        primaryKey := "mcc_id"
        foreignKey := "card_card_refnumber"
        updateColumn := "is_applicable"
        values := []interface{}{}

        // **Construct "VALUES" clause for bulk update (i*3 logic)**
        var valuePlaceholders []string
        for i, row := range keyValueList {
            // Ensure required keys exist
            if row[primaryKey] == nil || row[foreignKey] == nil || row[updateColumn] == nil {
                fmt.Println("Skipping row due to missing keys:", row)
                continue
            }

            // Append placeholders (Using `i*3` approach)
            offset := i * 3 // Only 3 columns per row now
            valuePlaceholders = append(valuePlaceholders,
                fmt.Sprintf("($%d::BIGINT, $%d::BIGINT, $%d::BOOLEAN)", offset+1, offset+2, offset+3))

            // Append actual values
            values = append(values, row[primaryKey], row[foreignKey], row[updateColumn])
        }

        //  Prevent Empty UPDATE
        if len(valuePlaceholders) == 0 {
            fmt.Println(" No data to update for MCC")
            return nil, fmt.Errorf("no valid MCC records to update")
        }

        // Fix: Ensure Column Names Are Correct in `SELECT` Clause**
        query := fmt.Sprintf(
            `UPDATE %s AS t 
         SET is_applicable = v.value3
         FROM (VALUES %s) AS v(value1, value2, value3)
         WHERE t.%s = v.value1 AND t.%s = v.value2;`,
            tableName,
            strings.Join(valuePlaceholders, ", "),
            primaryKey, foreignKey)

        fmt.Println("PostgresOperation : Query prepared successfully for card MCC.")
        // Construct QueryObject
        return &QueryObject{
            Query:  query,
            Values: values,
        }, nil
    }

    //  Invalid Operation
    return nil, fmt.Errorf("invalid operation type")
}



