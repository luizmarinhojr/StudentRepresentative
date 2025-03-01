package config

import "os"

var (
	PSQL_HOST_PROD   = os.Getenv("PSQL_HOST")
	PSQL_PORT_PROD   = os.Getenv("PSQL_PORT")
	PSQL_USER_PROD   = os.Getenv("PSQL_USER")
	PSQL_PASS_PROD   = os.Getenv("PSQL_PASS")
	PSQL_DBNAME_PROD = os.Getenv("PSQL_DBNAME")
)
