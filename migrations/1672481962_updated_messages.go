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
		new_replySummary := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "eq9rrtsl",
			"name": "replySummary",
			"type": "json",
			"required": false,
			"unique": false,
			"options": {}
		}`), new_replySummary)
		collection.Schema.AddField(new_replySummary)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tt0b8hchute5mz4")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("eq9rrtsl")

		return dao.SaveCollection(collection)
	})
}
