package services

import (
	"context"
	"fmt"
	_"math/rand"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/chipong/hello-mysql/models"
)

type Car struct {}

func NewCarService() Car {
	return Car{}
}

func (c Car) Register(ctx context.Context, name string) (int, error) {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("transaction begin failed, %w", err)
	}

	carId := int(uuid.New().ID())
	
	car := models.Car{
		ID: carId,
		Name: null.StringFrom(name),
	}

	err = car.Insert(ctx, tx, boil.Infer())
	if err != nil {
		if err := tx.Rollback(); err != nil {
			fmt.Printf("transaction rollback err %v\n", err)
		}
		return 0, fmt.Errorf("car register failed %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("transaction commit failed %w", err)
	}

	return carId, nil
}