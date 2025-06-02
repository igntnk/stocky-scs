package models

type User struct {
	Id             string `json:"id" bson:"id"`
	Name           string `json:"name" bson:"name"`
	Description    string `json:"description" bson:"description"`
	DocumentType   string `json:"document_type" bson:"document_type"`
	DocumentNumber string `json:"document_number" bson:"document_number"`
	CreationDate   string `json:"creation_date" bson:"creation_date"`
	Blocked        bool   `json:"blocked" bson:"blocked"`
	AuthId         string `json:"auth_id" bson:"auth_id"`
}
