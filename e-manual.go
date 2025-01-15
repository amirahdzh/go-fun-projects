package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received: " + r.URL.Path) // this is how we log information to the console
	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)                       // this is how we send a specific HTTP status code
		fmt.Fprintf(w, "%v Server Error\n", http.StatusNotFound) // this is how we send a response to the client
		fmt.Fprintf(w, "Error: %v\n", err)                       // this is how we send a response to the client
		return                                                   // this is how we stop the handler function
	}
	pages, _ := scandir("./manuals") // this is how we call a function from another file
	fmt.Println(pages)
	t.Execute(w, pages) // this is how we render a template
}

func handlerNews(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received: " + r.URL.Path) // this is how we log information to the console
	t, err := template.ParseFiles("templates/news.html")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)                       // this is how we send a specific HTTP status code
		fmt.Fprintf(w, "%v Server Error\n", http.StatusNotFound) // this is how we send a response to the client
		fmt.Fprintf(w, "Error: %v\n", err)                       // this is how we send a response to the client
		return                                                   // this is how we stop the handler function
	}
	date := time.Now().String() // this is how we get the current date and time
	t.Execute(w, date)          // this is how we render a template
}

func main() {
	http.HandleFunc("/", handler)                                                                    // this is how we register a request handler
	http.HandleFunc("/news", handlerNews)                                                            // this is how we register a request handler
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))              // this is how we serve static files
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))     // this is how we serve static files
	http.Handle("/manuals/", http.StripPrefix("/manuals/", http.FileServer(http.Dir("./manuals/")))) // this is how we serve static files
	fmt.Println("Server is listening on port 3000...")
	http.ListenAndServe(":3000", nil) // this is how we start the server
}

// Run the server by running the following command in the terminal:
// go run .
