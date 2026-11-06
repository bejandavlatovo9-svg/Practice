package main

import (
	
	"encoding/json"
	"net/http"
	
)

func CreatPost(w http.ResponseWriter, r *http.Request){
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil{
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}
	err = DB.QueryRow(
       `INSERT INTO posts (title, body)
	   VALUES
	   ($1,$2)
	   RETURNING id`,
	   post.Title,
	   post.Body,
	).Scan(&post.ID)
	if err != nil{
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)

} 
func Getposts(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(
		`SELECT id, title, body FROM posts ORDER BY id`,)
	if err != nil{
		http.Error(w, "Databese error", 500)
	}
	defer rows.Close()
	posts := []Post{}
	for rows.Next(){
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Body,
		)
		if err != nil{
			http.Error(w, "Scan error", 500)
			return
		}
		posts = append(posts, post)
	}
 	w.Header().Set("Content-type","application/json")
    json.NewEncoder(w).Encode(posts)
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    result, err := DB.Exec(
        `DELETE FROM posts WHERE id = $1`, id,
    )
    if err != nil {
        http.Error(w, "Database error", 500)
        return
    }
    affected, err := result.RowsAffected()
    if err != nil {
        http.Error(w, "Error checking result", 500)
        return
    }
    if affected == 0 {
        http.Error(w, "Post not found", 404)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Post deleted successfully",
    })
}