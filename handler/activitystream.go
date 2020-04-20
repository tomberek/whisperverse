package handler

import (
	"github.com/benpate/activitystream/writer"
	"github.com/benpate/ghost/service"
	"github.com/labstack/echo/v4"
)

// GetInbox returns an inbox for a particular ACTOR
func GetInbox(factory service.FactoryMaker) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		w := writer.Block(
			writer.Person("John Connor", "en"),
			writer.Article().Summary("hey-oh", "en").Icon("http://www.blah.com"),
		)

		return ctx.JSON(200, w)
	}
}

// PostInbox accepts messages to a particular ACTOR
func PostInbox(factory service.FactoryMaker) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		return nil
	}
}

// GetOutbox returns an inbox for a particular ACTOR
func GetOutbox(factory service.FactoryMaker) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		return nil
	}
}

// PostOutbox accepts messages to a particular ACTOR
func PostOutbox(factory service.FactoryMaker) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		return nil
	}
}
