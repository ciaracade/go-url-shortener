package main

import (
	"net/http"
	yaml "gopkg.in/yaml.v2"
	"encoding/json"
	"github.com/boltdb/bolt"
	"database/sql"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// Return handler func
	return func(writer  http.ResponseWriter, response *http.Request){
		// if we can match a path
		// true/false is second argument and 
		if dest, ok := pathsToUrls[response.URL.Path]; ok {
			// go to it
			http.Redirect(writer, response, dest, http.StatusFound)
			return
		}
		// otherwise fallback to it
		fallback.ServeHTTP(writer, response)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yamlData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. Parse the yaml
	var pathURLS []pathURL
	err := yaml.Unmarshal(yamlData, &pathURLS)
	if err != nil {
		panic(err)
	}
	// 2. Convert YAML array into map
	pathsToUrls := map[string]string{}
	for _, data := range pathURLS {
		pathsToUrls[data.Path] = data.URL
	}

	// 3. return a maphanlder using map
	return MapHandler(pathsToUrls, fallback), nil
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string	`yaml:"url"`
}

func JSONHandler ( jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. Parse JSON data
	var pathURLS []jsonPathURL

	// loading byte info into structs
	err := json.Unmarshal(jsonData, &pathURLS)
	if err != nil {
		panic(err)
	}

	// 2. Convert JSON into map
	pathsToUrls := map[string]string{}
	for _, data := range pathURLS {
		pathsToUrls[data.Path] = data.URL
	}

	// 3. return a maphandler with pathToUrls
	return MapHandler(pathsToUrls, fallback), nil
}

type jsonPathURL struct {
	Path string `json:"path"`
	URL string `json:"url"`
}

func BoltDBHandler(db *bolt.DB, fallback http.Handler) http.HandlerFunc {
	// 1. Access database
	pathsToUrls := map[string]string{}

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("pathstourls"))

		// 2. Load path's url from database and turn into path
		b.ForEach(func(path, url []byte) error {
			pathsToUrls[string(path)] = string(url)
			return nil
		})
		return nil
	})
	if err != nil {
		panic(err)
	}

	// 3. return a maphandler with pathToUrls
	return MapHandler(pathsToUrls, fallback)
}

func SQLDBHanlder (db *sql.DB, fallback http.Handler) http.HandlerFunc {
	// 1. Access database
	rows, err := db.Query(`SELECT path, url FROM pathstourls;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//2. Load path's url from database or turn into path
	pathToUrls := map[string]string {}

	for rows.Next() {
		var path string
		var url string
		err := rows.Scan(&path, &url)
		if err != nil {
			panic(err)
		}
		pathToUrls[path] = url
	}
	if err = rows.Err(); err != nil {
        panic(err)
    }

	// 3. return a maphandler with pathToUrls
	return MapHandler(pathToUrls, fallback)
}

