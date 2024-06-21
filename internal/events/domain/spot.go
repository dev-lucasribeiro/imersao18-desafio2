package domain

type Spot struct {
	ID       string
	EventID  string
	Name     string // all name uses the rule: Letter+Number. Ex: A1, B2, C3, etc.
	Status   SpotStatus
	TicketID string
}

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)