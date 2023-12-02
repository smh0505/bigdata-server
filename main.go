package main

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
	"slices"
	"strings"
)

type Row struct {
	//id,ingredients_str,allergens_str,title,allergens
	Id          string   `json:"id"`
	Ingredients []string `json:"ingredients_str"`
	Allergens   []string `json:"allergens_str"`
	Title       string   `json:"title"`
	AllergenStr string   `json:"allergens"`
}

func main() {
	db, _ := os.Open("Ready.csv")
	defer db.Close()

	reader := csv.NewReader(db)
	data, _ := reader.ReadAll()
	data = data[1:]

	var database []Row
	for _, row := range data {
		database = append(database, Row{
			Id:          row[0],
			Ingredients: strings.Split(row[1], "\t"),
			Allergens:   strings.Split(row[2], "\t"),
			Title:       row[3],
			AllergenStr: row[4],
		})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body struct {
			Index     int      `json:"index"`
			Allergens []string `json:"allergens"`
		}
		decoder.Decode(&body)

		var filtered []Row
		for _, row := range database {
			check := false
			for _, allergen := range body.Allergens {
				check = check || slices.Contains(row.Allergens, allergen)
			}
			if !check {
				filtered = append(filtered, row)
			}
		}

		begin := max(0, min(body.Index, len(filtered)-1))
		end := min(begin+20, len(filtered))

		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(filtered[begin:end])
	})
	http.ListenAndServe(":8080", nil)
}
