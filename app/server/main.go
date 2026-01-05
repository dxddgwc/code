package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Post represents a blog post
type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageURL string `json:"imageUrl"`
}

var posts = []Post{
	{
		ID:       1,
		Title:    "探索 Go 語言的並發",
		Content:  "Go 語言的 Goroutine 是其最著名的特性之一，它使得並發編程變得無比簡單。與傳統的線程相比，Goroutine 非常輕量，你可以在一個程序中輕鬆地創建成千上萬個 Goroutine。",
		ImageURL: "https://images.unsplash.com/photo-1579546929518-9e396f3cc809",
	},
	{
		ID:       2,
		Title:    "Vue 3 Composition API 入門",
		Content:  "Vue 3 引入了 Composition API，這是一種全新的、更靈活的方式來組織和重用組件邏輯。通過 `setup` 函數，你可以將相關的邏輯組合在一起，使得代碼更具可讀性和可維護性。",
		ImageURL: "https://images.unsplash.com/photo-1633356122544-f134324a6cee",
	},
}

func main() {
	http.HandleFunc("/api/posts", handlePosts)
	http.HandleFunc("/api/posts/", handlePost)

	log.Println("INFO: server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("FATAL: could not start server: %v", err)
	}
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		log.Printf("ERROR: could not encode posts: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	// Extract post ID from URL
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	// Find the post
	var foundPost *Post
	for i := range posts {
		if posts[i].ID == id {
			foundPost = &posts[i]
			break
		}
	}

	if foundPost == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(foundPost); err != nil {
		log.Printf("ERROR: could not encode post: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
