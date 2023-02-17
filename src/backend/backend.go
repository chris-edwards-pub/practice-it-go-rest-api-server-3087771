package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type App struct{
	DB   *sql.DB
	Port string
}

func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func (a *App) Initialize() {
	DB, err := sql.Open("sqlite3", "../../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	a.DB = DB
}

func (a *App) Run(){
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server stated and listening on port ", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, nil))
}

