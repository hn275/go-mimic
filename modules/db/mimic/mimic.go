package mimic

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const dbPath = "go-mimic"

var mimicDb = &MimicDb{}

type MimicDb struct {
	*mongo.Database
}

func Init(d *mongo.Client) *MimicDb {
	mimicDb.Database = d.Database(dbPath)
	return mimicDb
}
