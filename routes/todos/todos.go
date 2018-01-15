package todos

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"todo-api-server/models"
)

func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	todos := []models.Todo{}
	err := db.C(models.CollectionTodo).Find(nil).Sort("-_id").All(&todos)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, todos)
}

func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	todo := models.Todo{}
	err := c.Bind(&todo)
	if err != nil {
		c.Error(err)
		return
	}

	todo.CreatedOn = time.Now().UnixNano() / int64(time.Millisecond)
	todo.UpdatedOn = time.Now().UnixNano() / int64(time.Millisecond)

	err = db.C(models.CollectionTodo).Insert(todo)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func Show(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	todo := models.Todo{}
	oId := bson.ObjectIdHex(c.Param("_id"))

	err := db.C(models.CollectionTodo).FindId(oId).One(&todo)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	todo := models.Todo{}
	err := c.Bind(&todo)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	doc := bson.M{
		"title":       todo.Title,
		"status":      todo.Status,
		"description": todo.Description,
		"updated_on":  time.Now().UnixNano() / int64(time.Millisecond),
	}

	err = db.C(models.CollectionTodo).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
