package anypoint

import (
	"github.com/rhoegg/go-anypoint"
	"log"
)

type Config struct {
	Username string
	Password string
}

func (c *Config) Client() (*go_anypoint.Client, error) {
	client := go_anypoint.NewClient(nil)
	log.Printf("[INFO] Anypoint Client configured for URL: %s", client.BaseURL.String())

	return client, nil
}