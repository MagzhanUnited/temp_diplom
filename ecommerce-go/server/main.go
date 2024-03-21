package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/rs/cors"
	middlewares "github.com/umangraval/Go-Mongodb-REST-boilerplate/handlers"
	"github.com/umangraval/Go-Mongodb-REST-boilerplate/routes"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	port := middlewares.DotEnvVariable("PORT")
	color.Cyan("🌏 Server running on 89.46.33.148:" + port)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	router := routes.Routes()
	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Origin", "Accept", "*"},
		AllowedOrigins:   []string{"http://89.46.33.148:2001", "http://89.46.33.148:2003"},
		AllowCredentials: true, // Enable credentials
	})

	handler := c.Handler(router)
	http.ListenAndServe("0.0.0.0:"+port, middlewares.LogRequest(handler))
}
