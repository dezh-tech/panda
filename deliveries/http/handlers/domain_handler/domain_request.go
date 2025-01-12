package domainhandler

type DomainCreateRequest struct {
	Domain                 string `form:"domain"                    json:"domain"                    validate:"required,hostname"`
	BasePricePerIdentifier uint   `form:"base_price_per_identifier" json:"base_price_per_identifier" validate:"required,min=1"`
	DefaultTTL             uint32 `form:"default_ttl"               json:"default_ttl"               validate:"required,min=1"`
	Status                 string `form:"status"                    json:"status"                    validate:"required,oneof=active inactive"`
}
