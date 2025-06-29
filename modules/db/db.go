package db

import (
	"context"
	"log"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbClient *mongo.Client
)

func init() {
	ctx, cancel := context.WithTimeout(context.WithoutCancel(context.Background()), time.Second*5)
	defer cancel()

	dbOpts := options.BSONOptions{
		UseJSONStructTags:       true,
		ErrorOnInlineDuplicates: false,
		IntMinSize:              false,
		NilMapAsEmpty:           false,
		NilSliceAsEmpty:         false,
		NilByteSliceAsEmpty:     false,
		OmitZeroStruct:          false,
		StringifyMapKeysWithFmt: false,
		AllowTruncatingDoubles:  false,
		BinaryAsSlice:           false,
		DefaultDocumentD:        false,
		DefaultDocumentM:        false,
		UseLocalTimeZone:        false,
		ZeroMaps:                false,
		ZeroStructs:             false,
	}

	dbClientOpt := options.Client().ApplyURI(DbURI).SetBSONOptions(&dbOpts)

	var err error
	dbClient, err = mongo.Connect(ctx, dbClientOpt)
	if err != nil {
		log.Fatal(err)
	}

	err = dbClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Connected to MongoDB.", "url", DbURI)
}

func New() *mongo.Client {
	return dbClient
}
