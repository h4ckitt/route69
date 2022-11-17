package config

import "fmt"

type Configuration struct {
	routes map[string]string
}

type routes struct {
	Token string
	Route string
}

func (c *Configuration) Get(key string) string {
	if route, exists := c.routes[key]; exists {
		return route
	}

	return ""
}

func (c *Configuration) PrintRoutes() {
	for token, route := range c.routes {
		fmt.Printf("%s -> %s\n", token, route)
	}

}
