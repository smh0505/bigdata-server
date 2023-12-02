package main

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

func main() {
	db, _ := os.Open("Ready.csv")
	defer db.Close()

	reader := csv.NewReader(db)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	headers, data := data[0], data[1:]

	var database []map[string]any
	for _, row := range data {
		database = append(database, map[string]any{
			headers[0]: row[0],
			headers[1]: strings.Split(row[1], "\t"),
			headers[2]: strings.Split(row[2], "\t"),
			headers[3]: row[3],
			headers[4]: row[4],
		})
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body struct {
			Index 	  int 	   `json:"index"`
			Allergens []string `json:"allergens"`
		}
		decoder.Decode(&body)
		
		begin := max(0, min(body.Index, len(database) - 1))
		end := min(begin + 20, len(database))
		
		json.NewEncoder(w).Encode(database[begin:end])
	})
	http.ListenAndServe(":8080", nil)
}
