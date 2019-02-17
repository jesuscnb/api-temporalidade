package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Temporalidade struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Categoria   string        `bson:"categoria" json:"categoria"`
	Destinacao  string        `bson:"destinacao" json:"destinacao"`
	Tipo        string        `bson:"tipo" json:"tipo"`
	DataCriacao time.Time     `bson:"dataCriacao" json:"dataCriacao"`
	Tempo       int           `bson:"tempo" json:"tempo"`
	Observacao  string        `bson:"observacao" json:"observacao"`
}
