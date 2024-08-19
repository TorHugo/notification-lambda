package domain

type Notification struct {
	Contact    string      `json:"contact"`
	Template   string      `json:"template"`
	Parameters []Parameter `json:"parameters"`
}
