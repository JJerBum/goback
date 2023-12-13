package main

import (
	"goback/entgo/config"
	"goback/entgo/middleware"
	"goback/entgo/router"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("faild to load for .env, err:", err)
		os.Exit(1)
	}
}

func main() {
	// TODO: 프로젝트 전체에 오류가 있음 고처야 함!
	port := os.Getenv("PORT")

	client, err := config.NewEntClient()
	if err != nil {
		log.Fatalln("err : %s", err.Error())
	}
	defer client.Close()

	if err != nil {
		log.Fatalln("Fail to initalize client")
	}

	config.SetClient(client)

	r := mux.NewRouter()
	r.Use(middleware.Header)
	router.RegisterRouter(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on port " + port)
	log.Fatal(srv.ListenAndServe())
}
