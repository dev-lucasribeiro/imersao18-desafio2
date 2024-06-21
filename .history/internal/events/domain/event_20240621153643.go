package domain

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}