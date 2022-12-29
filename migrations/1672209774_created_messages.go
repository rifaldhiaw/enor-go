package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "tt0b8hchute5mz4",
			"created": "2022-12-28 06:42:54.045Z",
			"updated": "2022-12-28 06:42:54.045Z",
			"name": "messages",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "nmn9bkta",
					"name": "body",
					"type": "json",
					"required": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "venx6bpd",
					"name": "user",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false
					}
				},
				{
					"system": false,
					"id": "ekgmgerj",
					"name": "reaction",
					"type": "json",
					"required": false,
					"unique": false,
					"options": {}
				}
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tt0b8hchute5mz4")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
