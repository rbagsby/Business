package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageurl"`
}

var db *sql.DB

func initDB() {
	var err error
	connString := "server=localhost;user id=sa;password=Pr0fessi0nal1!;database=Vetspertise"
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	fmt.Println("Connected to SQL Server")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT ID, Name, Description, Price, ImageURL FROM Products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.ImageURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/homePage.html"))
	tmpl.Execute(w, nil)
}

//func productsPageHandler(w http.ResponseWriter, r *http.Request) {
//tmpl := template.Must(template.ParseFiles("templates/products.html"))
//	tmpl.Execute(w, nil)
//}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/About.html"))
	tmpl.Execute(w, nil)
}

func shopHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/shop.html"))
	tmpl.Execute(w, nil)
}

func main() {
	initDB()
	defer db.Close()

	mediaDir := "/media"

	// Create a file server handler
	fs := http.FileServer(http.Dir(mediaDir))

	// Serve static files at /media URL path
	http.Handle("/media/", http.StripPrefix("/media/", fs))

	// Serve static files (CSS, JS, images)
	fss := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fss))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/products", productsHandler)
	//http.HandleFunc("/products-page", productsPageHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/shop", shopHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
