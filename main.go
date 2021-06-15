package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
type User struct {
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email string `json:"email"`
}

type Post struct {
	Title string `json:"title"`
	Body string `json:"body"`
	Author User `json:"author"`
}

var posts []Post


func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/test", test)
	router.HandleFunc("/posts", addPost).Methods("POST")
	router.HandleFunc("/posts", listAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	log.Fatal( http.ListenAndServe(":8080", router))
}

func test(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(struct {
		ID string
	}{"555"})
}

func listAllPosts(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
func getPost(w http.ResponseWriter, r *http.Request)  {
	idParam := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	
	id, err := strconv.Atoi(idParam)
	
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID connot convert to integer"))
		return
	}
	
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}
	post := posts[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func addPost(w http.ResponseWriter, r *http.Request)  {
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	posts = append(posts, newPost)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)
}

func updatePost(w http.ResponseWriter, r *http.Request)  {
	idParam := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	
	id, err := strconv.Atoi(idParam)
	
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID connot convert to integer"))
		return
	}
	
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}
	var updatePost Post
	json.NewDecoder(r.Body).Decode(&updatePost)
	posts[id] = updatePost
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatePost)
}

func patchPost(w http.ResponseWriter, r *http.Request)  {
	idParam := mux.Vars(r)["id"]
	
	id, err := strconv.Atoi(idParam)
	
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID connot convert to integer"))
		return
	}
	
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	post := posts[id]
	json.NewDecoder(r.Body).Decode(&post)
	posts[id] = post
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func deletePost(w http.ResponseWriter, r *http.Request)  {
	idParam := mux.Vars(r)["id"]
	
	id, err := strconv.Atoi(idParam)
	
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID connot convert to integer"))
		return
	}
	
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}	
	posts = append(posts[:id], posts[id+1:]...)
	w.WriteHeader(200)
}