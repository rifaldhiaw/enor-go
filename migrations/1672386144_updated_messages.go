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
		edit_body := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "nmn9bkta",
			"name": "body",
			"type": "json",
			"required": true,
			"unique": false,
			"options": {}
		}`), edit_body)
		collection.Schema.AddField(edit_body)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tt0b8hchute5mz4")
		if err != nil {
			return err
		}

		// update
		edit_body := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "nmn9bkta",
			"name": "body",
			"type": "json",
			"required": false,
			"unique": false,
			"options": {}
		}`), edit_body)
		collection.Schema.AddField(edit_body)

		return dao.SaveCollection(collection)
	})
}
