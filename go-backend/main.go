package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	graphql "github.com/hasura/go-graphql-client"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("this_is_a_very_strong_super_secret_key_123")

var hasuraEndpoint string
var hasuraAdminSecret string
var graphqlClient *graphql.Client

type AuthRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

type users_insert_input struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type UserInsertMutation struct {
	InsertUsersOne struct {
		ID string `json:"id"`
	} `graphql:"insert_users_one(object: $object)"`
}

type GetUserByEmailQuery struct {
	Users []struct {
		ID           string `json:"id" graphql:"id"`
		Email        string `json:"email" graphql:"email"`
		PasswordHash string `json:"password_hash" graphql:"password_hash"`
		Name         string `json:"name" graphql:"name"`
	} `graphql:"users(where: {email: {_eq: $email}})"`
}

func main() {
	_ = godotenv.Load()

	hasuraEndpoint = os.Getenv("HASURA_GRAPHQL_ENDPOINT")
	hasuraAdminSecret = os.Getenv("HASURA_ADMIN_SECRET")

	httpClient := &http.Client{
		Transport: &transportWithHeaders{
			headers: map[string]string{
				"x-hasura-admin-secret": hasuraAdminSecret,
			},
		},
	}
	graphqlClient = graphql.NewClient(hasuraEndpoint, httpClient)

	http.HandleFunc("/signup", enableCors(signupHandler))
	http.HandleFunc("/login", enableCors(loginHandler)) // ✅ Action-ready login

	log.Println("✅ Server running on http://localhost:8082")
	http.ListenAndServe(":8082", nil)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	vars := map[string]interface{}{
		"object": users_insert_input{
			Name:         req.Name,
			Email:        req.Email,
			PasswordHash: string(hashed),
		},
	}

	var mutation UserInsertMutation
	err = graphqlClient.Mutate(context.Background(), &mutation, vars)
	if err != nil {
		log.Println("Signup error:", err)
		http.Error(w, "Email already exists or failed", http.StatusBadRequest)
		return
	}

	token := generateToken(mutation.InsertUsersOne.ID, req.Email)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TokenResponse{
		Token:  token,
		UserID: mutation.InsertUsersOne.ID,
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// ✅ Compatible with Hasura Action structure
	var input struct {
		Input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"input"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Input.Email == "" || input.Input.Password == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Query the user by email
	var query GetUserByEmailQuery
	vars := map[string]interface{}{
		"email": graphql.String(input.Input.Email),
	}

	err = graphqlClient.Query(context.Background(), &query, vars)
	if err != nil || len(query.Users) == 0 {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	user := query.Users[0]

	// Check the password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Input.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate the token for the user
	token := generateToken(user.ID, user.Email)

	// Return the token and user info, including the avatar (first letter of name)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"user": map[string]interface{}{
			"id":      user.ID,
			"email":   user.Email,
			"name":    user.Name,
			"avatar":  getAvatar(user.Name), // Function to get first letter as avatar
		},
	})
}

// Helper function to get the first letter of the name as avatar
func getAvatar(name string) string {
	if len(name) > 0 {
		return string(name[0]) // Get the first letter of the name
	}
	return "U" // Return "U" if the name is empty
}

func generateToken(userID, email string) string {
	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-allowed-roles": []string{"user"},
			"x-hasura-default-role":  "user",
			"x-hasura-user-id":       userID,
		},
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString(jwtSecret)
	return signed
}

type transportWithHeaders struct {
	headers map[string]string
}

func (t *transportWithHeaders) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, val := range t.headers {
		req.Header.Set(key, val)
	}
	return http.DefaultTransport.RoundTrip(req)
}

func enableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		if r.Method == "OPTIONS" {
			return
		}
		next(w, r)
	}
}
