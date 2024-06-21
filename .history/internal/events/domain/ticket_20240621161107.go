package domain

import "errors"

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

type TicketKind string

const (
	TicketKindHalf TicketKind = "half" // Half-price 
	TicketKindFull TicketKind = "full" // Full-price ticket
)

func IsValidTicketKind(ticketKind TicketKind) bool {
	return ticketKind == TicketKindHalf || ticketKind == TicketKindFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketKindHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return errors.New("ticket price must be greater than zero")
	}
	return nil
}