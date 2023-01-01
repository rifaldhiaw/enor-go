package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("npbg8ayih40mdup")
		if err != nil {
			return err
		}

		// add
		new_lastMessageClock := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "pdvjfsgu",
			"name": "lastMessageClock",
			"type": "number",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), new_lastMessageClock)
		collection.Schema.AddField(new_lastMessageClock)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("npbg8ayih40mdup")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("pdvjfsgu")

		return dao.SaveCollection(collection)
	})
}
