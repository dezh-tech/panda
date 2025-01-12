package domainhandler

type DomainCreateResponse struct {
	ID interface{} `json:"id"`
}

type DomainGetResponse struct {
	Domain                 string `json:"domain"`
	BasePricePerIdentifier uint   `json:"base_price_per_identifier"`
	DefaultTTL             uint32 `json:"default_ttl"`
	Status                 string `json:"status"`
}
