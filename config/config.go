package config

import (
	"os"
)

type Config struct {
	Debug            bool
	Timeout          int
	Addr             string
	VersionTimestamp string
}



var DefaultConfig Config

func init() {
	InitWithEnv()
}


func InitWithEnv() error {

	_debug := os.Getenv("HEDWI_API_DEBUG")
	debug := _debug == "true"

	httpAddr := os.Getenv("HEDWI_DOCUMENT_HTTP_ADDR")        //Addr
	if httpAddr == "" {
		httpAddr = ":40012"
	}
	
	
	DefaultConfig = Config{
		VersionTimestamp: "",
		Debug:            debug,
		Addr:             httpAddr,
	}

	return nil
}