package config

import (
	"context"
	"fmt"
	"goback/entgo/ent"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	client *ent.Client
)

func GetClient() *ent.Client {
	return client
}

func SetClient(newClient *ent.Client) {
	client = newClient
}
func NewEntClient() (*ent.Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	client, err := ent.Open("mysql", dsn, ent.Debug(), ent.Log(func(i ...any) {
		for _, v := range i {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), v)
			fmt.Print("\n")
		}
	}))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	// Run auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, err
}
