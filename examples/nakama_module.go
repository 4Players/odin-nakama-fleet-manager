package main

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"strconv"

	odinfleet "odinfleet-nakama"

	"github.com/heroiclabs/nakama-common/runtime"
)

// InitModule registers the ODIN fleet manager so Nakama can use it for matchmaker allocations.
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	appIDRaw := os.Getenv("ODIN_APP_ID")
	if appIDRaw == "" {
		return errors.New("missing ODIN_APP_ID")
	}
	appID, err := strconv.Atoi(appIDRaw)
	if err != nil {
		return err
	}
	apiToken := os.Getenv("ODIN_API_TOKEN")
	if apiToken == "" {
		return errors.New("missing ODIN_API_TOKEN")
	}

	cfg := odinfleet.GetDefaultConfiguration(int32(appID), apiToken)
	fm, err := odinfleet.CreateOdinFleetManager(logger, cfg, odinfleet.ServerToInstanceInfo)
	if err != nil {
		return err
	}

	if err := initializer.RegisterFleetManager(fm); err != nil {
		return err
	}
	logger.Info("ODIN fleet manager registered")
	return nil
}
