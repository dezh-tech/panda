package domainhandler

import (
	"fmt"
	"net/http"

	domainhandler "github.com/dezh-tech/panda/deliveries/http/handlers/domain_handler/dto"
	"github.com/dezh-tech/panda/pkg/validator"
	domainService "github.com/dezh-tech/panda/services/domain"
	"github.com/labstack/echo/v4"
)

// CreateDomain creates a new domain.
//
//	@Summary      Create a new domain
//	@Description  Accepts a JSON payload to create a new domain with the specified attributes.
//	@Tags         domain
//	@Accept       json
//	@Produce      json
//	@Param        domain  body      domainhandler.DomainCreateRequest  true  "Domain creation payload"
//	@Success      200     {object}  domainhandler.DomainCreateResponse "Domain created successfully"
//	@Failure      400     {object}  map[string]string                  "Bad Request - Invalid input"
//	@Failure      500     {object}  map[string]string                  "Internal Server Error"
//	@Router       /domains [post]
func (h Handler) domainCreate(c echo.Context) error {
	req := new(domainhandler.DomainCreateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	v := validator.NewValidator()
	validationErrors := v.Validate(req)
	if validationErrors != nil {
		return echo.NewHTTPError(http.StatusBadRequest, &validator.Varror{ValidationErrors: validationErrors})
	}

	resp, err := h.domainSvc.Create(domainService.DomainInsertArgs{
		Domain:                 req.Domain,
		BasePricePerIdentifier: req.BasePricePerIdentifier,
		DefaultTTL:             req.DefaultTTL,
		Status:                 req.Status,
	})
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, &domainhandler.DomainCreateResponse{ID: resp.ID})
}
