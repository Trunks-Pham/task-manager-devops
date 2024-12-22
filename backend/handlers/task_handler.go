package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-manager/database"
	"task-manager/models"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rows, _ := database.DB.Query("SELECT id, title, done FROM tasks")
		defer rows.Close()

		var tasks []models.Task
		for rows.Next() {
			var t models.Task
			rows.Scan(&t.ID, &t.Title, &t.Done)
			tasks = append(tasks, t)
		}
		json.NewEncoder(w).Encode(tasks)

	case "POST":
		var t models.Task
		json.NewDecoder(r.Body).Decode(&t)
		_, err := database.DB.Exec("INSERT INTO tasks (title, done) VALUES ($1, $2)", t.Title, t.Done)
		if err != nil {
			http.Error(w, "Failed to insert task", http.StatusInternalServerError)
		}
		fmt.Fprintln(w, "Task added!")
	}
}
