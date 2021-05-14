package main

import (
	// "net/http"

	"github.com/nanagami1369/CodingShare/module"
)

func main() {
	sm := module.NewStupModule()
	router, err := sm.GetRouter()
	if err != nil {
		panic(err)
	}
	router.Run(":8081")
}
