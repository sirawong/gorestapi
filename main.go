package main

import (
	"log"
	"net/http"
	"postapi/app"
	"postapi/app/database"
)

func main() {
	app := app.New()
	app.DB = &database.DB{}
	err := app.DB.Open()
	check(err, "Cannot Connect")

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("App running...")
	err = http.ListenAndServe(":8000", nil)
	check(err, "Cannot Serve")
}

func check(e error, msg string) {
	if e != nil {
		log.Fatalln(msg+":", e)
	}
}
