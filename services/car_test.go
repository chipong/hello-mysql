package services

import (
	"context"
	"testing"
	"fmt"

	db "github.com/chipong/hello-mysql/database"
)

func init() {
	err := db.SetBoilerDatabas()
	if err != nil {
		panic(err)
	}
}

func TestCarRegist(t *testing.T) {
	carService := NewCarService()
	id, err := carService.Register(context.Background(), "firstCar")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Car regist Id : %d", id)
}