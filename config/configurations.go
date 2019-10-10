package config

import (
	"github.com/spf13/viper"
)

const (
	//PortNumber is used to configure the port of the http listener
	PortNumber = "api.handlers.http.portnumber"
)

var defaultConfigPaths = []string{
	"/etc/viper",
	"./",
}

func init() {
	viper.SetDefault(PortNumber, ":8080")
}