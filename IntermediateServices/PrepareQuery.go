// package IntermediateServices

// import (
// 	"fmt"
// 	"strings"
// )

// // QueryObject represents a structured PostgreSQL query
// type QueryObject struct {
// 	Query  string
// 	Values []interface{}
// }

// // PrepareQuery generates a SQL query for inserting or updating a single record
// func PrepareQuery(prepareQueryObj map[string]interface{}, bookMark string) (*QueryObject, error) {
// 	// Extract required values
// 	data, dataExists := prepareQueryObj["data"].(map[string]interface{})
// 	operation, operationExists := prepareQueryObj["operation"].(string)
// 	tableName, tableExists := prepareQueryObj["tableName"].(string)

// 	if !dataExists || !operationExists || !tableExists {
// 		return nil, fmt.Errorf("invalid query preparation object")
// 	}

// 	columns := []string{}
// 	values := []interface{}{}

// 	// Extract column names and values
// 	for key, value := range data {
// 		if value != nil { // Ignore null values
// 			columns = append(columns, key)
// 			values = append(values, value)
// 		}
// 	}

// 	var query string

// 	if operation == "CREATE" {
// 		// Construct an INSERT query
// 		placeholders := []string{}
// 		for i := range values {
// 			placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
// 		}
// 		query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
// 			tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

// 	} else if operation == "UPDATE" {
// 		// Construct an UPDATE query
// 		if len(columns) == 0 {
// 			return nil, fmt.Errorf("no columns to update")
// 		}

// 		primaryKeyName := columns[0]  // First column is the primary key
// 		primaryKeyValue := values[0]  

// 		columnsToUpdate := columns[1:] // Exclude the primary key
// 		valuesToUpdate := values[1:]   // Exclude the primary key value

// 		setClauses := []string{}
// 		for i, col := range columnsToUpdate {
// 			setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, i+1))
// 		}

// 		whereClause := fmt.Sprintf("WHERE %s = $%d", primaryKeyName, len(columnsToUpdate)+1)
// 		valuesToUpdate = append(valuesToUpdate, primaryKeyValue)

// 		query = fmt.Sprintf("UPDATE %s SET %s %s;",
// 			tableName, strings.Join(setClauses, ", "), whereClause)
// 		values = valuesToUpdate
// 	}

// 	// Construct the QueryObject
// 	queryObj := &QueryObject{
// 		Query:  query,
// 		Values: values,
// 	}

// 	fmt.Println("Generated Query:", queryObj.Query)
// 	return queryObj, nil
// }


package IntermediateServices

import (
	"fmt"
	"strings"
)

// QueryObject represents a structured PostgreSQL query
type QueryObject struct {
	Query  string
	Values []interface{}
}

// PrepareQuery generates a SQL query for inserting or updating a single record while maintaining column order
func PrepareQuery(prepareQueryObj map[string]interface{}, bookMark string) (*QueryObject, error) {
	// Extract required values
	data, dataExists := prepareQueryObj["data"].(map[string]interface{})
	operation, operationExists := prepareQueryObj["operation"].(string)
	tableName, tableExists := prepareQueryObj["tableName"].(string)

	if !dataExists || !operationExists || !tableExists {
		return nil, fmt.Errorf("invalid query preparation object")
	}

	// Step 1: Get Column Order from FilterData.go
	orderedColumns := getColumnOrderFromFilterData(tableName)
	orderedValues := make([]interface{}, 0, len(orderedColumns))

	for _, col := range orderedColumns {
		if val, exists := data[col]; exists {
			orderedValues = append(orderedValues, val)
		}
	}

	if len(orderedColumns) == 0 {
		return nil, fmt.Errorf("no columns available for operation")
	}

	var query string

	// Step 2: Construct INSERT Query
	if operation == "CREATE" {
		placeholders := make([]string, len(orderedValues))
		for i := range orderedValues {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
		}
		query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
			tableName, strings.Join(orderedColumns, ", "), strings.Join(placeholders, ", "))

		// Step 3: Construct UPDATE Query
	} else if operation == "UPDATE" {
		primaryKeyName := orderedColumns[0] // First column is the primary key
		primaryKeyValue := orderedValues[0]

		columnsToUpdate := orderedColumns[1:]
		valuesToUpdate := orderedValues[1:]

		setClauses := make([]string, len(columnsToUpdate))
		for i, col := range columnsToUpdate {
			setClauses[i] = fmt.Sprintf("%s = $%d", col, i+1)
		}

		whereClause := fmt.Sprintf("WHERE %s = $%d", primaryKeyName, len(columnsToUpdate)+1)
		valuesToUpdate = append(valuesToUpdate, primaryKeyValue)

		query = fmt.Sprintf("UPDATE %s SET %s %s;",
			tableName, strings.Join(setClauses, ", "), whereClause)

		orderedValues = valuesToUpdate
	}

	// Step 4: Return Query Object
	queryObj := &QueryObject{
		Query:  query,
		Values: orderedValues,
	}

	fmt.Println("Generated Query:", queryObj.Query)
	return queryObj, nil
}

// getColumnOrderFromFilterData retrieves the column order based on the table type
func getColumnOrderFromFilterData(tableName string) []string {
	switch tableName {
	case "customer":
		return []string{
			"user_id", "user_name", "first_name", "last_name", "user_role",
			"address", "city", "state", "pincode", "self_onboard",
			"email", "mobile_number", "approver_role", "approver_id", "approved_by",
			"created_by", "is_active", "admin_name", "creator_id", "creator_role",
			"created_date", "updated_date", "address1", "admin_role", "channel",
			"country", "is_child", "is_block", "block_remarks", "block_timestamp",
			"admin_id", "kyc_status", "kyc_type", "is_fraud", "fraud_remarks",
			"fraud_timestamp", "remarks", "param_a", "param_b", "param_c",
			"kyc_number", "gender", "date_of_birth", "latlong",
		}
	case "card":
		return []string{
			// "card_refnumber", "customer_id", "atm_limit", "bin_number", "parent_customer_id",
			// "ecommerce_limit", "is_ecommerce_allowed", "pos_limit", "is_pos_allowed", "is_atmallowed",
			// "mobile_number", "card_brand", "card_type", "card_limit", "encrypted_card_number",
			// "last_fourdigit", "jit_allowed", "wallet_status", "wallet_id", "service_code",
			// "expiry_date", "is_active", "pinmailer_flag", "pinoffset", "pin_try_count",
			// "created_date", "updated_date", "product_id", "product_name", "product_type",
			// "product_category", "is_child", "is_block", "block_remarks", "block_timestamp",
			// "is_fraud", "fraud_remarks", "fraud_timestamp", "name_on_card", "param_a",
			// "param_b", "param_c", "shipping_addressline1", "shipping_addressline2", "shipping_city",
			// "shipping_country", "shipping_pincode", "shipping_state", "is_close", "is_embossa",
			// "is_green_pin", "is_hotlist", "is_lost", "is_parallel_add_on", "is_physical_pin",
			// "is_si", "is_stolen", "is_suspended", "is_virtual", "is_damage",
			// "damage_timestamp", "lost_timestamp", "stolen_timestamp", "de_active", "deactive_timestamp",
			// "hotlist_timestamp", "is_contactless_allowed", "per_transaction_max_limit_in_contactless",
			"card_refnumber", "customer_id","created_date", "updated_date",
		}
	case "card_acs":
		return []string{
			"encrypted_card_number", "card_refnumber", "mobile_number", "expiry_date", "is_active",
			"created_date", "updated_date", "is_si_enable", "is_block", "is_fraud",
			"is_hotlist", "is_lost", "is_stolen",
		}
	default:
		return []string{} // Return empty list if table is unknown
	}
}
