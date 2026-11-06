package main
import (
	"log"
	"net/http"
)
func main(){
	ConnectDb()
	defer DB.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/posts", CreatPost)
	mux.HandleFunc("GET /api/v1/posts", Getposts)
	mux.HandleFunc("DELETE /api/v1/posts/{id}", DeletePost)
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil{
      log.Fatal(err)
	}
}