// main_test.go

package main_test

import (
    "os"
    "testing"
    "log"
    "net/http"
    "net/http/httptest"
    "bytes"
//    "encoding/json"
    "."
)

var a main.App

func TestMain(m *testing.M) {
    a = main.App{}
    a.Initialize(
        os.Getenv("APP_DB_USERNAME"),
        os.Getenv("APP_DB_PASSWORD"),
        os.Getenv("APP_DB_NAME"),
        os.Getenv("APP_DB_HOST"))

    ensureTableExists()

    code := m.Run()

    clearTable()

    os.Exit(code)


}

func ensureTableExists() {
    if _, err := a.DB.Exec(tableCreationQuery); err != nil {
        log.Fatal(err)
    }
}

func clearTable() {
    a.DB.Exec("DELETE FROM names")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS names
(
    name TEXT NOT NULL,
    count INT NOT NULL DEFAULT 0,
    CONSTRAINT names_pkey PRIMARY KEY (name)
)`

func TestEmptyTable(t *testing.T) {
    clearTable()

    req, _ := http.NewRequest("GET", "/counts", nil)
    response := executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)

    if body := response.Body.String(); body != "[]" {
        t.Errorf("Expected an empty array. Got %s", body)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)
    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

func TestCreateName(t *testing.T) {
    clearTable()

    payload := []byte("testme")

    req, _ := http.NewRequest("GET", "/hello/:", bytes.NewBuffer(payload))
    response := executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)

}

func TestGetNames(t *testing.T) {
    clearTable()

    req, _ := http.NewRequest("GET", "/counts", nil)
    response := executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)

}
