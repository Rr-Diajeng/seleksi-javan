package main

import (
	"os"
	"seleksi-javan/server"

	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("APP_ENV") == "production" {
		return
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

}

func main() {
	r := server.Start()
	r.Run(":8080")
}
