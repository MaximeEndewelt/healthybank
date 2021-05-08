package clients

import (
	"log"
	"net/http"
	"time"
)

const (
	GetAccountInformation = "https://api-sandbox.starlingbank.com/api/v2/accounts"
)

type service struct {
	httpClient *http.Client
}

func New() StarlingClient {
	return &service{
		httpClient: &http.Client{
			Timeout: 50 * time.Second,
		},
	}
}

func (s *service) GetAccountInformation() error {

	req, err := http.NewRequest("GET", GetAccountInformation, nil)
	if err != nil {
		log.Printf("Error initializing request: %s", err)
	}

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJQUzI1NiIsInppcCI6IkdaSVAifQ")
	req.Header.Add("accept", "application/json")

	res, err := s.httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	if res.StatusCode > 400 {
		log.Printf("Bad request : %s", res.Body)
	} else {
		log.Printf("Good request : %s", res.Body)
	}

	defer res.Body.Close()
	return nil
}
