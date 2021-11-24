package clicky

type User struct {
	ID             int    `json:"id,omitempty"`
	Color          string `json:"color,omitempty"`
	CustomRole     string `json:"custom_role,omitempty"`
	DateInvited    string `json:"date_invited,omitempty"`
	DateJoined     string `json:"date_joined,omitempty"`
	Email          string `json:"email,omitempty"`
	Initials       string `json:"initials,omitempty"`
	LastActive     string `json:"last_active,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	Role           int    `json:"role,omitempty"`
	Username       string `json:"username,omitempty"`
}

type Member struct {
	User User `json:"user,omitempty"`
}

type Team struct {
	Avatar  string   `json:"avatar,omitempty"`
	Color   string   `json:"color,omitempty"`
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Members []Member `json:"members,omitempty"`
}

type Teams struct {
	Teams []Team `json:"teams,omitempty"`
}

type Status struct {
	Status     string `json:"status,omitempty"`
	Type       string `json:"type,omitempty"`
	Orderindex int    `json:"orderindex,omitempty"`
	Color      string `json:"color,omitempty"`
}

type FeatureDueDates struct {
	Enabled            bool `json:"enabled,omitempty"`
	StartDate          bool `json:"start_date,omitempty"`
	RemapDueDates      bool `json:"remap_due_dates,omitempty"`
	RemapClosedDueDate bool `json:"remap_closed_due_date,omitempty"`
}

type Feature struct {
	Enabled bool `json:"enabled,omitempty"`
}

type Features struct {
	DueDates          FeatureDueDates `json:"due_dates,omitempty"`
	TimeTracking      Feature         `json:"time_tracking,omitempty"`
	Tags              Feature         `json:"tags,omitempty"`
	TimeEstimates     Feature         `json:"time_estimates,omitempty"`
	Checklists        Feature         `json:"checklists,omitempty"`
	CustomFields      Feature         `json:"custom_fields,omitempty"`
	RemapDependencies Feature         `json:"remap_dependencies,omitempty"`
	DependencyWarning Feature         `json:"dependency_warning,omitempty"`
}

type Space struct {
	ID                string   `json:"id,omitempty"`
	Name              string   `json:"name,omitempty"`
	Private           bool     `json:"private,omitempty"`
	Statuses          []Status `json:"statuses,omitempty"`
	MultipleAssignees bool     `json:"multiple_assignees,omitempty"`
	Features          Features `json:"features,omitempty"`
}
