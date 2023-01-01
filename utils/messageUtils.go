package messageUtils

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"golang.org/x/exp/slices"
)

type ReplySummary struct {
	Count int      `json:"count"`
	Users []string `json:"users"`
}

// update replySummary function
func UpdateReplySummary(app *pocketbase.PocketBase, e *core.RecordCreateEvent) error {
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
		return err
	}

	// update parent record
	parentRecord.Set("replySummary", newReplySummaryBytes)
	err = app.Dao().SaveRecord(parentRecord)
	if err != nil {
		return err
	}

	return nil
}

func IncrementLastMessageClock(app *pocketbase.PocketBase, e *core.RecordCreateEvent) error {
	record := e.Record

	channelID := record.Get("channel").(string)
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

	return nil
}
