// app.go

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"log"
	"net/http"
	"runtime"
	"strconv"
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

//All routes for app
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/hello/:{name:[a-zA-Z]+}", a.getName).Methods("GET")
	a.Router.HandleFunc("/health", a.getHardwareData).Methods("GET")
//	a.Router.HandleFunc("/hello/:{name:[a-zA-Z]+}", a.createName).Methods("POST")
	a.Router.HandleFunc("/counts", a.getNames).Methods("GET")
//	a.Router.HandleFunc("/hello/:{name:[a-zA-Z]+}", a.updateCount).Methods("PUT")
	a.Router.HandleFunc("/delete", a.deleteCounts).Methods("DELETE")
}

//Show name

func (a *App) getName(w http.ResponseWriter, r *http.Request) {
      vars := mux.Vars(r)
      name := vars["name"]
      fmt.Fprintf(w, "Hello, %s!", name)
      n := namest{Name: name}
      if nomatch := n.getName(a.DB); nomatch !=nil {
            n.createName(a.DB)
         }
     defer r.Body.Close()
     n.updateCount(a.DB)
}      
//Creaet a name
//func (a *App) createName(w http.ResponseWriter, r *http.Request) {
//	var n namest
//	decoder := json.NewDecoder(r.Body)
//	if err := decoder.Decode(&n); err != nil {
//		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
//		return
//	}
//	defer r.Body.Close()
//
//	if err := n.createName(a.DB); err != nil {
//		respondWithError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	respondWithJSON(w, http.StatusCreated, n)
//}

//Update count on existing name
//func (a *App) updateCount(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	name := vars["name"]
//	var n namest
//	decoder := json.NewDecoder(r.Body)
//	if err := decoder.Decode(&n); err != nil {
//		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
//		return
//	}
//	defer r.Body.Close()
//
//	n.Name = name
//
//	if err := n.updateCount(a.DB); err != nil {
//		respondWithError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	respondWithJSON(w, http.StatusOK, n)
//}

//Get Names and Counts
func (a *App) getNames(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 100 || count < 1 {
		count = 100
	}
	if start < 0 {
		start = 0
	}

	names, err := getNames(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, names)
}

//Delete all counts
func (a *App) deleteCounts(w http.ResponseWriter, r *http.Request) {
	var n namest
	if err := n.deleteCounts(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

//Health Check With Hardware Info
func (a *App) getHardwareData(w http.ResponseWriter, r *http.Request) {
	runtimeOS := runtime.GOOS
	// memory
	vmStat, err := mem.VirtualMemory()
	retErr(err)

	// disk
	diskStat, err := disk.Usage("/")
	retErr(err)

	html := "<html>OS : " + runtimeOS + "<br>"
	html = html + "Total memory: " + strconv.FormatUint(vmStat.Total, 10) + " bytes <br>"
	html = html + "Free memory: " + strconv.FormatUint(vmStat.Free, 10) + " bytes<br>"
	html = html + "Percentage used memory: " + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%<br>"

	html = html + "Total disk space: " + strconv.FormatUint(diskStat.Total, 10) + " bytes <br>"
	html = html + "Used disk space: " + strconv.FormatUint(diskStat.Used, 10) + " bytes<br>"
	html = html + "Free disk space: " + strconv.FormatUint(diskStat.Free, 10) + " bytes<br>"
	html = html + "Percentage disk space usage: " + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%<br>"

	html = html + "</html>"

	w.Write([]byte(html))
}

////Error Handler for json
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

//JSON Handler for Marshalling
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//Error Handler for health check
func retErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
