package api

import (
	"net/http"

	db "github.com/amupxm/todo_api/db/sqlc"
	"github.com/amupxm/todo_api/util"
	"github.com/gin-gonic/gin"
)

type CreateTodoRequest struct {
	Task  string `json:"task" binding:"required"`
	Title string `json:"title"`
}

type ListTodosRequest struct {
	Limit  int32 `form:"limit"`
	Offset int32 `form:"offset"`
}

func (s *Server) CreateNewTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	todo := db.CreateTodoParams{
		Task:  util.ToSqlNullString(req.Task),
		Title: req.Title,
	}
	createdTodo, err := s.store.CreateTodo(c, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, createdTodo)
}

func (s *Server) GetAllTodos(c *gin.Context) {
	var req ListTodosRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := db.ListTodosParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	if req.Limit == 0 {
		params.Limit = 10
	}
	if req.Offset == 0 {
		params.Offset = 1
	}
	todos, err := s.store.ListTodos(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, todos)
}
