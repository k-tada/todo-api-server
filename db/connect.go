package db

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	Mongo   *mgo.DialInfo
)

const (
	MongoDBUrl = "mongodb://mongo/todo-api-server"
)

func Connect() {
	url := os.Getenv("MONGODB_URL")
	if len(url) == 0 {
		url = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(url)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}

	s, err := mgo.Dial(url)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}

	// safety mode => acknowledged
	s.SetSafe(&mgo.Safe{})

	fmt.Println("Connected to MongoDB, " + url)
	Session = s
	Mongo = mongo
}
