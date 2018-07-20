package db

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

const mongoAddr = "192.168.9.81:27017"

var (
	db      *mgo.Session
	session *mgo.Database
)

// db.Collection("users").find("*")

// Env contains the database environment
// type Env struct {
// 	Session    *mgo.Database
// 	Collection *mgo.Collection
// }

// Connect connects to a specified database and initializes internally
func Connect() error {
	if db != nil {
		// return fmt.Errorf("session has already been started")
		return nil
	}

	db, err := mgo.Dial(mongoAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to mongoDB: %v", err)
	}

	session = db.DB("botManager")

	return nil
}

// Collection returns a queriable *mgo.Collection
func Collection(collection string) *mgo.Collection {
	return session.C(collection)
}

// Connect connects to a specified database and collection
// func Connect(database, collection string) (Env, error) {
// 	env := Env{}

// 	var err error
// 	env.Session, err = mgo.Dial(mongoAddr)
// 	if err != nil {
// 		return env, fmt.Errorf("failed to connect to mongoDB: %v", err)
// 	}

// 	env.Collection = env.Session.DB(database).C(collection)

// 	return env, nil
// }
