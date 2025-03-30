package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/joho/godotenv"
)

// Define the input structure matching Hasura's `recipes_insert_input`
type recipes_insert_input struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	CreatorID   string    `json:"creator_id"`
	PrepTime    int       `json:"prep_time"`
	CreatedAt   time.Time `json:"created_at"`
}

// Define the GraphQL mutation structure
type InsertRecipeMutation struct {
	InsertRecipesOne struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	} `graphql:"insert_recipes_one(object: $object)"`
}

func main() {
	// Load environment variables
	_ = godotenv.Load()

	endpoint := os.Getenv("HASURA_GRAPHQL_ENDPOINT") // example: http://localhost:8080/v1/graphql
	adminSecret := os.Getenv("HASURA_ADMIN_SECRET")

	if endpoint == "" || adminSecret == "" {
		log.Fatal("HASURA_GRAPHQL_ENDPOINT or HASURA_ADMIN_SECRET is not set")
	}

	// Set up the GraphQL client
	httpClient := &http.Client{
		Transport: &transportWithHeaders{
			headers: map[string]string{
				"x-hasura-admin-secret": adminSecret,
			},
		},
	}
	client := graphql.NewClient(endpoint, httpClient)

	// Prepare the mutation variables
	variables := map[string]interface{}{
		"object": recipes_insert_input{
			Title:       "Test Recipe",
			Description: "Test Desc",
			ImageURL:    "https://test.com/image.jpg",
			CreatorID:   "1bd8dfe0-ed17-42a1-8c18-57ff699401b9", // this UUID must exist in your users table
			PrepTime:    30,
			CreatedAt:   time.Now(),
		},
	}

	// Run the mutation
	var mutation InsertRecipeMutation
	err := client.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		log.Fatalf("Mutation failed: %v", err)
	}

	log.Printf("✅ Recipe inserted! ID: %s, Title: %s", mutation.InsertRecipesOne.ID, mutation.InsertRecipesOne.Title)
}

// Custom transport to add headers
type transportWithHeaders struct {
	headers map[string]string
}

func (t *transportWithHeaders) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, val := range t.headers {
		req.Header.Set(key, val)
	}
	return http.DefaultTransport.RoundTrip(req)
}
