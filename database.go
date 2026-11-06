package main
import (
	"database/sql"
	"log"
	_"github.com/lib/pq"
)
var DB *sql.DB
func ConnectDb(){
	Connstr := "host=localhost port=5433 user=postgres password=BEZH0909 dbname=postgres sslmode=disable"
	db,err := sql.Open("postgres", Connstr)
	if err != nil{
		log.Fatal("Database connection error:", err)
	}
   err = db.Ping()
   if err != nil{
	log.Fatal("Database is not responding:",err)
   }
   DB = db
   log.Println("Database connected successfully!")
}