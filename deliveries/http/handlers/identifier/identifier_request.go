package handlers

type IdentifierCreateRequest struct {
	DomainID string `json:"domain_id"             validate:"required"`
	Pubkey   string `json:"pubkey"               validate:"required"`
	Name     string `json:"name"                  validate:"required"`
}
