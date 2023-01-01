package main

import (
	"encoding/json"
	_ "enor-go/migrations"
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"golang.org/x/exp/slices"
)

type ReplySummary struct {
	Count int      `json:"count"`
	Users []string `json:"users"`
}

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true, // auto creates migration files when making collection changes
	})

	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {

		// if collection is "messages" and parent field is not empty,
		// update parent record
		if e.Record.Collection().Name == "messages" && e.Record.Get("parent") != nil {
			updateReplySummary(app, e)
		}

		// if collection is "messages" and parent field is empty
		// then increment lastMessageClock in it's channel record
		if e.Record.Collection().Name == "messages" && e.Record.Get("parent") == nil {
			channelID := e.Record.Get("channel").(string)
			channelRecord, err := app.Dao().FindRecordById("channels", channelID)
			if err != nil {
				return err
			}

			lastMessageClock := channelRecord.GetInt("lastMessageClock")

			channelRecord.Set("lastMessageClock", lastMessageClock+1)
			err = app.Dao().SaveRecord(channelRecord)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}

// update replySummary function
func updateReplySummary(app *pocketbase.PocketBase, e *core.RecordCreateEvent) error {
	record := e.Record

	// get parent record
	parentID := e.Record.Get("parent").(string)
	parentRecord, err := app.Dao().FindRecordById("messages", parentID)
	if err != nil {
		return err
	}

	// update replySummary field,
	// which is a json object containing the count of replies and a list of users who have replied
	var replySummary ReplySummary
	err = parentRecord.UnmarshalJSONField("replySummary", &replySummary)

	// if reply summary is empty, create new
	// otherwise, unmarshal existing reply summary
	if err != nil {
		replySummary = ReplySummary{
			Count: 0,
			Users: []string{},
		}
	}

	// increment reply count
	replySummary.Count++

	// append user to list if not already in list
	user := record.Get("user")
	if user != nil {
		userStr, ok := user.(string)
		if ok && !slices.Contains(replySummary.Users, userStr) {
			replySummary.Users = append(replySummary.Users, userStr)
		}
	}

	// marshal reply summary
	newReplySummaryBytes, err := json.Marshal(replySummary)
	if err != nil {
		log.Println("error marshaling reply summary", err)
		return err
	}

	// update parent record
	parentRecord.Set("replySummary", newReplySummaryBytes)
	err = app.Dao().SaveRecord(parentRecord)
	if err != nil {
		log.Println("error updating parent record", err)
		return err
	}

	return nil
}
