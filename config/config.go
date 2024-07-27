package config

import "os"

var (
	ConnectionString = os.Getenv("CONNECTION_STRING")
)
