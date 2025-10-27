package plex

import (
	// "os"
	"fmt"
	"time"
	"net/http"
	cfg "github.com/PFNASS/go-arrs/pkg/config"
)

type PlexConfig struct {
	IP         string
	Identifier string
	Port       int
}

type PlexClient struct {
	httpClient *http.Client
	Config *PlexConfig
}

func NewPlexClient(host string, apiKey, apiSecret string, timeout time.Duration) *PlexClient {
 client := &http.Client{
  Timeout: timeout,
 }
 return &PlexClient{
  httpClient: client,
  Config: &PlexConfig{
   IP:         host,
   Identifier: apiKey,
   Port:       32400,
  },
 }
}

func loadPlexConfig() (*PlexConfig, error) {
	v, err := cfg.LoadConfig()
	if err != nil {
		return nil, err
	}

	plexConfig := &PlexConfig{
		IP:         v.GetString("plex.ip"),
		Identifier: v.GetString("plex.identifier"),
		Port:       v.GetInt("plex.port"),
	}

	return plexConfig, nil
}

func (c *PlexClient) do(method, endpoint string, params map[string]string) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
	return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	for key, val := range params {
	q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}