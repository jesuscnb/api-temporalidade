package dao

import (
	"log"

	. "github.com/jesuscnb/api-temporalidade/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TemporalidadesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "temporalidades"
)

func (t *TemporalidadesDAO) Connect() {
	session, err := mgo.Dial(t.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(t.Database)
}

func (t *TemporalidadesDAO) FindAll() ([]Temporalidade, error) {
	var temporalidades []Temporalidade
	err := db.C(COLLECTION).Find(bson.M{}).All(&temporalidades)
	return temporalidades, err
}

func (m *TemporalidadesDAO) FindById(id string) (Temporalidade, error) {
	var temporalidade Temporalidade
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&temporalidade)
	return temporalidade, err
}

func (m *TemporalidadesDAO) Insert(temporalidade Temporalidade) error {
	err := db.C(COLLECTION).Insert(&temporalidade)
	return err
}

func (m *TemporalidadesDAO) Delete(temporalidade Temporalidade) error {
	err := db.C(COLLECTION).Remove(&temporalidade)
	return err
}

func (m *TemporalidadesDAO) Update(temporalidade Temporalidade) error {
	err := db.C(COLLECTION).UpdateId(temporalidade.ID, &temporalidade)
	return err
}
