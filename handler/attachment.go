package handler

import (
	"github.com/benpate/derp"
	"github.com/benpate/ghost/model"
	"github.com/benpate/ghost/server"
	"github.com/benpate/steranko"
	"github.com/labstack/echo/v4"
)

func GetAttachment(factoryManager *server.FactoryManager) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		factory, err := factoryManager.ByContext(ctx)

		if err != nil {
			return derp.Wrap(err, "ghost.handler.GetAttachment", "Cannot load Domain")
		}

		var stream model.Stream
		streamService := factory.Stream()
		streamToken := getStreamToken(ctx)

		if err := streamService.LoadByToken(streamToken, &stream); err != nil {
			return derp.Wrap(err, "ghost.handler.GetAttachment", "Error loading Stream", streamToken)
		}

		// Try to find the action requested by the user.  This also enforces user permissions...
		sterankoContext := ctx.(*steranko.Context)
		if _, err := factory.Renderer(sterankoContext, &stream, "view"); err != nil {
			return derp.Wrap(err, "ghost.handler.GetAttachment", "Cannot create renderer")
		}

		attachmentService := factory.Attachment()
		attachment, err := attachmentService.LoadByToken(ctx.Param("attachment"))

		if err != nil {
			return derp.Wrap(err, "ghost.handler.GetAttachment", "Error loading attachment")
		}

		file, err := attachmentService.File(&attachment)

		if err != nil {
			return derp.Wrap(err, "ghost.handler.GetAttachment", "Error accessing attachment file")
		}

		return ctx.Stream(200, attachment.MimeType(), file)
	}
}