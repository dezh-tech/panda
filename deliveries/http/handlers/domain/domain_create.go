package handlers

import (
	"errors"
	"net/http"

	"github.com/dezh-tech/panda/pkg"
	"github.com/dezh-tech/panda/pkg/validator"
	domainService "github.com/dezh-tech/panda/services/domain"
	"github.com/labstack/echo/v4"
)

// CreateDomain creates a new domain.
//
//	@Summary      Create a new domain
//	@Description  Create a new domain with the specified attributes.
//	@Tags         domains
//	@Accept       json
//	@Produce      json
//	@Param        domain  body      DomainCreateRequest  true  "Domain creation payload"
//	@Success      200     {object}  pkg.ResponseDto{data=DomainCreateResponse} "Domain created successfully"
//	@Failure      400     {object}  pkg.ResponseDto[validator.Varror]                    "Bad Request - Validation error"
//	@Failure      500     {object}  pkg.ResponseDto[string]                              "Internal Server Error"
//	@Router       /domains [post]
func (dh Domain) create(c echo.Context) error {
	req := new(DomainCreateRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, pkg.ResponseDto{
			Success: false,
			Error:   validator.Varror{Error: "invalid input"},
		})
	}

	// Validate the request payload
	v := validator.NewValidator()
	validationErrors := v.Validate(req)
	if validationErrors != nil {
		return echo.NewHTTPError(http.StatusBadRequest, pkg.ResponseDto{
			Success: false,
			Error:   validator.Varror{ValidationErrors: validationErrors},
		})
	}

	// Call the domain service to create the domain
	ctx := c.Request().Context()
	resp, err := dh.service.Create(ctx, req.Domain, req.Status, req.BasePricePerIdentifier, req.DefaultTTL)
	if err != nil {
		if errors.Is(err, domainService.ErrIsExist) {
			return echo.NewHTTPError(http.StatusConflict, pkg.ResponseDto{
				Success: false,
				Error:   validator.Varror{Error: err.Error()},
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, pkg.ResponseDto{
			Success: false,
			Error:   validator.Varror{Error: echo.ErrInternalServerError.Error()},
		})
	}

	return c.JSON(http.StatusOK, pkg.ResponseDto{Success: true, Data: DomainCreateResponse{ID: resp}})
}
