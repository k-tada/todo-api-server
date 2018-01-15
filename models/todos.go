package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	// collections
	CollectionTodo = "todos"

	// todo statuses
	TodoStatus  string = "todo"
	DoingStatus string = "doing"
	DoneStatus  string = "done"
)

type Todo struct {
	Id          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string        `form:"title" binding:"required" json:"title" bson:"title"`
	Status      string        `form:"status" binding:"required" json:"status" bson:"status"`
	Description string        `form:"description" binding:"required" json:"description" bson:"description"`
	CreatedOn   int64         `json:"created_on" bson:"created_on"`
	UpdatedOn   int64         `json:"updated_on" bson:"updated_on"`
}

type Todos []Todo
