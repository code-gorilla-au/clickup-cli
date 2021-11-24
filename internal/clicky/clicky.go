package clicky

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/code-gorilla-au/fetch"
)

const (
	baseURL = "https://api.clickup.com/api"
)

var (
	teamsURL = fmt.Sprintf("%s/v2/team", baseURL)
)

type Service struct {
	fetch fetchClient
}

var _ fetchClient = (*fetch.Client)(nil)

func New(fetch fetchClient) Service {
	return Service{
		fetch: fetch,
	}
}

func (s *Service) GetAllWorkspaces(token string) (Teams, error) {
	var teams Teams
	resp, err := s.fetch.Get(teamsURL, map[string]string{
		"Authorization": token,
	})
	if err != nil {
		log.Println("Error getting teams: ", err)
		return teams, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("Error closing body: ", err)
		}
	}()
	if err := json.NewDecoder(resp.Body).Decode(&teams); err != nil {
		log.Println("Error decoding response: ", err)
		return teams, err
	}
	return teams, nil
}

func (s *Service) GetSpaces(workID string, token string) (Spaces, error) {
	url := fmt.Sprintf("%s/%s/space", teamsURL, workID)
	var spaces Spaces
	resp, err := s.fetch.Get(url, map[string]string{
		"Authorization": token,
	})
	if err != nil {
		log.Println("Error getting teams: ", err)
		return spaces, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("Error closing body: ", err)
		}
	}()
	if err := json.NewDecoder(resp.Body).Decode(&spaces); err != nil {
		log.Println("Error decoding response: ", err)
		return spaces, err
	}
	return spaces, nil
}
