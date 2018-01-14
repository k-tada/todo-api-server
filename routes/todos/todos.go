package todos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id          int    `form:"id" json:"id"`
	Text        string `form:"title" json:"title"`
	Status      string `form:"status" json:"status"`
	Description string `form:"description" json:"description"`
}

type Todos []Todo

func (t Todos) findById(id int) *Todo {
	for _, v := range t {
		if v.Id == id {
			return &v
		}
	}
	return nil
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

func Create(c *gin.Context) {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	c.JSON(http.StatusOK, *todos.findById(id))
}

func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "no implements"})
}
