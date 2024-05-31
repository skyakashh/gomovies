package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skyakashh/mongo/router"
)

func main() {
	fmt.Println("mongo db api")
	r := router.Router()
	fmt.Println("server is getting started ...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listen and serve 4000")

}

//mongodb+srv://sky:<password>@cluster0.g6gxgfm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0
