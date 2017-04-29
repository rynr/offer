package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

var offers Offers

func init() {
	RepoCreateOffer(Offer{Name: "BMW 3er"})
	RepoCreateOffer(Offer{Name: "BMWi"})
	RepoCreateOffer(Offer{Name: "Mini"})
}

func RepoFindOffer(id uuid.UUID) Offer {
	for _, o := range offers {
		if o.Id == id {
			return o
		}
	}
	// return empty Offer if not found
	return Offer{}
}

func RepoCreateOffer(o Offer) Offer {
	o.Id = uuid.NewV4()
	offers = append(offers, o)
	return o
}

func RepoDestroyOffer(id uuid.UUID) error {
	for i, o := range offers {
		if o.Id == id {
			offers = append(offers[:i], offers[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Offer with id of %d to delete", id)
}
