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
		new_parent := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "krhtm2di",
			"name": "parent",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"collectionId": "tt0b8hchute5mz4",
				"cascadeDelete": false
			}
		}`), new_parent)
		collection.Schema.AddField(new_parent)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tt0b8hchute5mz4")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("krhtm2di")

		return dao.SaveCollection(collection)
	})
}
