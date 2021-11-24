package commands

import (
	"github.com/code-gorilla-au/clickup-cli/internal/clicky"
)

type clickupClient interface {
	GetAllWorkspaces(token string) (clicky.Teams, error)
	GetSpaces(workID string, token string) (clicky.Spaces, error)
}

type storage interface {
	Create() error
	Save(data interface{}) error
	Load(dest interface{}) error
	Exists() bool
}
