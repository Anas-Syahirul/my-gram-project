package main

import (
	"log"
	"my-gram-project/database"
	"my-gram-project/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	log.Println("starting app...")
	r.Run(":8080")
}
