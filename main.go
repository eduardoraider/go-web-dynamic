package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Data struct to hold template data
type Data struct {
	Title string
}

func main() {
	// Compile the template
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Define a handler function to serve the template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a Data struct with the title
		data := Data{Title: "My Dynamic Page"}

		// Execute the template with the data
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
