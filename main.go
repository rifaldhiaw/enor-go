package main

import (
	_ "enor-go/migrations"
	messageUtils "enor-go/utils"
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true, // auto creates migration files when making collection changes
	})

	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {

		isMessageRecord := e.Record.Collection().Name == "messages"

		if isMessageRecord {
			parentId := e.Record.GetString("parent")
			isReply := parentId != ""

			if isReply {
				messageUtils.UpdateReplySummary(app, e)
			} else {
				messageUtils.IncrementLastMessageClock(app, e)
			}

		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}
