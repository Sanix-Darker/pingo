package main

import (
	"os"
)

var (
	VERSION         = os.Getenv("VERSION")
	PING_PORT       = os.Getenv("PING_PORT")
	MONGO_DB        = os.Getenv("MONGO_DB")
	MONGO_HOST      = os.Getenv("MONGO_HOST")
	PING_COLLECTION = os.Getenv("PING_COLLECTION")
	ASSETS_PATH     = "../static"
	TEMPLATES_PATH  = "../templates"
)
