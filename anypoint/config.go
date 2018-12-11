package anypoint

import (
	"github.com/rhoegg/go-anypoint/anypointplatform"
	"log"
)

type Config struct {
	Username string
	Password string
}

func (c *Config) Client() (*anypointplatform.Client, error) {
	client := anypointplatform.NewClient(nil)
	client.Username = c.Username
	client.Password = c.Password
	log.Printf("[INFO] Anypoint Client configured for URL: %s", client.BaseURL.String())
	log.Printf("[DEBUG] Anypoint username to use: %s", c.Username)

	return client, nil
}
