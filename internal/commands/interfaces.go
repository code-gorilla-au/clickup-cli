package commands

import (
	"github.com/code-gorilla-au/clickup-cli/internal/clicky"
)

type clickUpClient interface {
	GetAllWorkspaces(token string) (clicky.Teams, error)
	GetSpaces(workID string, token string) (clicky.Spaces, error)
	GetFolders(spaceID string, token string) (clicky.Folders, error)
}

type storage interface {
	Create() error
	Save(data interface{}) error
	Load(dest interface{}) error
	Exists() bool
}
