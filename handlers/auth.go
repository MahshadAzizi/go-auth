package handlers

import (
	"authentication/config"
	"authentication/models"
	"authentication/utils"
	"encoding/json"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	_, err = config.DB.Exec("INSERT INTO users (username, password, firstname, lastname) "+
		"VALUES ($1, $2, $3, $4)", user.Username, hashedPassword, user.FirstName, user.LastName)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var storedPassword string
	err := config.DB.QueryRow("SELECT password FROM users WHERE username=$1", user.Username).Scan(&storedPassword)
	if err != nil {
		http.Error(w, "Invalid credentials 2", http.StatusBadRequest)
		return
	}
	if err := utils.CheckPassword(storedPassword, user.Password); err != nil {
		http.Error(w, "Invalid credentials 3", http.StatusUnauthorized)
		return
	}

	token, error := utils.GenerateJWT(user.Username)
	if error != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})

}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-Username")
	if username == "" {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	// Step 2: Fetch user information from the database
	var user models.User
	err := config.DB.QueryRow("SELECT id, username, firstname, lastname FROM users WHERE username=$1", username).
		Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Step 3: Return user details as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
