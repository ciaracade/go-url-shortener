package main

import (
	"fmt"
	"net/http"
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
			fmt.Printf("New dest:", dest)
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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Write code that parses the provided YAML data
	// extracts paths and URLs, and creates a map.
	// Then, use the logic from MapHandler to handle requests.
	return nil, nil
}