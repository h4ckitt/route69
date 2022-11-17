package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	conf        *ProxyConfiguration
	routesCache map[string]string
)

func ReadInConfig() (*ProxyConfiguration, error) {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	fmt.Println(string(file))
	routesCache = make(map[string]string)

	var c config

	err = yaml.Unmarshal(file, &c)

	if err != nil {
		return nil, err
	}

	for _, route := range c.Routes {
		routesCache[route.Token] = route.Route
	}

	if c.ListenOn == "" {
		c.ListenOn = "localhost:6969"
	}

	conf = &ProxyConfiguration{c.ListenOn, routesCache}

	return conf, nil
}

// TODO: Complete When I Have Enough Brain Cells. A Normal Golang Map Shouldn't Be Concurrently Written To And Read From. Need To Re-Implement Using Sync.Map Or A Custom Solution.
func RefreshConfig() {
	file, err := os.ReadFile("config.yaml")

	if err != nil {
		log.Println(err)
		return
	}

	c := config{}

	err = yaml.Unmarshal(file, &c)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, route := range c.Routes {
		routesCache[route.Token] = route.Route
	}

}

func GetConfig() *ProxyConfiguration {
	return conf
}
