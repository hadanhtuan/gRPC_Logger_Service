package conf

import "os"

type config struct {
	Env     string
	Version string
	DBName  string
	DBAuth  string
}

// Config main config object
var Config *config

func init() {
	env := os.Getenv("env")

	DBName := os.Getenv("DBName")
	DBAuth := os.Getenv("DBAuth")
	version := os.Getenv("version")
	if env == " " {
		env = "local"
	}

	switch env {
	case "local":
		Config = &config{
			Env:     "local",
			Version: version,
			DBName:  DBName,
			DBAuth:  DBAuth,
		}
		break
	case "uat":
		break
	case "prd":
		break
	}
}
