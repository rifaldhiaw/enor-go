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
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2022-12-25 06:07:47.078Z",
				"updated": "2022-12-27 04:30:33.315Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif"
							],
							"thumbs": null
						}
					},
					{
						"system": false,
						"id": "rdxnydyq",
						"name": "organization",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "qs1w0clyttknnz5",
							"cascadeDelete": false
						}
					},
					{
						"system": false,
						"id": "uuze3g9m",
						"name": "appRole",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"admin",
								"member"
							]
						}
					}
				],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": false
				}
			},
			{
				"id": "2x7xas5ydcwxxmr",
				"created": "2022-12-27 03:50:13.087Z",
				"updated": "2022-12-27 04:32:49.494Z",
				"name": "teams",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ytk2tusw",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "jmbo0chj",
						"name": "organization",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "qs1w0clyttknnz5",
							"cascadeDelete": true
						}
					}
				],
				"listRule": "organization = @request.auth.organization.id",
				"viewRule": "organization = @request.auth.organization.id",
				"createRule": "organization = @request.auth.organization.id",
				"updateRule": "organization = @request.auth.organization.id",
				"deleteRule": "organization = @request.auth.organization.id && @request.auth.appRole = \"admin\"",
				"options": {}
			},
			{
				"id": "npbg8ayih40mdup",
				"created": "2022-12-27 03:50:41.317Z",
				"updated": "2022-12-27 10:51:00.447Z",
				"name": "channels",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "zbovo78t",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "0pi3wusu",
						"name": "team",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "2x7xas5ydcwxxmr",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "qlq20q5i",
						"name": "type",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"textRoom",
								"voiceRoom",
								"drawBoard",
								"document",
								"kanban"
							]
						}
					}
				],
				"listRule": "@request.auth.organization.id = team.organization.id",
				"viewRule": "@request.auth.organization.id = team.organization.id",
				"createRule": "@request.auth.organization.id = team.organization.id",
				"updateRule": "@request.auth.organization.id = team.organization.id",
				"deleteRule": "@request.auth.organization.id = team.organization.id && @request.auth.appRole = \"admin\"",
				"options": {}
			},
			{
				"id": "qs1w0clyttknnz5",
				"created": "2022-12-27 04:21:59.573Z",
				"updated": "2022-12-27 04:30:16.174Z",
				"name": "organizations",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "utvlfn5f",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": null,
				"viewRule": "id = @request.auth.organization.id",
				"createRule": "id = @request.auth.organization.id && @request.auth.appRole = \"admin\"",
				"updateRule": "id = @request.auth.organization.id && @request.auth.appRole = \"admin\"",
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
