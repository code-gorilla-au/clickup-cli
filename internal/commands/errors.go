package commands

import "errors"

var (
	ErrNoToken       = errors.New("personal access token is required")
	ErrNoWorkspaceID = errors.New("workspace id is required")
	ErrNoSpaceID     = errors.New("workspace space id is required")
)
