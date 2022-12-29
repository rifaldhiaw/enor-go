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

		collection, err := dao.FindCollectionByNameOrId("tt0b8hchute5mz4")
		if err != nil {
			return err
		}

		// add
		new_channel := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "bgi4vhxh",
			"name": "channel",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"collectionId": "npbg8ayih40mdup",
				"cascadeDelete": false
			}
		}`), new_channel)
		collection.Schema.AddField(new_channel)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tt0b8hchute5mz4")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("bgi4vhxh")

		return dao.SaveCollection(collection)
	})
}
