# URL Shortener in Golang 

A URL shortener written in [GoLang](https://go.dev/) that will look at the path of any incoming web request and determine if it should redirect the user to a new page.

## Features 

### MapHandler
- Returns an http.HandlerFunc that will attempt to map any paths to their corresponding URL 
- If the path is not provided in the map, then the fallback http.Handler will be called instead.

### YAMLHandler
- Parses the provided YAML and then returns an http.HandlerFunc that will attempt to map any paths to their corresponding URL. 
- If the path is not provided in the YAML, then thefallback http.Handler  will be called instead.

### JSONHandler 
- Paths provided by JSON data
- Returns an http.HandlerFunc that will attempt to map any paths to their corresponding URL 
- If the path is not provided in the map, then the fallback http.Handler will be called instead.


### BOLTDBHandler
- DB by [Bolt](https://pkg.go.dev/github.com/boltdb/bolt) 
- Returns an http.HandlerFunc that will attempt to find in a database any paths to their corresponding URL 
- If the path is not provided in the database, then the fallback http.Handler will be called instead.

### SQLDBHandler
- DB by [SQLite](https://pkg.go.dev/github.com/mattn/go-sqlite3) 
- Returns an http.HandlerFunc that will attempt to find in a database any paths to their corresponding URL 
- If the path is not provided in the database, then the fallback http.Handler will be called instead.



## Acknowledgement 
**From [Gophercises](https://gophercises.com/)**