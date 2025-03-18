package IntermediateServices

import (
	"fmt"
)

// FilterMccData extracts and formats MCC data
func FilterMccData(cardMccList []map[string]interface{}, bookMark string) (map[string]interface{}, error) {
	// Define values that should be treated as null
	nullArr := map[interface{}]bool{nil: true, "null": true, "" : true}

	getColumnValue := func(column interface{}) interface{} {
		if _, exists := nullArr[column]; exists {
			return nil
		}
		return column
	}

	var valueList [][]interface{}
	var keyValueList []map[string]interface{}

	// Iterate over the cardMccList to extract and format data
	for _, mccData := range cardMccList {
		filteredData := map[string]interface{}{
			"created_date":        getColumnValue(mccData["column29"]),
			"is_applicable":       getColumnValue(mccData["column4"]),
			"mcc_id":              getColumnValue(mccData["column1"]),
			"updated_date":        getColumnValue(mccData["column30"]),
			"card_card_refnumber": getColumnValue(mccData["column2"]),
		}
 
		fmt.Println("INFO:", bookMark, "Filtered MCC Data:", filteredData)

		// Append only values for bulk insertion
		valueList = append(valueList, []interface{}{
			filteredData["created_date"],
			filteredData["is_applicable"],
			filteredData["mcc_id"],
			filteredData["updated_date"],
			filteredData["card_card_refnumber"],
		})

		// Append key-value pairs for updates
		keyValueList = append(keyValueList, filteredData)
	}

	// Return processed MCC data
	filteredListObj := map[string]interface{}{
		"valueList":   valueList,
		"keyValueList": keyValueList,
	}

	fmt.Println("INFO:", bookMark, "Filtered MCC Data:", filteredListObj)

	return filteredListObj, nil
}
