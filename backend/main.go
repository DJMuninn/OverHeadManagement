package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DJMuninn/OverHeadManagement/OverHead"
	_ "github.com/go-sql-driver/mysql"
)

// ✅ Corrected CORS middleware
func enableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {
	dsn := "DavisTour:1234@tcp(127.0.0.1:3306)/OverHeadManagement"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Health check
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w, r)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		fmt.Fprintln(w, "Backend running")
	})

	http.HandleFunc("/api/action", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w, r)

		fmt.Println("Received request to /api/action")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.Method != "POST" {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var req OverHead.APIRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Decoded request:", req)

		switch req.Action {
		case 1:
			fmt.Println("Action 1: InsertPallet triggered")
			err = OverHead.InsertPallet(db, req.Data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintln(w, "Pallet inserted successfully!")
		case 2:
			fmt.Println("Action 2: DeletePallet triggered")
			err = OverHead.DeletePallet(db, req.Data.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintln(w, "Pallet deleted successfully!")
		case 3: // action 3: update pallet
			fmt.Println("Action 3: UpdatePallet triggered")
			err = OverHead.UpdatePallet(db, req.Data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintln(w, "Pallet updated successfully!")
		default:
			fmt.Println("Unknown action:", req.Action)
			http.Error(w, "Unknown action", http.StatusBadRequest)
		}
	})

	log.Println("Backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
