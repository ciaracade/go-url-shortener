package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"flag"
	"github.com/boltdb/bolt"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)


func main() {
	// Accept command line arguments to see which setting to be used
	var yamlFlag, jsonFlag, boltdbFlag, sqldbFlag string

	// Command line arguments -> variable, flag name, default, desc
	flag.StringVar(&yamlFlag, "yaml", "", "file name for yaml URL paths")
	flag.StringVar(&jsonFlag, "json", "", "file name for json URL paths")
	flag.StringVar(&boltdbFlag, "boltdb", "", "file name for bolt database URl paths")
	flag.StringVar(&sqldbFlag, "sqldb", "", "file name for SQL database URl paths")

	// Parse the flags
	flag.Parse()

	mux := defaultMux()

	// Check if mux is nil before passing it as fallback
  if mux == nil {
        panic("Mux is nil")
    }

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	// Check the mode and execute corresponding block of code
	if yamlFlag != "" {
		fmt.Println("Loading via yaml file:", yamlFlag)
		
		// Get YAML Data from yaml file
		yamlData, err := ioutil.ReadFile(yamlFlag)
		if err != nil {
			panic(err)
		}

		// Build the YAMLHandler using the mapHandler as the fallback
		yamlHandler, err := YAMLHandler(yamlData, mapHandler)
		if err != nil {
			panic(err)
		}

		// Start server with handler
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", yamlHandler)

		return
	} else if jsonFlag != "" {
		fmt.Println("Loading via json file:", jsonFlag)
		
		// Get JSON data from JSON file
		jsonData, err := ioutil.ReadFile(jsonFlag)
		if err != nil {
			panic(err)
		}

		// Build the JSONHandler using the mapHandler as a fallback
		jsonHandler, err := JSONHandler(jsonData, mapHandler)
		if err != nil {
			panic(err)
		}

		// Start server with handler
		fmt.Println("starting the server :8080")
		http.ListenAndServe(":8080", jsonHandler)

		return
	} else if boltdbFlag != "" {
		fmt.Println("Loading via Bolt db file:", boltdbFlag)
		
		// Start a db and open file
		db, err := bolt.Open(boltdbFlag, 0600, nil)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		// Load data into database
		err = db.Update(func(tx *bolt.Tx) error {
			// Create bucket if it doesn't exist
			b := tx.Bucket([]byte("pathstourls"))
			if b == nil {
				var err error
				b, err = tx.CreateBucketIfNotExists([]byte("pathstourls"))
				if err != nil {
					return err
				}
			}

			err = b.Put([]byte("/wiki"),[]byte("https://www.wikipedia.org/"))
			return err
		})
		if err != nil {
			panic(err)
		}

		// Build boltdbHanlder using Bolt db and mapHandler as a fallback
		boltdbHandler := BoltDBHandler(db, mapHandler)
		if err != nil {
			panic(err)
		}

		// Start server with handler
		fmt.Println("starting the server :8080")
		http.ListenAndServe(":8080", boltdbHandler)

		return
	} else if sqldbFlag != "" {
		fmt.Println("Loading via SQLite db file.")

		// Start a db and open a file
		db, err := sql.Open("sqlite3", sqldbFlag)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		const create = `
		CREATE TABLE IF NOT EXISTS pathstourls (
			path TEXT NOT NULL PRIMARY KEY,
			url TEXT NOT NULL
		);`

		if _, err := db.Exec(create); err != nil {
			panic(err)
		}

		example := `INSERT OR IGNORE INTO pathstourls VALUES(?,?);`

		_, err = db.Exec(example, "/sql", "https://sqlite.org/")
		if err != nil {
			panic(err)
		}

		// Build sqldbHanlder using SQLite db and mapHandler as a fallback
		sqldbHandler := SQLDBHanlder(db, mapHandler)

		// Start server with handler
		fmt.Println("starting the server :8080")
		http.ListenAndServe(":8080", sqldbHandler)

		return
	}

	// Default case -> use MapHandler
	fmt.Println("Loading via default case:", mapHandler)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}