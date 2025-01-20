package handlers

import (
	"errors"
	"net/http"

	"github.com/dezh-tech/panda/pkg"
	"github.com/dezh-tech/panda/pkg/validator"
	userService "github.com/dezh-tech/panda/services/user"
	"github.com/labstack/echo/v4"
)

// CreateUser creates a new user.
//
//	@Summary      Create a new user
//	@Description  Creates a new user using the provided public key. The request payload must include a valid public key for successful user creation.
//	@Tags         users
//	@Accept       json
//	@Produce      json
//	@Param        user  body      UserCreateRequest  true  "Payload containing the public key for user creation"
//	@Success      200   {object}  pkg.ResponseDto                           "User created successfully"
//	@Failure      400   {object}  pkg.ResponseDto[validator.Varror]         "Bad Request - Invalid input or validation errors"
//	@Failure      409   {object}  pkg.ResponseDto[validator.Varror]         "Conflict - User with the specified public key already exists"
//	@Failure      500   {object}  pkg.ResponseDto[string]                   "Internal Server Error - An unexpected error occurred"
//	@Router       /users [post]
func (uh User) create(c echo.Context) error {
	req := new(UserCreateRequest)
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
	_, err := uh.service.Create(ctx, req.Pubkey)
	if err != nil {
		if errors.Is(err, userService.ErrIsExist) {
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

	return c.JSON(http.StatusOK, pkg.ResponseDto{Success: true, Data: nil})
}
