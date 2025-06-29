package db

import (
	"errors"
	"mimic/lib/utils"
)

var (
	ErrEmptyURI = errors.New("empty MongoDB URI")
	DbURI       = utils.EnvOrDefault("MONGO_URL", "mongodb://localhost:27017")
)
