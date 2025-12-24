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

func (p *FactProvider) getFact(factURL string) string {
	resp, err := http.Get(factURL)
	if err != nil {
		log.Printf("can't get fact: %s", err)
	}
	defer resp.Body.Close()

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

func (p *FactProvider) Recieve() (int64, string) {
	var s string = p.getFact(factsURL)
	return 0, s
}
