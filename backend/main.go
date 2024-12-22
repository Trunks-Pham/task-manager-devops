package main

import (
	"fmt"
	"log"
	"net/http"
	"task-manager/database"
	"task-manager/handlers"
)

func main() {
	// Kết nối database
	database.ConnectDB()

	// Đăng ký route
	http.HandleFunc("/tasks", handlers.TaskHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
