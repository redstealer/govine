package main

import (
    "log"
    "net/http"
    "hexakleo/govine"
)

func main() {
    app := govine.NewApp()

    app.Get("/", func(c *govine.Context) {
        c.String(http.StatusOK, "Welcome to Govine!")
    })

    log.Fatal(http.ListenAndServe(":8080", app))
}
