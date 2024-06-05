package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)


func main() {
	// Accept command line arguments to see which setting to be used
	var option string
	flag.StringVar(&option, )

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

	switch option{
	case "yaml":
			// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := ioutil.ReadFile("data.yaml")
		yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
		if err != nil {
			panic(err)
		}
	
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", mapHandler)

	case "json":
	case "db":
	default:
		fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}
	


	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

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