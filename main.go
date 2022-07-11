package main

import (
	"Library-API/config"
	"fmt"
	"Library-API/router"
	"log"
	"net/http"
)

func main(){
	config.Initialize()
	fmt.Println(config.ConnectionString)
	r := router.Create()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}