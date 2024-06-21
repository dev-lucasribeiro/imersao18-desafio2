package domain

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

type TicketKind string

const (
	TicketKindHalf TicketKind = "half" // Half-price ticket
	TicketKindFull TicketKind = "full" // Full-price ticket
)
