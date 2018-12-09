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
	client.Username = c.Username
	client.Password = c.Password
	log.Printf("[INFO] Anypoint Client configured for URL: %s", client.BaseURL.String())
	log.Printf("[DEBUG] Anypoint username to use: %s", c.Username)

	return client, nil
}
