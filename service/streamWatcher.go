package service

import (
	"context"

	"github.com/benpate/derp"
	"github.com/benpate/ghost/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// StreamWatcher initiates a mongodb change stream to on every updates to Stream data objects
func StreamWatcher(uri string, database string) chan model.Stream {

	result := make(chan model.Stream)

	ctx := context.Background()

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		derp.Report(err)
		return result
	}

	if err := client.Connect(ctx); err != nil {
		derp.Report(err)
		return result
	}

	collection := client.Database(database).Collection("Stream")

	cs, err := collection.Watch(ctx, mongo.Pipeline{})

	if err != nil {
		derp.Report(derp.Wrap(err, "ghost.service.Watcher", "Unable to open Mongodb Change Stream"))
		return result
	}

	go func() {

		for cs.Next(ctx) {

			var event struct {
				Stream model.Stream `bson:"fullDocument"`
			}

			if err := cs.Decode(&event); err != nil {
				derp.Report(err)
				continue
			}

			// Skip "zero" sreams
			if event.Stream.StreamID.IsZero() {
				continue
			}

			result <- event.Stream
		}
	}()

	return result
}