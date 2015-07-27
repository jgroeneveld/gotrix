package cfg

import (
	"os"
	"strings"
)

var (
	ProdEnv = Env() == "production"
	TestEnv = Env() == "test"
	DevEnv  = Env() == "development"
)

func Env() string {
	switch {
	case strings.HasSuffix(os.Args[0], ".test"): // must be first (there might be tests being run in app env)
		return "test"
	case os.Getenv("APP_ENV") == "production":
		return "production"
	case os.Getenv("APP_ENV") == "staging":
		return "staging"
	default:
		return "development"
	}
}
