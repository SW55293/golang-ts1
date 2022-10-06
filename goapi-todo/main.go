package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type todo struct {
	ID        string	`json:"id"`
	Item      string	`json:"item"`
	Completed bool		`json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Wash Clothes", Completed: false},
	{ID: "3", Item: "Mop", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos) 
}

func addTodos(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return 
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
	 
}

func getTodoId(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	} 
	return nil, errors.New("todo not found")
}

func getOneTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoId(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found sorry"})
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoId(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found sorry"})
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getOneTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodos)
	router.Run("localhost:9090")
}