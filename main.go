package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-api-automatic-versioning-tutorial/internal/database"
)

// Item struct holds the information for a specific item
type Item struct {
	ID   string `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
}

// APIVersion holds the version of the API we're building.
// We populate this variable during the build
var APIVersion string

// handleRoot handles get requests to: /
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Root called, you probably want to hit one of the API endpoints.")
}

// handleAPIRoot handles get requests to: /{version}
func handleAPIRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API root called")
	fmt.Fprintf(w, "API Version: %s", APIVersion)
}

// handleGetAllItems handles get requests to: /{version}/items
func handleGetAllItems(w http.ResponseWriter, r *http.Request) {
	db := database.New("db.json")

	fileData, err := db.ReadFullDB()
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, "%s", fileData)
}

func main() {
	// If the API Version is empty then we haven't released v1 yet
	// So we just populate it with v0
	if APIVersion == "" {
		APIVersion = "v0"
	}

	// Setup our handlers
	http.HandleFunc("/", handleRoot)
	http.HandleFunc(fmt.Sprintf("/%s", APIVersion), handleAPIRoot)
	http.HandleFunc(fmt.Sprintf("/%s/items", APIVersion), handleGetAllItems)

	// Listen and serve until we exit
	log.Fatal(http.ListenAndServe(":80", nil))
}
