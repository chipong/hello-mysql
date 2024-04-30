package main

import (
	db "github.com/chipong/hello-mysql/database"
)

func init() {
	err := db.SetBoilerDatabas()
	if err != nil {
		panic(err)
	}
}

func main() {
	
}