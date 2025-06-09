package database
import (
	
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the database connection using the provided connection string.
func InitDB(connStr string) {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Check if the connection is established
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	fmt.Println("Database connection established successfully")
}

// CloseDB closes the database connection.
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error closing the database connection: %v", err)
		} else {
			fmt.Println("Database connection closed successfully")
		}
	} else {
		fmt.Println("No database connection to close")
	}
}
