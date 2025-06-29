package state

import (
	"mimic/modules/db/mimic"

	"go.mongodb.org/mongo-driver/mongo"
)

type StateDb struct {
	*mongo.Collection
}

func New(d *mimic.MimicDb) StateDb {
	return StateDb{d.Collection("state")}
}

func (s *StateDb) GetState(key string) (string, error) {
	// Implement the logic to get the state by key
	return "", nil
}

func (s *StateDb) SetState(key string, value string) error {
	// Implement the logic to set the state by key
	return nil
}
