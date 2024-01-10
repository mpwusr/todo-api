package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Clean Room", Completed: false},
	{ID: "3", Item: "Clean Room", Completed: false},
}

func GetTodos(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, todos)
}
