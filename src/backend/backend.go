package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

type App struct{
	DB     *sql.DB
	Port   string
	Router *mux.Router
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
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(){
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server stated and listening on port ", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, nil))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", helloWorld)
}