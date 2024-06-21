package domain

import (
	"errors"
	"time"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

func (e *Event) Validate() error {
	if e.Name == "" {
		return errors.New("event name is required")
	}

	if e.Date.Before(time.Now()) {
		return errors.New("event date must be in the future")
	}

	if e.Capacity <= 0 {
		return errors.New("event capacity must be greater than zero")
	}

	if e.Price <= 0 {
		return errors.New("event price must be greater than zero")
	}

	return nil
}