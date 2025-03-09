package main

import (
	"authentication/config"
	"authentication/routes"
	"fmt"
	"net/http"
)

func main() {
	config.ConnectDB()
	mux := http.NewServeMux()
	routes.SetupRoutes(mux)

	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		return
	}
}
