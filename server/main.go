package main

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type todo struct { // need to start with UPPERCASE to export the value
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Id: "1", Item: "abc", Completed: false},
	{Id: "2", Item: "abcd", Completed: false},
	{Id: "3", Item: "abcde", Completed: true},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	var id string = context.Param("id")
	t, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, t)
}

func addTodos(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, todos)
}

func main() {
	// create a server with gin
	router := gin.Default()
	// middlewares, cors
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://localhost:9000"},
		AllowMethods: []string{"GET", "POST", "PUT"},
	}))

	// routes
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodos)
	router.Run("localhost:9090") // path, port
}
