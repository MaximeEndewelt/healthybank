package clients

import (
	"github.com/faceit/go-tools/logger"
	"net/http"
	"time"
)

const (
	GetAccountInformation = "https://api-sandbox.starlingbank.com/api/v2/accounts"
)

type service struct {
	lgr        *logger.StandardLogger
	httpClient *http.Client
}

func New() StarlingClient {
	return &service{
		lgr: logger.Get(),
		httpClient: &http.Client{
			Timeout: 50 * time.Second,
		},
	}
}

func (s *service) GetAccountInformation() error {

	req, err := http.NewRequest("GET", GetAccountInformation, nil)
	if err != nil {
		s.lgr.Errorf("Error initializing request", err)
	}

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJQUzI1NiIsInppcCI6IkdaSVAifQ")
	req.Header.Add("accept", "application/json")

	res, err := s.httpClient.Do(req)
	if err != nil {
		s.lgr.Error(err)
	}

	if res.StatusCode > 400 {
		s.lgr.Errorf("Bad request : %s", res.Body)
	} else {
		s.lgr.Infof("Good request : %s", res.Body)
	}

	defer res.Body.Close()
	return nil
}
