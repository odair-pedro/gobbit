package main

import (
	"context"
	"github.com/odair-pedro/gobbit/retry/domain"
	"github.com/odair-pedro/gobbit/retry/infrastructure"
	"github.com/odair-pedro/gobbit/retry/infrastructure/persistence/mongo"
	"log"
)

func main() {
	ctx := context.Background()

	rep, err := mongo.NewEventRepository(ctx, infrastructure.NewDatabase())
	if err != nil {
		log.Printf(err.Error())
	}

	event := new(domain.Event)
	err = rep.Create(ctx, event)
	if err != nil {
		log.Printf(err.Error())
	}

	result := rep.FindById(ctx, event.Id)
	log.Println(result.Id)

	result.Count++
	rep.Update(ctx, result)

	rep.Delete(ctx, result.Id)
}
