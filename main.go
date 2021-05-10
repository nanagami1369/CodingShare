package main

import (
	"fmt"

	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/module"
)

func main() {
	sm := module.NewStupModule()
	db, err := sm.OpenDB()
	if err != nil {
		panic(err)
	}
	user := &model.User{}
	db.First(user)
	fmt.Println("%v", user)
}
