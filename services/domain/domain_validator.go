package domainService

type insertDomainValidation struct {
	Domain                 string `validate:"required,url"`
	BasePricePerIdentifier uint   `validate:"required"`
	Status                 string `validate:"oneof=ACTIVE INACTIVE"`
}
