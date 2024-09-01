package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

// Check if a port is open and accessible
func isPortOpen(host string, port string) bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// // Setup ArangoDB resources
// func setupArangoDB() error {
// 	if !isPortOpen("arangodb", "8529") {
// 		return fmt.Errorf("ArangoDB port 8529 is not open or accessible")
// 	}

// 	conn, err := http.NewConnection(http.ConnectionConfig{
// 		Endpoints: []string{"http://arangodb:8529"},
// 	})
// 	if err != nil {
// 		return fmt.Errorf("failed to create ArangoDB connection: %w", err)
// 	}

// 	client, err := driver.NewClient(driver.ClientConfig{
// 		Connection:     conn,
// 		Authentication: driver.BasicAuthentication("root", "rootpassword"),
// 	})
// 	if err != nil {
// 		return fmt.Errorf("failed to create ArangoDB client: %w", err)
// 	}

// 	// Create database if it doesn't exist
// 	dbExists, err := client.DatabaseExists(context.Background(), "example_db")
// 	if err != nil {
// 		return fmt.Errorf("failed to check ArangoDB database existence: %w", err)
// 	}
// 	if !dbExists {
// 		_, err = client.CreateDatabase(context.Background(), "example_db", nil)
// 		if err != nil {
// 			return fmt.Errorf("failed to create ArangoDB database: %w", err)
// 		}
// 		fmt.Println("ArangoDB database 'example_db' created")
// 	}

// 	// Create collection if it doesn't exist
// 	db, err := client.Database(context.Background(), "example_db")
// 	if err != nil {
// 		return fmt.Errorf("failed to open ArangoDB database: %w", err)
// 	}
// 	collExists, err := db.CollectionExists(context.Background(), "example_collection")
// 	if err != nil {
// 		return fmt.Errorf("failed to check ArangoDB collection existence: %w", err)
// 	}
// 	if !collExists {
// 		_, err = db.CreateCollection(context.Background(), "example_collection", nil)
// 		if err != nil {
// 			return fmt.Errorf("failed to create ArangoDB collection: %w", err)
// 		}
// 		fmt.Println("ArangoDB collection 'example_collection' created")
// 	}

// 	// Insert a test document into the collection
// 	coll, err := db.Collection(context.Background(), "example_collection")
// 	if err != nil {
// 		return fmt.Errorf("failed to access ArangoDB collection: %w", err)
// 	}
// 	doc := map[string]interface{}{"_key": "test_doc", "name": "test_name"}
// 	_, err = coll.CreateDocument(context.Background(), doc)
// 	if err != nil {
// 		return fmt.Errorf("failed to insert test document into ArangoDB collection: %w", err)
// 	}
// 	fmt.Println("ArangoDB test document inserted")

// 	return nil
// }

func setupArangoDB() error {
	if !isPortOpen("arangodb", "8529") {
		return fmt.Errorf("ArangoDB port 8529 is not open or accessible")
	}

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://arangodb:8529"},
	})
	if err != nil {
		return fmt.Errorf("failed to create ArangoDB connection: %w", err)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "rootpassword"),
	})
	if err != nil {
		return fmt.Errorf("failed to create ArangoDB client: %w", err)
	}

	// Create database if it doesn't exist
	dbExists, err := client.DatabaseExists(context.Background(), "example_db")
	if err != nil {
		return fmt.Errorf("failed to check ArangoDB database existence: %w", err)
	}
	if !dbExists {
		_, err = client.CreateDatabase(context.Background(), "example_db", nil)
		if err != nil {
			return fmt.Errorf("failed to create ArangoDB database: %w", err)
		}
		fmt.Println("ArangoDB database 'example_db' created")
	}

	// Create collection if it doesn't exist
	db, err := client.Database(context.Background(), "example_db")
	if err != nil {
		return fmt.Errorf("failed to open ArangoDB database: %w", err)
	}
	collExists, err := db.CollectionExists(context.Background(), "example_collection")
	if err != nil {
		return fmt.Errorf("failed to check ArangoDB collection existence: %w", err)
	}
	if !collExists {
		_, err = db.CreateCollection(context.Background(), "example_collection", nil)
		if err != nil {
			return fmt.Errorf("failed to create ArangoDB collection: %w", err)
		}
		fmt.Println("ArangoDB collection 'example_collection' created")
	}

	// Check if the test document already exists
	coll, err := db.Collection(context.Background(), "example_collection")
	if err != nil {
		return fmt.Errorf("failed to access ArangoDB collection: %w", err)
	}

	// Attempt to read the document to see if it exists
	var existingDoc map[string]interface{}
	_, err = coll.ReadDocument(context.Background(), "test_doc", &existingDoc)
	if driver.IsNotFound(err) {
		// Document does not exist; create it
		doc := map[string]interface{}{"_key": "test_doc", "name": "test_name"}
		_, err = coll.CreateDocument(context.Background(), doc)
		if err != nil {
			return fmt.Errorf("failed to insert test document into ArangoDB collection: %w", err)
		}
		fmt.Println("ArangoDB test document inserted")
	} else if err != nil {
		// An error other than "not found" occurred
		return fmt.Errorf("failed to check for existing ArangoDB document: %w", err)
	} else {
		fmt.Println("ArangoDB test document already exists, skipping insertion")
	}

	return nil
}


// Setup PostgreSQL resources
func setupPostgreSQL() error {
	if !isPortOpen("postgres", "5432") {
		return fmt.Errorf("PostgreSQL port 5432 is not open or accessible")
	}

	connStr := "user=postgres password=yourpassword dbname=postgres sslmode=disable host=postgres port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	defer db.Close()

	// Create database if it doesn't exist
	_, err = db.Exec("CREATE DATABASE example_db")
	if err != nil && err.Error() != `pq: database "example_db" already exists` {
		return fmt.Errorf("failed to create PostgreSQL database: %w", err)
	}
	fmt.Println("PostgreSQL database 'example_db' created or already exists")

	// Connect to the new database
	connStr = "user=postgres password=yourpassword dbname=example_db sslmode=disable host=postgres port=5432"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL example_db: %w", err)
	}
	defer db.Close()

	// Create table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS example_table (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	)`)
	if err != nil {
		return fmt.Errorf("failed to create PostgreSQL table: %w", err)
	}
	fmt.Println("PostgreSQL table 'example_table' created or already exists")

	// Insert a test row into the table
	_, err = db.Exec("INSERT INTO example_table (name) VALUES ($1) ON CONFLICT DO NOTHING", "test_name")
	if err != nil {
		return fmt.Errorf("failed to insert test row into PostgreSQL table: %w", err)
	}
	fmt.Println("PostgreSQL test row inserted")

	return nil
}

// Setup Redis resources
func setupRedis() error {
	if !isPortOpen("redis", "6379") {
		return fmt.Errorf("Redis port 6379 is not open or accessible")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	// Set a key-value pair
	err := client.Set(ctx, "example_key", "example_value", 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set Redis key: %w", err)
	}
	fmt.Println("Redis key 'example_key' set")

	// Set another key for additional check
	err = client.Set(ctx, "sanity_check_key", "sanity_check_value", 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set additional Redis key: %w", err)
	}
	fmt.Println("Redis sanity check key 'sanity_check_key' set")

	return nil
}

// Sanity check for ArangoDB
func checkArangoDB() error {
	var client driver.Client
	var err error

	for i := 0; i < 5; i++ { // Retry up to 5 times
		conn, err := http.NewConnection(http.ConnectionConfig{
			Endpoints: []string{"http://arangodb:8529"},
		})
		if err != nil {
			log.Printf("Failed to create ArangoDB connection: %v. Retrying...", err)
			time.Sleep(5 * time.Second)
			continue
		}

		client, err = driver.NewClient(driver.ClientConfig{
			Connection:     conn,
			Authentication: driver.BasicAuthentication("root", "rootpassword"),
		})
		if err != nil {
			log.Printf("Failed to create ArangoDB client: %v. Retrying...", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Connection check
		_, err = client.Version(context.Background())
		if err == nil {
			break // Exit the loop if successful
		}

		log.Printf("Failed to connect to ArangoDB: %v. Retrying...", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to ArangoDB after retries: %w", err)
	}

	// Database and collection existence check
	db, err := client.Database(context.Background(), "example_db")
	if err != nil {
		return fmt.Errorf("failed to open ArangoDB database: %w", err)
	}
	coll, err := db.Collection(context.Background(), "example_collection")
	if err != nil {
		return fmt.Errorf("failed to check ArangoDB collection existence: %w", err)
	}

	// Document existence check
	var doc map[string]interface{}
	_, err = coll.ReadDocument(context.Background(), "test_doc", &doc)
	if err != nil {
		return fmt.Errorf("ArangoDB test document does not exist or could not be read: %w", err)
	}

	// Data integrity check
	if doc["name"] != "test_name" {
		return fmt.Errorf("ArangoDB test document has unexpected data: %v", doc)
	}

	fmt.Println("ArangoDB Sanity Check passed with document validation")
	return nil
}

// Sanity check for PostgreSQL
func checkPostgreSQL() error {
	connStr := "user=postgres password=yourpassword dbname=example_db sslmode=disable host=postgres port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	defer db.Close()

	// Connection check
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	// Table existence check and data validation
	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'example_table')").Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check PostgreSQL table existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("PostgreSQL table 'example_table' does not exist")
	}

	// Test row existence check
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM example_table WHERE name = $1", "test_name").Scan(&count)
	if err != nil || count == 0 {
		return fmt.Errorf("PostgreSQL test row does not exist or could not be validated")
	}

	// Performance check (dummy value)
	start := time.Now()
	_ = db.Ping() // Simulating a query
	elapsed := time.Since(start)
	if elapsed > 10*time.Second {
		return fmt.Errorf("PostgreSQL query response time is too high: %v", elapsed)
	}

	fmt.Println("PostgreSQL Sanity Check passed with row validation and performance check")
	return nil
}

// Sanity check for Redis
func checkRedis() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	// Connection and ping check
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to ping Redis: %w", err)
	}
	fmt.Println("Redis ping response:", pong)

	// Key existence and value check
	val, err := client.Get(ctx, "sanity_check_key").Result()
	if err != nil {
		return fmt.Errorf("failed to retrieve Redis sanity check key: %w", err)
	}
	if val != "sanity_check_value" {
		return fmt.Errorf("Redis sanity check key has unexpected value: %s", val)
	}

	// Performance check (dummy value)
	start := time.Now()
	_, err = client.Ping(ctx).Result()
	elapsed := time.Since(start)
	if elapsed > 10*time.Second {
		return fmt.Errorf("Redis query response time is too high: %v", elapsed)
	}

	fmt.Println("Redis Sanity Check passed with key validation and performance check")
	return nil
}

func main() {
	// Set up resources
	if err := setupArangoDB(); err != nil {
		log.Fatalf("ArangoDB setup failed: %v", err)
	}
	if err := setupPostgreSQL(); err != nil {
		log.Fatalf("PostgreSQL setup failed: %v", err)
	}
	if err := setupRedis(); err != nil {
		log.Fatalf("Redis setup failed: %v", err)
	}

	// Perform sanity checks
	if err := checkArangoDB(); err != nil {
		log.Fatalf("ArangoDB Sanity Check failed: %v", err)
	}
	if err := checkPostgreSQL(); err != nil {
		log.Fatalf("PostgreSQL Sanity Check failed: %v", err)
	}
	if err := checkRedis(); err != nil {
		log.Fatalf("Redis Sanity Check failed: %v", err)
	}

	fmt.Println("All sanity checks passed. Sleeping for 10 seconds...")
	time.Sleep(10 * time.Second) // Delay to allow logs to be viewed
}
