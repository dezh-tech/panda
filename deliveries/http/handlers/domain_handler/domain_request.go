package domainhandler

type DomainCreateRequest struct {
	Domain                 string `json:"domain" validate:"required,hostname" form:"domain" query:"domain"`
	BasePricePerIdentifier uint   `json:"base_price_per_identifier" validate:"required,min=1" form:"base_price_per_identifier" query:"base_price_per_identifier"`
	DefaultTTL             uint32 `json:"default_ttl" validate:"required,min=1" form:"default_ttl" query:"default_ttl"`
	Status                 string `json:"status" validate:"required,oneof=active inactive" form:"status" query:"status"`
}
