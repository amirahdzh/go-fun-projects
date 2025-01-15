package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received: " + r.URL.Path) // Print the path of the request
	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)                       // 404
		fmt.Fprintf(w, "%v Server Error\n", http.StatusNotFound) // Print the error
		return
	}
	expirationDate := time.Now().AddDate(0, 0, 30).String()
	t.Execute(w, expirationDate)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))          // this is how we serve static files
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))) // this is how we serve static files
	fmt.Println("Listening on port 3002...")
	http.ListenAndServe(":3002", nil) // Listen on port 3002

}
