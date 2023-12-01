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

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		var body []map[string]any
		for _, row := range data[:20] {
			body = append(body, map[string]any{
				headers[0]: row[0],
				headers[1]: strings.Split(row[1], "\t"),
				headers[2]: strings.Split(row[2], "\t"),
				headers[3]: row[3],
				headers[4]: row[4],
			})
		}
		json.NewEncoder(w).Encode(body)
	})
	http.ListenAndServe(":8080", nil)
}
