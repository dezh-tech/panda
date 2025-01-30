package handlers

import (
	"errors"
	"net/http"

	"github.com/dezh-tech/panda/pkg"
	"github.com/dezh-tech/panda/pkg/validator"
	domainService "github.com/dezh-tech/panda/services/domain"
	identifierService "github.com/dezh-tech/panda/services/identifier"
	userService "github.com/dezh-tech/panda/services/user"
	"github.com/labstack/echo/v4"
)

// CreateIdentifier creates a new identifier.
//
//	@Summary      Create a new identifier
//	@Description  Creates a new identifier with the specified attributes. Returns success if the identifier is created successfully or relevant error messages if the creation fails.
//	@Tags         identifiers
//	@Accept       json
//	@Produce      json
//	@Param        identifier  body      IdentifierCreateRequest  true  "Identifier creation payload"
//	@Success      200         {object}  pkg.ResponseDto "Identifier created successfully"
//	@Failure      400         {object}  pkg.ResponseDto{error=validator.Varror}        "Bad Request - Validation error or invalid input"
//	@Failure      409         {object}  pkg.ResponseDto{error=validator.Varror}        "Conflict - Identifier already exists"
//	@Failure      500         {object}  pkg.ResponseDto{error=validator.Varror}        "Internal Server Error - Unexpected server error"
//	@Router       /identifiers [post]
func (dh Identifier) create(c echo.Context) error {
	req := new(IdentifierCreateRequest)
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
	_, err := dh.service.Create(ctx, req.Name, req.DomainID, req.Pubkey)
	if err != nil {
		if errors.Is(err, domainService.ErrNotFound) || errors.Is(err, userService.ErrNotFound) || errors.Is(err, identifierService.ErrIsExist) || errors.Is(err, identifierService.Err) {
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

	return c.JSON(http.StatusOK, pkg.ResponseDto{Success: true})
}
