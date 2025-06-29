package main

import (
	"mimic/modules/api"
	"mimic/modules/db"
	"mimic/modules/db/mimic"
	"mimic/modules/db/mimic/blocks"
	"mimic/modules/db/mimic/state"
)

func main() {
	dbClient := db.New()
	mimicDb := mimic.Init(dbClient)
	_ = blocks.New(mimicDb)
	_ = state.New(mimicDb)

	/*
		plugins := []aggregate.Plugin{
			// db,
			// mimicDb,
			hiveBlocks,
			stateDb,
		}

		agg := aggregate.New(plugins)


		agg.Init()
		agg.Start().Await(context.Background())
		defer agg.Stop()
	*/

	router := api.NewAPIServer()
	router.Init()
	router.Start()
}
