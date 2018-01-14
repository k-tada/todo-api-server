package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id          int    `form:"id" json:"id"`
	Text        string `form:"title" json:"title"`
	Status      string `form:"status" json:"status"`
	Description string `form:"description" json:"description"`
}

var todos = Todos{
	{1, "hoge", TodoStatus, "Hoge Todo"},
	{2, "fuga", DoingStatus, "Fuga Todo"},
	{3, "piyo", DoneStatus, "Piyo Todo"},
	{4, "mogo", DoneStatus, "Mogo Todo"},
}

const (
	TodoStatus  string = "todo"
	DoingStatus string = "doing"
	DoneStatus  string = "done"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}
