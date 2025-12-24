package catfacts

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const factsURL string = "https://catfact.ninja/fact"

type FactProvider struct {
	ProviderID string
}

func New() *FactProvider {
	return &FactProvider{ProviderID: "catfacts"}
}

func (p *FactProvider) getFact(factURL string, client http.Client) string {
	resp, err := client.Get(factURL)
	if err != nil {
		log.Printf("can't get fact: %s", err)
	}
	defer resp.Body.Close()
	defer client.CloseIdleConnections()

	var result struct {
		Result string `json:"fact"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("can't get fact: %s", err)
	}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("can't get fact: %s", err)
	}
	return result.Result
}

func (p *FactProvider) Recieve(client http.Client) (int64, string) {
	var s string = p.getFact(factsURL, client)
	return 0, s
}
