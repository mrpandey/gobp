package util

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	_ "time/tzdata"

	"github.com/mrpandey/gobp/src/util/projectpath"

	"github.com/joho/godotenv"
)

type Config struct {
	Name     string
	DGN      DGNType
	LogLevel string

	Host     string
	RESTPort uint
	GRPCPort uint

	PGHost    string
	PGDB      string
	PGUser    string
	PGPass    string
	PGPort    uint
	PgSslMode string

	SqlIdleConnectionsCount int32
	SqlMaxConnectionsCount  int32
	SqlConnectionLifetime   time.Duration

	SecretKey string

	ExpiryTimezone              *time.Location
	DefaultCreditExpiryInMonths uint
	DefaultRefundExpiryInMonths uint

	ContextTimeout       time.Duration
	ContextTimeoutBuffer time.Duration
}

type DGNType string

const (
	Local DGNType = "local"
	Dev   DGNType = "dev"
	Prod  DGNType = "prod"
)

func NewConfig() *Config {
	envPath := projectpath.Root + "/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("Unable to load .env file. Falling back to environment variables.")
	}

	config := getDefaultConfig()
	environmentConfigurableVars := getEnvVarNames()

	err = setConfigFieldsFromEnv(environmentConfigurableVars, config)
	if err != nil {
		panic(err)
	}

	return config
}

func getDefaultConfig() *Config {
	ISTLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}

	return &Config{
		Name:     "gobp",
		DGN:      Local,
		LogLevel: "info", // debug, info, warn, error, dpanic, panic, fatal

		RESTPort: 3000,
		GRPCPort: 3001,

		PGHost: "localhost",
		PGDB:   "gobp",
		PGUser: "postgres",
		PGPass: "postgres",
		PGPort: 5432,

		SqlIdleConnectionsCount: 50,
		SqlMaxConnectionsCount:  50,
		SqlConnectionLifetime:   time.Hour,

		ExpiryTimezone:              ISTLocation,
		DefaultCreditExpiryInMonths: 6,
		DefaultRefundExpiryInMonths: 6,

		ContextTimeout:       time.Second * 5,
		ContextTimeoutBuffer: time.Millisecond,
	}
}

func getEnvVarNames() []string {
	return []string{
		"DGN",
		"LOG_LEVEL",

		"REST_PORT",
		"GRPC_PORT",

		"POSTGRES_HOST",
		"POSTGRES_DB",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_PORT",
		"POSTGRES_SSL_MODE",

		"SECRET_KEY",
	}
}

var errSecretKeyNotSet error = fmt.Errorf("secret key is not set")

//nolint:cyclop,gocognit
func setConfigFieldsFromEnv(environmentConfigurableVars []string, config *Config) error {
	for _, env := range environmentConfigurableVars {
		if os.Getenv(env) != "" {
			switch env {
			case "DGN":
				config.DGN = DGNType((os.Getenv(env)))
			case "LOG_LEVEL":
				config.LogLevel = os.Getenv(env)
			case "REST_PORT":
				RESTPort, err := strconv.Atoi(os.Getenv(env))
				if err != nil {
					return err
				}
				config.RESTPort = uint(RESTPort)
			case "GRPC_PORT":
				GRPCPort, err := strconv.Atoi(os.Getenv(env))
				if err != nil {
					return err
				}
				config.GRPCPort = uint(GRPCPort)
			case "POSTGRES_HOST":
				config.PGHost = os.Getenv(env)
			case "POSTGRES_DB":
				config.PGDB = os.Getenv(env)
			case "POSTGRES_USER":
				config.PGUser = os.Getenv(env)
			case "POSTGRES_PASSWORD":
				config.PGPass = os.Getenv(env)
			case "POSTGRES_PORT":
				DBport, err := strconv.Atoi(os.Getenv(env))
				if err != nil {
					return err
				}
				config.PGPort = uint(DBport)
			case "POSTGRES_SSL_MODE":
				config.PgSslMode = os.Getenv(env)
			case "SECRET_KEY":
				secretKey := os.Getenv(env)
				if secretKey == "" {
					return errSecretKeyNotSet
				}
				config.SecretKey = secretKey
			}
		}
	}

	if config.DGN == Local {
		config.Host = "127.0.0.1"
	} else {
		config.Host = "0.0.0.0"
	}

	return nil
}
