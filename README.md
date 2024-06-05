# URL Shortener in Golang 

A URL shortener writtenin GoLang that will look at the path of any incoming web request and determine if it should redirect the user to a new page.


From [Gophercises](https://gophercises.com/)


## Features 
### MapHandler
- Returns an http.HandlerFunc that will attempt to map any paths to their corresponding URL 
- If the path is not provided in the map, then the fallback http.Handler will be called instead.

### YAMLHandler
- Parses the provided YAML and then return an http.HandlerFunc that will attempt to map any paths to their corresponding URL. 
- If the path is not provided in the YAML, then thefallback http.Handler  will be called instead.

### JSONHandler 
- Paths provided by JSON data
- Returns an http.HandlerFunc that will attempt to map any paths to their corresponding URL 
- If the path is not provided in the map, then the fallback http.Handler will be called instead.


### DBHandler
- Returns an http.HandlerFunc that will attempt to find in a database any paths to their corresponding URL 
- If the path is not provided in the database, then the fallback http.Handler will be called instead.