// main_test.go

package main_test

import (
    "os"
    "testing"
    "log"
    "net/http"
    "net/http/httptest"
    "bytes"
    "encoding/json"
    "."
)

var a main.App

func TestMain(m *testing.M) {
    a = main.App{}
    a.Initialize(
        os.Getenv("TEST_DB_USERNAME"),
        os.Getenv("TEST_DB_PASSWORD"),
        os.Getenv("TEST_DB_NAME"),
        os.Getenv("TEST_DB_HOST"))

    ensureTableExists()

    code := m.Run()

    clearTable()

    os.Exit(code)
}
