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

		// update
		edit_replySummary := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "eq9rrtsl",
			"name": "replySummary",
			"type": "json",
			"required": false,
			"unique": false,
			"options": {}
		}`), edit_replySummary)
		collection.Schema.AddField(edit_replySummary)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tt0b8hchute5mz4")
		if err != nil {
			return err
		}

		// update
		edit_replySummary := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "eq9rrtsl",
			"name": "replySummary",
			"type": "json",
			"required": true,
			"unique": false,
			"options": {}
		}`), edit_replySummary)
		collection.Schema.AddField(edit_replySummary)

		return dao.SaveCollection(collection)
	})
}
