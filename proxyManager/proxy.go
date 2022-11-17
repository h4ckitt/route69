package proxy

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"route69/config"
)

type ProxyMan struct {
	*config.ProxyConfiguration
}

func NewProxyManager(r *config.ProxyConfiguration) *ProxyMan {
	return &ProxyMan{r}
}

func (p *ProxyMan) Start() {
	log.Printf("Starting Proxy Server On: %s\n", p.ListenAddress)
	l, err := net.Listen("tcp", p.ListenAddress)

	if err != nil {
		log.Fatalln(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Printf("Encounteredc An Error Trying To Accept New Connection: %v\n", err)
			continue
		}

		go p.handle(conn)

	}
}

func (p *ProxyMan) handle(c net.Conn) {
	resp := &http.Response{}
	defer func() {
		_ = c.Close()
	}()

	buf := make([]byte, 2048)

	n, err := c.Read(buf)

	if err != nil {
		log.Println("Error occurred while trying to read connection contents")
		return
	}

	payload := buf[:n]
	req, err := http.ReadRequest(bufio.NewReader(bytes.NewBuffer(payload)))

	if err != nil {
		log.Println("Error occurred while converting payload to valid request")
		return
	}

	token := req.Header.Get("X-Telegram-Bot-Api-Secret-Token")

	route := p.GetRoute(token)

	if route == "" {
		log.Println("Token Not Recognized, Ignoring .....")
		resp.StatusCode = 418
		_ = resp.Write(c)
		return
	}

	//	req.RemoteAddr = fmt.Sprintf("http://%s", route)
	u, _ := url.Parse(fmt.Sprintf("http://%s%s", route, req.RequestURI))
	req.RequestURI = ""
	req.URL = u

	resp, err = http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("An error occurred while serving that request ...... is the server up?: %v\n", err)
		resp = &http.Response{}
		resp.StatusCode = 422
		_ = resp.Write(c)
		return
	}

	if err := resp.Write(c); err != nil {
		log.Printf("An error occurred forwarding response back to webserver: %v\n", err)
	}
}
