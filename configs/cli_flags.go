package configs

import (
	"flag"
	"log/slog"
	"os"
)

var cliFlags CLIFlag

type CLIFlag struct {
	AppName *string
	Env     *string
	Port    *int
}

func CLIFlags() CLIFlag {
	return cliFlags
}

func SetCLIFlags(flags CLIFlag) {
	cliFlags = flags
}

func ParseFlags() CLIFlag {
	slog.Info("Parsing cli flags ...")
	cliDetails := CLIFlag{}
	cliDetails.AppName = flag.String("a", fromENVorDefault("APP_NAME", "billion").(string), "Name of the application")
	cliDetails.Env = flag.String("e", fromENVorDefault("ENV", "development").(string), "Environment in which the application will work")
	cliDetails.Port = flag.Int("p", fromENVorDefault("PORT", 1323).(int), "Port on which the application will work")
	flag.Parse()
	slog.Info("Successfully parsed cli flags ...", slog.Any("flags", cliDetails))
	return cliDetails
}

func fromENVorDefault(env string, defaultVal any) any {
	envData := os.Getenv(env)
	if envData != "" {
		return envData
	}
	return defaultVal
}
