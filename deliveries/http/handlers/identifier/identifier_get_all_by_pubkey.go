package handlers

import (
	"net/http"

	"github.com/dezh-tech/panda/pkg"
	"github.com/dezh-tech/panda/pkg/validator"
	"github.com/labstack/echo/v4"
)

// IdentifierGetAllByPubkey retrieves all identifiers.
//
//	@Summary      Retrieve all identifiers
//	@Description  Get a list of all identifiers associated with the provided public key.
//	@Tags         identifiers
//	@Accept       json
//	@Produce      json
//	@Param        Authorization  header    string  true   "Authorization"
//	@Success      200     {object}  pkg.ResponseDto{data=[]IdentifierGetResponse} "identifiers retrieved successfully"
//	@Failure      500     {object}  pkg.ResponseDto[string]                        "Internal Server Error"
//	@Router       /identifiers [get]
func (dh Identifier) getAllByPubkey(c echo.Context) error {
	ctx := c.Request().Context()
	idns, err := dh.service.GetAllByPubKey(ctx, c.Get("pubkey").(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, pkg.ResponseDto{
			Success: false,
			Error:   validator.Varror{Error: echo.ErrInternalServerError.Error()},
		})
	}

	idnRes := make([]IdentifierGetResponse, 0)
	for _, d := range *idns {
		idnRes = append(idnRes, IdentifierGetResponse{
			Name:           d.Name,
			Pubkey:         d.Pubkey,
			DomainID:       d.DomainID,
			FullIdentifier: d.FullIdentifier,
			ExpiresAt:      d.ExpiresAt,
		})
	}

	// Respond with the created domain's ID
	return c.JSON(http.StatusOK, pkg.ResponseDto{Success: true, Data: idnRes})
}
