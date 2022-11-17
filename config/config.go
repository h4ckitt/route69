package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var conf *Configuration

func ReadInConfig() (*Configuration, error) {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	fmt.Println(string(file))
	routesCache := make(map[string]string)

	var routes []routes

	err = yaml.Unmarshal(file, &routes)

	if err != nil {
		return nil, err
	}

	for _, route := range routes {
		routesCache[route.Token] = route.Route
	}

	conf = &Configuration{routesCache}

	return conf, nil
}

func GetConfig() *Configuration {
	return conf
}
