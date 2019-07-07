package main

import (
	"os"
    "app/route"
    "app/datastore"
)

func main() {
	e := route.Init()
	
	datastore.Init()

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":"+port))
}

