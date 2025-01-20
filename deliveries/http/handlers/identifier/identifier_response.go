package handlers

import "time"

type IdentifierGetResponse struct {
	Name           string    `json:"name"`
	Pubkey         string    `json:"pubkey"`
	DomainID       string    `json:"domain_id"`
	ExpiresAt      time.Time `json:"expires_at"`
	FullIdentifier string    `json:"full_identifier"`
}
