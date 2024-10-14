package main

import (
	"fmt"
	"gyanasetu/backend/server"
	"os"

	"github.com/joho/godotenv"
)

const (
	DEFAULT_HOST = "127.0.0.1"
	DEFAULT_PORT = "5000"
)

func GetAddr() string {
	var host string = os.Getenv("GYANASETU_HOST")
	var port string = os.Getenv("GYANASETU_PORT")
	if host == "" {
		host = DEFAULT_HOST
	}
	if port == "" {
		port = DEFAULT_PORT
	}
	return fmt.Sprintf("%s:%s", host, port)
}

func main() {
	godotenv.Load()
	addr := GetAddr()
	fmt.Printf("Addr: %s\n", addr)
	server.StartServer(addr)
}
