package controller

import "go.mongodb.org/mongo-driver/bson/primitive"

type Anggota struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nama       string             `json:"nama,omitempty"`
	Email      string             `json:"email,omitempty"`
	Password   string             `json:"password,omitempty"`
	IdKoperasi string             `json:"idKoperasi,omitempty"`
	LinkFoto   string             `json:"link_foto,omitempty"`
	Rule       []FormatRule       `json:"rule,omitempty"`
}

type FormatRule struct {
	Judul string `json:"judul,omitempty"`
	Value string `json:"value,omitempty"`
}

type Id struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}
