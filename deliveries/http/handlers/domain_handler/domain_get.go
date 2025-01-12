package domainhandler

import (
	"net/http"

	"github.com/dezh-tech/panda/pkg"
	"github.com/dezh-tech/panda/pkg/validator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// DomainGetAll retrieves all domains.
//
//	@Summary      Retrieve all domains
//	@Description  Get a list of all domains with their attributes.
//	@Tags         domains
//	@Accept       json
//	@Produce      json
//	@Success      200  {object}  pkg.ResponseDto{data=[]DomainGetResponse} "Domains retrieved successfully"
//	@Failure      500  {object}  pkg.ResponseDto[string]                                "Internal Server Error"
//	@Router       /domains [get]
func (h Handler) domainGetAll(c echo.Context) error {
	ctx := c.Request().Context()
	domains, err := h.domainService.GetAll(ctx, bson.M{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, pkg.ResponseDto{Success: false, Error: validator.Varror{Error: echo.ErrInternalServerError.Error()}})
	}

	domainsRes := make([]DomainGetResponse,0)
	for _, d := range *domains {
		domainsRes = append(domainsRes, DomainGetResponse{
			Domain:                 d.Domain,
			BasePricePerIdentifier: d.BasePricePerIdentifier,
			DefaultTTL:             d.DefaultTTL,
			Status:                 d.Status,
		})
	}

	// Respond with the created domain's ID
	return c.JSON(http.StatusOK, pkg.ResponseDto{Success: true, Data: domainsRes})
}
