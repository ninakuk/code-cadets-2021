package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
)

const axilisFeedURL = "http://18.193.121.232/axilis-feed"

type AxilisOfferFeed struct {
	updates    chan models.Odd
	httpClient *http.Client
}

func NewAxilisOfferFeed(
	httpClient *http.Client,
) *AxilisOfferFeed {
	return &AxilisOfferFeed{
		updates:    make(chan models.Odd),
		httpClient: httpClient,
	}
}

func (a *AxilisOfferFeed) Start(ctx context.Context) error {
	defer close(a.updates)
	defer log.Printf("shutting down %s", a)

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-time.After(time.Second * 3):
			response, err := a.httpClient.Get(axilisFeedURL)
			if err != nil {
				log.Println("axilis offer feed, http get", err)
				continue
			}
			a.processResponse(ctx, response)
		}
	}
}

func (a *AxilisOfferFeed) GetUpdates() chan models.Odd {
	return a.updates
}

func (a *AxilisOfferFeed) String() string {
	return "axilis offer feed"
}

func (a *AxilisOfferFeed) processResponse(ctx context.Context, response *http.Response) {
	defer response.Body.Close()

	var axilisOfferOdds []axilisOfferOdd
	err := json.NewDecoder(response.Body).Decode(&axilisOfferOdds)
	if err != nil {
		log.Println("axilis offer feed, json decode", err)
		return
	}

	for _, axilisOdd := range axilisOfferOdds {
		odd := models.Odd{
			Id:          axilisOdd.Id,
			Name:        axilisOdd.Name,
			Match:       axilisOdd.Match,
			Coefficient: axilisOdd.Details.Price,
			Timestamp:   time.Now(),
		}

		// IMPORTANT SELECT!!!
		// show an example
		select {
		case <-ctx.Done():
			return
		case a.updates <- odd:
			// do nothing
		}
	}
}

type axilisOfferOdd struct {
	Id      string
	Name    string
	Match   string
	Details axilisOfferOddDetails
}

type axilisOfferOddDetails struct {
	Price float64
}
