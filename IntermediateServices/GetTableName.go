package IntermediateServices

import (
	"fmt"
	"time"
)


var bigqueryTableName = "bigquery_table_prefix_" 

// GetTableName generates a table name based on the transaction creation date
func GetTableName(transactionCreatedDate string, bookMark string) (string, error) {
	
	createdDate, err := time.Parse(time.RFC3339, transactionCreatedDate)
	if err != nil {
		fmt.Println("ERROR:", bookMark, "Invalid transaction created date:", err)
		return "", fmt.Errorf("error parsing transaction created date: %v", err)
	}

	// Convert to IST (Asia/Kolkata timezone)
	loc, _ := time.LoadLocation("Asia/Kolkata")
	currDate := createdDate.In(loc)

	// Extract date components (YYYYMMDD format)
	cDate := fmt.Sprintf("%02d", currDate.Day())
	cMonth := fmt.Sprintf("%02d", currDate.Month())
	cYear := currDate.Year()

	// Generate table name
	tableName := fmt.Sprintf("%s%d%s%s", bigqueryTableName, cYear, cMonth, cDate)

	return tableName, nil
}
