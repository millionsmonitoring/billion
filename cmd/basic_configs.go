package cmd

import (
	"context"
	"fmt"

	"github.com/anurag925/billion/configs"
	_ "github.com/joho/godotenv/autoload"
	"github.com/millionsmonitoring/millionsgocore/env"
	"github.com/millionsmonitoring/millionsgocore/initializers"
	"github.com/millionsmonitoring/millionsgocore/logger"
)

// autoload env file
// load and parse the flags
// set env of the application
// loads the settings required for the application
// init logger
func BasicConfig(ctx context.Context) {
	flags := configs.ParseFlags()
	configs.SetCLIFlags(flags)
	env.Env(env.Environment(*flags.Env))
	logger.Init(logger.WithBlacklistKeys("password"))
	settings, err := initializers.LoadSettings[configs.Setting](ctx)
	if err != nil {
		panic(fmt.Sprintf("error in loading settings %s", err))
	}
	configs.SetSettings(settings)
}
