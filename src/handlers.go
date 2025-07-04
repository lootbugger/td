package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) addTask(c *gin.Context) {
	body := c.Query("body")
	if body == "" {
		c.String(http.StatusBadRequest, "Missing 'body' param")
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'body' param"})
		return
	}

	_, err := app.tasks.DB.Exec("INSERT INTO tasks (body) VALUES (?)", body)
	if err != nil {
		c.String(http.StatusInternalServerError, "DB insert error")
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert error"})
		return
	}

	c.String(http.StatusOK, "Task added")
	// c.JSON(http.StatusOK, gin.H{"message": "Task added", "body": body})
}

func (app *application) listTasks(c *gin.Context) {
	rows, err := app.tasks.DB.Query("SELECT id, body FROM tasks")
	if err != nil {
		c.String(500, "DB querry error: %v\n", err)
		return
	}
	defer rows.Close()

	output := "ID | TASK\n"

	for rows.Next() {
		var id int
		var body string

		if err := rows.Scan(&id, &body); err != nil {
			c.String(500, "Failed to scan row: %v\n", err)
			return
		}
		output += fmt.Sprintf("%d | %s\n", id, body)
	}

	c.Data(200, "text/plain; sharset=utf-8", []byte(output))
}
