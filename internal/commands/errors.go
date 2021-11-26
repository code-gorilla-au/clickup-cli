package commands

import "errors"

var (
	ErrNoToken   = errors.New("no token")
	ErrNoTeamID  = errors.New("no team / workspace id")
	ErrNoSpaceID = errors.New("no space id")
)
