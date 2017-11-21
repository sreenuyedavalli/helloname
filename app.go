// app.go

package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
     "runtime"
    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

//App and DB Connection
func (a *App) Initialize(user, password, dbname, host string) {
    connectionString :=
        fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", user, password, dbname, host)

    var err error
    a.DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }

    a.Router = mux.NewRouter()
    a.initializeRoutes()
}

//Listen and Servce Requests
func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(":8000", a.Router))
}

//all routes for app
func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/hello/:{name:[a-zA-Z]+}", a.showName).Methods("GET")
    a.Router.HandleFunc("/health", a.getHardwareData).Methods("GET")
    a.Router.HandleFunc("/hello/:{name:[a-zA-Z]+}", a.createName).Methods("POST")
    a.Router.HandleFunc("/counts", a.getNames).Methods("GET")
    a.Router.HandleFunc("/hello/:{name:[a-zA-Z]+}", a.updateCount).Methods("PUT")
    a.Router.HandleFunc("/delete", a.deleteCounts).Methods("DELETE")
}
