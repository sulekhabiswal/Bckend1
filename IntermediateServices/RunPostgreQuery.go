package IntermediateServices

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Database configuration
const (
	host     = "localhost"
	port     = 5432               // Default PostgreSQL port
	user     = "postgres"     
	password = "Sulekha@30" 
	dbname   = "mydb"     
)

// RunPostgresQuery executes a query (INSERT or UPDATE) in PostgreSQL
func RunPostgresQuery(queryObj *QueryObject, bookMark string) (map[string]interface{}, error) {
	// Construct PostgreSQL connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("ERROR: Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}
	defer db.Close()

	// Execute the query
	result, err := db.Exec(queryObj.Query, queryObj.Values...)
	if err != nil {
		fmt.Println("ERROR:", bookMark, "Error while executing query:", err)
		return nil, err
	}

	// Retrieve affected rows count
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("ERROR:", bookMark, "Failed to get affected rows:", err)
		return nil, err
	}

	// Return success response
	fmt.Println("INFO:", bookMark, "Query executed successfully. Rows affected:", rowsAffected)
	return map[string]interface{}{
		"status":       0,
		"statusCode":   200,
		"statusDesc":   "Query executed successfully.",
		"message":      "success",
		"rowsAffected": rowsAffected,
	}, nil
}
