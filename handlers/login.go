package handlers

import (
	"database/sql"
	"encoding/json"
	"go-auth-api/models"
	"go-auth-api/utils"
	"net/http"
)

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.ErrorResponse(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			utils.ErrorResponse(w, "Bad request", http.StatusBadRequest)
			return
		}

		storedPassword, err := user.GetPassword(db)
		if err == sql.ErrNoRows {
			utils.ErrorResponse(w, "Invalid user", http.StatusUnauthorized)
			return
		} else if err != nil {
			utils.ErrorResponse(w, "Server error", http.StatusInternalServerError)
			return
		}

		if user.Password == storedPassword {
			utils.SuccessResponse(w, "Valid user", http.StatusOK)
		} else {
			utils.ErrorResponse(w, "Invalid user", http.StatusUnauthorized)
		}
	}
}
