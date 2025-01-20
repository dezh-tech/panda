package handlers

type UserCreateRequest struct {
	Pubkey string `json:"pubKey"                    validate:"required"`
}
