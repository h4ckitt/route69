package config

import "fmt"

type ProxyConfiguration struct {
	ListenAddress string
	routes        map[string]string
}

type config struct {
	ListenOn        string
	RefreshInterval string
	Routes          []routes
}

type routes struct {
	Token string
	Route string
}

func (c *ProxyConfiguration) GetRoute(key string) string {
	if route, exists := c.routes[key]; exists {
		return route
	}

	return ""
}

func (c *ProxyConfiguration) PrintRoutes() {
	for token, route := range c.routes {
		fmt.Printf("%s -> %s\n", token, route)
	}

}
