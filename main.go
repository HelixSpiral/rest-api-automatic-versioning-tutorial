package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"rest-api-automatic-versioning-tutorial/internal/database"
)

// Item struct holds the information for a specific item
type Item struct {
	ID          string `json:"ID,omitempty"`
	Name        string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
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

// handleGetItems handles get requests to: /{version}/items/...
func handleGetItems(w http.ResponseWriter, r *http.Request) {
	var items map[string][]Item
	db := database.New("db.json")

	fileData, err := db.ReadFullDB()
	if err != nil {
		fmt.Fprintln(w, err)
	}

	json.Unmarshal(fileData, &items)

	pathSplit := strings.Split(r.URL.Path, "/")

	// If the third value is nothing then we didn't specify an item
	// So we return everything
	if pathSplit[3] == "" {
		jsonItems, err := json.MarshalIndent(items["Items"], "", "\t")
		if err != nil {
			fmt.Fprintln(w, err)
		}

		fmt.Fprintf(w, "%s", jsonItems)
		return
	}

	itemToGet := pathSplit[3]

	for _, item := range items["Items"] {
		if item.ID == itemToGet {
			jsonItem, err := json.MarshalIndent(item, "", "\t")
			if err != nil {
				fmt.Fprintln(w, err)
			}

			fmt.Fprintf(w, "%s", jsonItem)
			return
		}

	}

	fmt.Fprintln(w, "No item found:", itemToGet)
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
	http.HandleFunc(fmt.Sprintf("/%s/items/", APIVersion), handleGetItems)

	// Listen and serve until we exit
	log.Fatal(http.ListenAndServe(":80", nil))
}
