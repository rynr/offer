package main

import (
	"github.com/satori/go.uuid"
	"time"
)

type Offer struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Offers []Offer
