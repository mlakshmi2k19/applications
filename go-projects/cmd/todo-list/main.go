package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var tasks = []Task{
	{ID: 1, Title: "Task1", Content: "Description1"},
	{ID: 2, Title: "Task2", Content: "Description2"},
}

func getAllTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tasks)
}

func createTask(ctx *gin.Context) {
	var newTask Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	ctx.JSON(http.StatusCreated, newTask)
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getAllTasks)
	router.POST("/create-new-task", createTask)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %s", err.Error())
		}
	}()

	<-quit

	fmt.Println("\nShutting down the server...")
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Printf("Error shutting down the server: %s\n", err.Error())
	}
	fmt.Println("Server stopped...")
}
